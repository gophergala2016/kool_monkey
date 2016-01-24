Summary: Kool monkey is a distributed system to monitor webpages. This package provides the server part of Kool monkey.
Name: kool-server
Version: %{version}
Release: %{release}
License: GPLv3
Group: Applications/Multimedia
Requires: postgresql94
URL: https://github.com/gophergala2016/kool_monkey

BuildRoot: %{_tmppath}/%{name}-%{version}-%{release}-root
# Should work in centos, but it's not working in debian :(
#BuildRequires: golang
Requires(post): /sbin/chkconfig /usr/sbin/useradd
Requires(preun): /sbin/chkconfig, /sbin/service
Requires(postun): /sbin/service
Provides: kool-server

%description
Kool monkey is a distributed system to monitor webpages. This package
provides the server part of Kool monkey.

%build
PROD_BUILD=1 %{__make} kool-server
PROD_BUILD=1 %{__make} install

%install
%{__rm} -rf %{buildroot}
mkdir -p %{buildroot}%{_bindir}
mkdir -p %{buildroot}%{_sysconfdir}
mkdir -p %{buildroot}%{_datadir}
mkdir -p %{buildroot}%{_exec_prefix}/www
mkdir -p %{buildroot}%{_exec_prefix}/dashboard
%{__install} -Dp -m 0755 bin/kool-server %{buildroot}%{_bindir}/kool-server
%{__install} -Dp -m 0644 scripts/db/*.sql %{buildroot}%{_datadir}
%{__install} -Dp -m 0644 front/www/* %{buildroot}%{_exec_prefix}/www
%{__install} -Dp -m 0644 front/dashboard/* %{buildroot}%{_exec_prefix}/dashboard
%{__install} -Dp -m 0755 scripts/init/kool-server %{buildroot}%{_sysconfdir}/init.d/kool-server
%{__install} -Dp -m 0755 systemd/kool-server.service %{buildroot}%{_systemddir}/kool-server.service
%{__install} -Dp -m 0644 dev-env/conf/kool-server.conf %{buildroot}%{_sysconfdir}/kool-server.conf
%{__install} -p -d -m 0755 %{buildroot}%{pid_dir}

%pre
/usr/sbin/useradd -c 'monkey' -u 499 -s /bin/false -r -d %{_prefix} monkey 2> /dev/null || :

%preun
if [ $1 = 0 ]; then
    # when the preun section is run, we've got stdin attached.  If we
    # call stop() in the redis init script, it will pass stdin along to
    # the redis-cli script; this will cause redis-cli to read an extraneous
    # argument, and the redis-cli shutdown will fail due to the wrong number
    # of arguments.  So we do this little bit of magic to reconnect stdin
    # to the terminal
    term="/dev/$(ps -p$$ --no-heading | awk '{print $2}')"
    exec < $term

    /sbin/service kool-server stop > /dev/null 2>&1 || :
    /sbin/chkconfig --del kool-server
fi

%post
port="-p 5430"
/sbin/chkconfig --add kool-server
if [ -z `su postgres -c "/usr/bin/psql ${port} -l | grep monkey"` ]; then
	su postgres -c "/usr/bin/createdb ${port} monkey"
	su postgres -c "/usr/bin/psql ${port} monkey -f %{_datadir}/create_db.sql"
else
	su postgres -c "/usr/bin/psql ${port} monkey -f %{_datadir}/upgrade_db.sql"
fi

%postun
port="-p 5430"
su postgres -c "/usr/bin/dropdb ${port} --if-exists monkey"

%clean
%{__rm} -rf %{buildroot}

%files
%defattr(-, root, root, 0755)
%{_bindir}/kool-server
%{_exec_prefix}/www
%{_exec_prefix}/dashboard
%{_datadir}/create_db.sql
%{_datadir}/upgrade_db.sql
%{_datadir}/create_roles.sql
%{_datadir}/create_schema.sql
%{_sysconfdir}/init.d/kool-server
%{_systemddir}/kool-server.service
%config(noreplace) %{_sysconfdir}/kool-server.conf

%changelog
* Sat Jan 23 2016 Pablo Alvarez de Sotomayor Posadillo <palvarez@ritho.net> 0.1-0
- First version of the rpm package.

Summary: Kool monkey is a distributed system to monitor webpages. This package provides the server part of Kool monkey.
Name: kool-server
Version: %{version}
Release: %{release}
License: GPLv3
Group: Applications/Multimedia
Requires: postgresql
URL: https://github.com/gophergala2016/kool_monkey

BuildRoot: %{_tmppath}/%{name}-%{version}-%{release}-root
# Should work in centos, but it's not working in debian :(
#BuildRequires: golang
Requires(post): /sbin/chkconfig /usr/sbin/useradd /etc/init.d/postgresql
Requires(preun): /sbin/chkconfig, /sbin/service
Requires(postun): /sbin/service
Provides: kool-server

%description
Kool monkey is a distributed system to monitor webpages. This package
provides the server part of Kool monkey.

%build
%{__make} kool-server

%install
%{__rm} -rf %{buildroot}
mkdir -p %{buildroot}%{_bindir}
mkdir -p %{buildroot}%{_sysconfdir}
mkdir -p %{buildroot}%{_datadir}
%{__install} -Dp -m 0755 bin/kool-server %{buildroot}%{_bindir}/kool-server
%{__install} -Dp -m 0644 scripts/db/*.sql %{buildroot}%{_datadir}
%{__install} -Dp -m 0755 scripts/init/kool-server %{buildroot}%{_sysconfdir}/init.d/kool-server
%{__install} -Dp -m 0644 conf/kool-server.conf %{buildroot}%{_sysconfdir}/kool-server.conf
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
/sbin/chkconfig --add kool-server
su postgres -c "/usr/bin/createdb monkey"
su postgres -c "/usr/bin/psql monkey -f %{_datadir}/create_db.sql"

%clean
%{__rm} -rf %{buildroot}

%files
%defattr(-, root, root, 0755)
%{_bindir}/kool-server
%{_datadir}/create_db.sql
%{_datadir}/create_roles.sql
%{_datadir}/create_schema.sql
%{_sysconfdir}/init.d/kool-server
%config(noreplace) %{_sysconfdir}/kool-server.conf
%dir %attr(0755,redis,redis) %{_localstatedir}/run

%changelog
* Sat Jan 23 2016 Pablo Alvarez de Sotomayor Posadillo <palvarez@ritho.net> 0.1-0
- First version of the rpm package.

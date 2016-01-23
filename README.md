 Kool Monkey
=============

Kool Monkey is a distributed monitoring system that collect metrics
related to a page loading time (time to respond, time to load, the
waterfall, ...).

The term distributed comes from the fact that the metrics are
measured by an small agent that can be installed in any computer, so
the same users that wants to collect the data to see how their site is
working can install themself the agent to recolect metrics and send
them back to the central server, making the user to contribute to a
common comunity and get information in more and more sites with the
growth of the comunity.

The idea of collecting the metrics is both to show the user (in a web
application) how their site evolve and get alerts (by email and maybe
SMS in the future) when the site get slower or offline for whatever
reason.

The project started in the [Gopher Gala](http://www.gophergala.com/)
2016 (22-24 of January) based on particular needs of the participants.

We are currently doing the management on the project in
[trello](https://trello.com/b/zNxSafya), so feel free to check it out
and see what we have done and what we are planning to do.

## License

GPLv3

## Architecture

This project have three important elements: the agent, the server and
the frontend.

### Agent

The agent is encharged of do the actual measures of the webpage. It's
developed in Go as a daemon that comunicates with the server in order
to obtain the urls to monitor with some parameters (when to start and
end the monitoring, the periodicity of the metrics, what metrics to
run, ...) and send back the actual metrics obtained.

### Server

The server is a central go application that manage the work between
the different agents, store in db the results and comunicate with the
front end in order to obtain the user requests and show back the
results recolected from the agents.

### Frontend

The frontend is where the user can request a new monitoring and see
the different information that have been recolected in a graphical and
easy way.

## Build and deploy

As you have seen, the development of the project is in Go, so we have
provided some Makefiles in order to work with the code. This Makefiles
not only are useful to develop in the project, but to build rpms
packages from the source.

We have also integrated the project with Travis to run the tests,
build the rpms and deploy into production both the server part
(metrics recolection and web frontend) and the agents, so it's become
quite easy to build and deploy the project.

The current status of travis is [![Build Status](https://travis-ci.org/gophergala2016/kool_monkey.svg?branch=master)](https://travis-ci.org/gophergala2016/kool_monkey) 

## Developers

The main developers (and maintainers) of the project are:

* [Guillermo de los Santos](https://github.com/MemoDLSG)
* [Iván Californias](https://github.com/ivan-californias)
* [Pablo Álvarez de Sotomayor](https://github.com/i02sopop)
* [Sergio Visinoni](https://github.com/piffio)

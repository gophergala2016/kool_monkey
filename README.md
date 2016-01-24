# Kool Monkey

Kool Monkey is a Open and Distributed Web monitoring system that collect metrics
related to web site perfomances: Response Time, Uptime, Network Roundtrip, etc.

## Why another monitoring tool?

Web page monitoring falls is usually implemented using one of the 2
following alternatives: Proactive Monitoring and RUM.

Both solutions are in generally expensive and typically lack support for
any but the most common cases. With any of the existing commercial
solutions you're limited by either the Pricing model, which makes it
suitable for establised enterprise businesses, or the provider infrastructure,
which generally has coverage of the most common / high internet density
areas, and leave out the biggest portion of the globe.

In order to overcome both this limitations we build and will continue
building Kool Monkeys (though the name might change in the future).

The way Kool Monkeys tries to be a completely new approach to the Web
Monitoring problem is by being both Distributed and Open.

### Kool Monkey is Distributed!

The idea behind Kool Monkey is that any user will operate one or more agents.
The agent is free to download, and can run on Linux / OSX with no major
issues.

Once each agent will be registered under the account of the user operating it,
and its location documented, it will be available for performig any web performance
test agains any possible web page from that specific location.

The key advantage here is that in order to operate we don't need to set up
a huge infrastructure for hosting agents in various locations across the
world. Instead we just need to make the sevice useful for our users and
provide incentives for them to run one or more agents.

All the metrics collected by all the agents are stored in a central location,
and can be shared with all the users of the platform.

Using the distributed potential of Kool Monkey, you will be able to answer questions such as:

* How fast (or slow) is my web page from a user visiting it from "name a remote location here"?
* How well does my web page perform compared to my main competitors?
* If I were to open up my Web page in coutry X, how well will perform my current hosting options?

Without the need to invest a lot of capital in setting up dedicated infrastructure. The only thing
you will have to do will be to offer your agents to the community.

### Kool Monkey is Open!

Kool Monkey is both an Open Source project and an Open Service.
The source code for Kool Monkey will be always available though GitHub or
any alternative service in the future.

Additionally, we'll be operating an Open Service that will be open to the whole community
of people interested in Web Performance Measurements.

The way this is going to work is very simple.
Any user will have the opportunity to sign up for an account in Kool Monkeys.
Once the account is created, the user will be able to download the agent software for various
platforms (Currenty only binaries for `Linux` are available).

The user will have to set up the agent and register it under its own account. Once the agent will
be up and running, this will be immediately availabe for the whole Kool Monkeys community in order
to perform tests from that specific location, and at the same time the new user will be granted
access to the whole network of agents that the community is operating, getting immediate access
to a vast worldwide network of web monitoring agents.

### Long Term Vision

### Origins

The project started in the [Gopher Gala](http://www.gophergala.com/)
2016 (22-24 of January) based on particular needs of the participants.

We are currently doing the management on the project in
[trello](https://trello.com/b/zNxSafya), so feel free to check it out
and see what we have done and what we are planning to do.

## Business Model for the [SaaS version](here https://github.com/gophergala2016/kool_monkey)

## License

Kool Monkey is distributed under the terms of the GPLv3

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

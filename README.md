# Kool Monkey

Kool Monkey is an Open and Distributed Web monitoring system that collects metrics
related to web site perfomances, such as Response Time, Site Uptime and Network Roundtrip.

## Why another monitoring tool?

Today we observe a proliferation of web monitoring tools, and each
one of those is usually in one of the 2 following categories:

1. Proactive Monitoring
2. RUM (Real User Monitoring)

Both these types of solution tend to be expensive, and will typically lack
support for any but the most common cases.

With any of the existing commercial solutions you're then limited by
either the Pricing model, which makes it suitable for establised enterprise businesses,
or the provider infrastructure, which generally has coverage of the most common / high internet density
areas, and leave out the biggest portion of the globe.

The vast majority of the current providers will offer you agents in the Bay Area, but seldomly
you will be offered the ability to monitor from Chiapas, Hawaii or Siberia.

In order to overcome both these limitations we decided to build - and will continue
building - Kool Monkeys.

The way Kool Monkeys tries to be a completely new approach to the Web
Monitoring problem is by being both Distributed and Open.

### Kool Monkey is Distributed!

The idea behind Kool Monkey is that any user will operate one or more agents.
The agent is free to download, and can run on `Linux` and `OSX` with no major
issues.

Once each agent is registered under the account of the user operating it,
and its location documented, it will be available for performing any web performance
test against any possible web page from that specific location.

The key advantage here is that in order to operate the service at scale,
we don't need to set up an expensive infrastructure for hosting agents
in various locations across the world.
Instead we just need to make the sevice useful for our users and
provide incentives for them to run one or more agents.
More on that later.

All of the metrics collected by all the agents are stored in a central location,
and can be shared with all the users of the platform.

Using the distributed potential of Kool Monkey, you will be able to answer questions such as:

* How fast (or slow) is my web page for a user visiting it from `name a remote location here`?
* How well does my web page perform compared to my main competitors?
* If I were to open up my Web page in country X, how well will it perform with my current hosting options?

Without the need to invest a lot of capital in setting up dedicated infrastructure. The only thing
you will have to do will be to share your agents with the community.

### Kool Monkey is Open!

Kool Monkey is both an Open Source project and an Open Service.
The source code for Kool Monkey will be always available though GitHub or
any alternative service in the future.

Additionally, we'll be operating an Open Service which will soon be open to the whole community
of people interested in Web Performance Measurements and that requires high geographical
granularity without prohibitive service fees.

The way this is going to work is very simple:

1. Register a new account with Kool Monkeys
2. Download the agent for your favorite platform
3. Set up the agent and link it to your account
4. Get access to the - potentially - biggest network of web monitoring agents in the world!
5. The more agents you operate, the more `fair usage slots` you will get in order to use the global network

By using such a model, every user is encouraged to operate as many agents as possible, which will give them
access to more monitoring slots on the global platform. `What goes around, comes around` is a great way to
summarize the core concept of the service!

### Long Term Vision

The idea behind Kool Monkeys is to create the largest distributed network for Web Monitoring in the world.
By doing that we'll be able to open up various possibilities for the future:

* Be the largest database of the `State of the Internet` by providing up to date comprehensive coverage of
  how the current trends in Web performance are evolving
* Provide all users access to all data for free!
* Help in developing a culture of speed when it comes to serving web pages
* Apply predictive analysis in order to prevent or alert about possible congestions in various areas in the World
* And a lot more!

### Origins

The project started in the [Gopher Gala](http://www.gophergala.com/)
2016 (22-24 of January) based on particular needs of the participants.

We are currently doing the management on the project in
[trello](https://trello.com/b/zNxSafya), so feel free to check it out
and see what we have done and what we are planning to do.

## License

Kool Monkey is distributed under the terms of the GPLv3

## Live Demo

A live demo is availale at [this URL](http://dashboard.koolmonkey.xyz/).
It might be unavailable from time to time, and will be updated regularly.
Also, all the stored data is to be considered volatile.

## Architecture

This project has three important elements: the agent, the server and
the frontend.

### Agent

The agent is in charge of doing the actual measurements of the webpage. It's
developed in Go as a daemon that comunicates with the server in order
to obtain the urls to monitor with some parameters (when to start and
end the monitoring, the periodicity of the metrics, what metrics to
run, ...) and send back the actual metrics obtained.

### Server

The server is a central go application that manages the work between
the different agents, stores in db the results and comunicates with the
front end in order to obtain the user requests and show back the
results collected from the agents.

### Frontend

The frontend is where the user can request a new monitoring and see
the different information that has been collected in a graphical and
easy way.

## Build and deploy

As you have seen, the development of the project is in Go, so we have
provided some Makefiles in order to work with the code. These Makefiles
are not only useful to develop in the project, but to build rpms
packages from the source.

We have also integrated the project with Travis to run the tests,
build the rpms and deploy into production both the server part
(metrics collection and web frontend) and the agents, so it's become
quite easy to build and deploy the project.

The current status of travis is [![Build Status](https://travis-ci.org/gophergala2016/kool_monkey.svg?branch=master)](https://travis-ci.org/gophergala2016/kool_monkey) 

### Dependencies

In order to build and run the project you will require the following:

* Go compiler (tested with Go 1.5)
* PostgreSQL (tested with PostgreSQL 9.4)
* PhantomJS (tested with PhantomJS 1.9.7)

### Starting up a local environment

If you want to set up Kool Monkeys for local development, just clone
the repo and go through the following steps:

1. Build the project by calling:

    `make`

2. Start up the dev environment by calling:

    `make start-environment`

3. Start the Kool Monkeys server as follows:

    `./bin/kool-server -conf dev-env/conf/kool-server.conf`

4. Start the monitoring agent:

    `./bin/kool-agent`

## Developers

The main developers (and maintainers) of the project are:

* [Guillermo de los Santos](https://github.com/MemoDLSG)
* [Iván Californias](https://github.com/ivan-californias)
* [Pablo Álvarez de Sotomayor](https://github.com/i02sopop)
* [Sergio Visinoni](https://github.com/piffio)

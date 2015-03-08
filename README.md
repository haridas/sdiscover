sdiscover
=========

A consul cluster management utility for docker and other environments.

Me and Aneesh working on the service discovery issues. He need one tool which
help him to do the DevOps job easily to deploy a SOA based app. And for me it's
a learning process for my related SOA works using Zookeeper in Java :).

Use case for this tool,

We need to do the service discovery on our application which was deployed in
docker using fleet. Since the application was already developed, so changing
the code to achieving the service discovery is cumbersome, and at this point of
time it's unnecessary. So we my friend found this new cluster based tool consul,
and it fits for our needs.

But we require some tool written on top of consul to manage all consul related
operation and making all happens automatically using fleet and docker bootup
time. We started to develop an tool which does the following things for us,

- Automatically manage the consul cluster ( Servers and agents ).
- Failover, if one consul agent or master failed. Consul doesn't have this
  feature inbuilt at this point of time.
- Solves the issues while working with containers ( Auto restart, networking
  etc.)
- Easy integration with your application and less manual setup.

WIP


Links
=====
- consul.io

Contributers
============
- Haridas N <hn@haridas.in>
- Aneesh Kumar <me@aneesh.info>

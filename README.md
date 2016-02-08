sdiscover
=========

A consul cluster management utility for docker and other environments.

Service discovery is the one main use case of the Consul. Consul adopts
eventually consistent model for the cluster consistency, ie; it gives more
importance to the availability of the cluster on network partition times. The
Consul uses Quoram based leader election, with multi datacenter support. With
all this features inbuilt Consul is a good pick for our service discovery
requirement.

But for deploying the distributed application with consul cluster is not
a trivial task. We need a simple tool to rol eout the application with consul
integrated well with it.

Feature set:-

1. Agent Auto join with AWS S3 account to sync the cluster state.
2. Easy cluster bootstrap.
3. Integrate this tool with your micro containers for easy deployment.


WIP


Links
=====
- consul.io

Contributers
============
- Haridas N <hn@haridas.in>

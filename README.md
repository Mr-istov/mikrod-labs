# Workshop/crash course on microservices using docker/kubernetes

## Overview

1. Brief introduction about monolithic / microservice architecture.
2. Intro to docker. The what, how and why.
3. Intro to kubernetes. The what, how and why.

### Monolithic / microservice architecture

Let's imagine buidling a enterprise server-side application with the following components:

Monolithic approach architecture:

![mono-architecture](https://microservices.io/i/DecomposingApplications.011.jpg)

Microservice approach architecture:

![micro-architecture](https://microservices.io/i/Microservice_Architecture.png)

> NOTE: The images are not mine, they are from the following page and you can learn more information about the architectures here: [Microservices.io](https://microservices.io/patterns/microservices.html)

### Docker

Docker is a set of platform as a service products that uses OS-level virtualization to deliver software in packages called containers. Containers are isolated from one another and bundle their own software, libraries and configuration files; they can communicate with each other through well-defined channels.

Containerized applications:

![container-image](https://www.docker.com/sites/default/files/d8/2018-11/docker-containerized-appliction-blue-border_2.png)

Application on VM:

![vm-image](https://www.docker.com/sites/default/files/d8/2018-11/container-vm-whatcontainer_2.png)

Combined:

![combined-image](https://i1.wp.com/www.docker.com/blog/wp-content/uploads/Are-containers-..-vms-image-2-1024x759.png?ssl=1)

# Workshop/crash course on microservices using docker/kubernetes

## Prerequisites

1. [docker](https://www.docker.com/get-started)
2. [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
3. [minikube](https://minikube.sigs.k8s.io/docs/)

## Overview

1. Brief introduction about monolithic / microservice architecture.
2. Intro to docker. The what, how and why.
3. Intro to kubernetes. The what, how and why.

### Monolithic / microservice architecture

Let's imagine buidling a enterprise server-side application with the following components:

_Monolithic approach architecture:_

![mono-architecture](https://microservices.io/i/DecomposingApplications.011.jpg)

_Microservice approach architecture:_

![micro-architecture](https://microservices.io/i/Microservice_Architecture.png)

> NOTE: The images are not mine, they are from the following page and you can learn more information about the architectures here: [Microservices.io](https://microservices.io/patterns/microservices.html)

### Docker

Docker is a set of platform as a service products that uses OS-level virtualization to deliver software in packages called containers. Containers are isolated from one another and bundle their own software, libraries and configuration files; they can communicate with each other through well-defined channels.

_Containerized applications:_

![container-image](https://www.docker.com/sites/default/files/d8/2018-11/docker-containerized-appliction-blue-border_2.png)

_Application on VM:_

![vm-image](https://www.docker.com/sites/default/files/d8/2018-11/container-vm-whatcontainer_2.png)

_Combined:_

![combined-image](https://i1.wp.com/www.docker.com/blog/wp-content/uploads/Are-containers-..-vms-image-2-1024x759.png?ssl=1)

Explain the following _docker concepts_:

- Image (How to build)
- Container (How to run)

### Kubernetes

Kubernetes is an open-source container-orchestration system for automating application deployment, scaling, and management.

_k8s architecture:_

![k8s-image](https://www.mediakwest.com/images/galeries/_2019_Galerie/2019_09_28_MK33_SERVICES_CLOUD/MKW33_OutilsCloud_kubernetes_architecture.jpg)

Explain the following _k8s concepts:_

- Pod
- Deployment
- Service

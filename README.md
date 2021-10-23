# Go microservices deployed to Kubernetes with Helm subcharts

## Purpose

Showcase good practices when developing for Kubernetes

Topics
- Helm chart with subcharts
- Multi-stage Dockerfile for build, testing, and deployment
- Fast developer workflow

## Prerequisites

- Install Docker
    - see https://docs.docker.com/get-docker/

- Install kubectl
    - On macOS via Homebrew: `brew install kubectl`
    - Others see https://kubernetes.io/docs/tasks/tools/

- Install minikube
    - On macOS via Homebrew: `brew install minikube`
    - Others see https://minikube.sigs.k8s.io/docs/start/

- Install Helm
    - On macOS via Homebrew: `brew install helm`
    - Others see https://helm.sh/docs/intro/install/

- Install Tilt
    - On macOS via Homebrew: `brew install tilt-dev/tap/tilt`
    - Others see https://docs.tilt.dev/install.html

- Install ctlptl
  - On macOS via Homebrew: `brew install tilt-dev/tap/ctlptl`
  - Others see https://github.com/kubernetes/examples/blob/master/guidelines.md

## Development

### Create cluster

create
```shell
ctlptl create cluster minikube --registry=ctlptl-registry --kubernetes-version=v1.22.0
```

delete
```shell
ctlptl delete cluster minikube
```

### Start services

```shell
tilt up
```

## References

### Helm
https://helm.sh/docs/chart_template_guide/subcharts_and_globals/  
https://faun.pub/helm-chart-how-to-create-helm-charts-from-kubernetes-k8s-yaml-from-scratch-d64901e36850  
https://github.com/kubernetes/examples/blob/master/guidelines.md  

### Go
https://github.com/powerman/go-monolith-example  
https://github.com/getkin/kin-openapi  

### Docker
https://docs.docker.com/develop/develop-images/multistage-build/  
https://medium.com/@lizrice/non-privileged-containers-based-on-the-scratch-image-a80105d6d341  

## TODO
- add ingress
- add more services to support popular pet store example 

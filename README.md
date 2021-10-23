# Example of Go microservice with a Helm chart

## TODO
add ingress
add more services
add attributes service
add key access service

## Prerequisites

- Install Docker
    - see https://docs.docker.com/get-docker/

- Install kubectl
    - On macOs via Homebrew: `brew install kubectl`
    - Others see https://kubernetes.io/docs/tasks/tools/

- Install kind
    - On macOS via Homebrew: `brew install kind`
    - Others see https://kind.sigs.k8s.io/docs/user/quick-start/#installation

- Install helm
    - On macOS via Homebrew: `brew install helm`
    - Others see https://helm.sh/docs/intro/install/

- Install Tilt
    - On macOS via Homebrew: `brew install tilt-dev/tap/tilt`
    - Others see https://docs.tilt.dev/install.html

## Development

### Create cluster

`kind create cluster --name ex`

create
```shell
ctlptl create cluster minikube --registry=ctlptl-registry --kubernetes-version=v1.22.0
```

delete
```shell
ctlptl delete cluster minikube
```

### tilt
```shell
tilt --debug --verbose up
```

### Start services

```shell
tilt up
```

## References
https://github.com/powerman/go-monolith-example
https://github.com/getkin/kin-openapi
https://faun.pub/helm-chart-how-to-create-helm-charts-from-kubernetes-k8s-yaml-from-scratch-d64901e36850
https://medium.com/@lizrice/non-privileged-containers-based-on-the-scratch-image-a80105d6d341

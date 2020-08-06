## Simple Tilt Demo

### Prerequisites

- [Tilt](https://tilt.dev)
- Kubernetes cluster

### Demo

1. Run `tilt up` + press space to open the browser
1. create a `Tiltfile` from the console
1. add `k8s_yaml('deploy.yaml')`
1. add `docker_build('tilt-demo', '.')
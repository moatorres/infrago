# Infrago

### A Static `FileServer` example application written in Go

Infrago is an open-source static file server application based on the native `net/http` module. This project was created as an exercise and should serve as a simple example of what you can build with [Go](https://go.dev/), [Docker](https://www.docker.com/) and [Kubernetes](https://kubernetes.io/).


https://github.com/moatorres/infrago/assets/44585769/8109d758-6215-48a6-b888-d1c1cec94a9b


## Build

#### Go Binary

To build the Go binary, ensure you're on the root directory of the project and run:

```sh
go build -o infrago
```

#### Docker Image

You can either build a local Docker Image manually or using the available [make](https://www.gnu.org/software/make/manual/make.html) commands.

**Buildind a Docker Image with `docker buildx`**

```sh
docker buildx build -f Dockerfile . -t docker.io/moatorres/infrago
```

**Building a Docker Image with `make`**

```sh
make build
```

## Deploy

<sub>This project is skaffold-friendly ✨ If you don't want to install `skaffold` on your machine, you can deploy the resources as you'd normally do with `kubectl`. Learn more about Skaffold [here](https://skaffold.dev/). </sub>

**Deploying with `kubectl`**

```sh
kubectl apply -f k8s/dev/*.yaml
```

**Deploying with `make`**

```sh
make deploy-dev
```

**Deploying with `skaffold`**

```sh
skaffold dev
```

## Graceful Shutdown

Let's see how our deployment behaves on Kubernetes. It's advised that you install `skaffold`. Skaffold spins resources up and down without requiring shell scripts. ✨

1. Open **three** terminals on your machine then run `skaffold dev` in the first one.
2. Run `kubectl get pods` to get the names of the running pods.
3. Now run `kubeclt logs -f pods/infrago-7f5dbb59cf-dvx8g` (remember to replace the pod name with yours). This will allow us to follow our pod's logs.
4. Go to the third terminal and run `kubectl get pods infrago-7f5dbb59cf-dvx8g -w` to watch the pod's state.
5. Now go back to the first terminal and press `ctrl c` to stop the `skaffold` process.

See what happened? Since we specified a `terminationGracePeriodSeconds: 120` on our Deployment, our server was able to catch the `'terminated'` signal, wait for 10 seconds then exit the process before being shut down by Kubernetes.

## Related Projects

- [Loggo](https://github.com/moatorres/go/modules/logger) is a zero-dependency JSON-based logger written in Go

<sub>⚡️ Powered by **OSS** — `< >` with ☕️ by [**Moa Torres**](https://github.com/moatorres)</sub>

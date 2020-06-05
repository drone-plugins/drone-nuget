# drone-nuget

[![Build Status](http://cloud.drone.io/api/badges/drone-plugins/drone-nuget/status.svg)](http://cloud.drone.io/drone-plugins/drone-nuget)
[![Gitter chat](https://badges.gitter.im/drone/drone.png)](https://gitter.im/drone/drone)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![](https://images.microbadger.com/badges/image/plugins/nuget.svg)](https://microbadger.com/images/plugins/nuget "Get your own image badge on microbadger.com")
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-nuget?status.svg)](http://godoc.org/github.com/drone-plugins/drone-nuget)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-nuget)](https://goreportcard.com/report/github.com/drone-plugins/drone-nuget)

Drone plugin to publish files and artifacts to a NuGet repository. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Build

Build the binary with the following command:

```console
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -v -a -tags netgo -o release/linux/amd64/drone-nuget
```

## Docker

Build the Docker image with the following command:

```console
docker build \
  --label org.label-schema.build-date=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
  --label org.label-schema.vcs-ref=$(git rev-parse --short HEAD) \
  --file docker/Dockerfile.linux.amd64 --tag plugins/nuget .
```

## Usage

```console
docker run --rm \
  -e PLUGIN_SOURCE=http://nuget.company.com \
  -e PLUGIN_API_KEY=SUPER_KEY \
  -e PLUGIN_FILES=*.nupkg \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/nuget
```

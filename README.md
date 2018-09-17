# Prometheus HetznerCloud SD

[![Build Status](http://github.dronehippie.de/api/badges/promhippie/prometheus-hcloud-sd/status.svg)](http://github.dronehippie.de/promhippie/prometheus-hcloud-sd)
[![Build status](https://ci.appveyor.com/api/projects/status/k96mjylckmsb89bn?svg=true)](https://ci.appveyor.com/project/tboerger/prometheus-hcloud-sd)
[![Stories in Ready](https://badge.waffle.io/promhippie/prometheus-hcloud-sd.svg?label=ready&title=Ready)](http://waffle.io/promhippie/prometheus-hcloud-sd)
[![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/d7900c4c246740edb77cf29a4b1d85ee)](https://www.codacy.com/app/promhippie/prometheus-hcloud-sd?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=promhippie/prometheus-hcloud-sd&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/promhippie/prometheus-hcloud-sd?status.svg)](http://godoc.org/github.com/promhippie/prometheus-hcloud-sd)
[![Go Report](http://goreportcard.com/badge/github.com/promhippie/prometheus-hcloud-sd)](http://goreportcard.com/report/github.com/promhippie/prometheus-hcloud-sd)
[![](https://images.microbadger.com/badges/image/promhippie/prometheus-hcloud-sd.svg)](http://microbadger.com/images/promhippie/prometheus-hcloud-sd "Get your own image badge on microbadger.com")

This project provides a server to automatically discover nodes within your HetznerCloud account in a Prometheus SD compatible format.


## Install

You can download prebuilt binaries from our [GitHub releases](https://github.com/promhippie/prometheus-hcloud-sd/releases), or you can use our Docker images published on [Docker Hub](https://hub.docker.com/r/promhippie/prometheus-hcloud-sd/tags/).


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.8.

```bash
go get -d github.com/promhippie/prometheus-hcloud-sd
cd $GOPATH/src/github.com/promhippie/prometheus-hcloud-sd

# install retool
make retool

# sync dependencies
make sync

# generate code
make generate

# build binary
make build

./bin/prometheus-hcloud-sd -h
```


## Security

If you find a security issue please contact thomas@webhippie.de first.


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

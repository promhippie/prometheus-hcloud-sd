# Prometheus Hetzner Cloud SD

[![Current Tag](https://img.shields.io/github/v/tag/promhippie/prometheus-hcloud-sd?sort=semver)](https://github.com/promhippie/prometheus-scw-sd) [![General Build](https://github.com/promhippie/prometheus-hcloud-sd/actions/workflows/general.yml/badge.svg)](https://github.com/promhippie/prometheus-hcloud-sd/actions/workflows/general.yaml) [![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/d7900c4c246740edb77cf29a4b1d85ee)](https://www.codacy.com/gh/promhippie/prometheus-hcloud-sd/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=promhippie/prometheus-hcloud-sd&amp;utm_campaign=Badge_Grade) [![Go Doc](https://godoc.org/github.com/promhippie/prometheus-hcloud-sd?status.svg)](http://godoc.org/github.com/promhippie/prometheus-hcloud-sd) [![Go Report](http://goreportcard.com/badge/github.com/promhippie/prometheus-hcloud-sd)](http://goreportcard.com/report/github.com/promhippie/prometheus-hcloud-sd)

This project provides a server to automatically discover nodes within your
Hetzner Cloud account in a Prometheus SD compatible format.

## Install

You can download prebuilt binaries from our [GitHub releases][releases], or you
can use our containers published on [Docker Hub][dockerhub] and [Quay][quayio].
If you need further guidance how to install this take a look at our
[documentation][docs].

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions][golang]. This project requires
Go >= v1.17, at least that's the version we are using.

```console
git clone https://github.com/promhippie/prometheus-hcloud-sd.git
cd prometheus-hcloud-sd

make generate build

./bin/prometheus-hcloud-sd -h
```

## Security

If you find a security issue please contact
[thomas@webhippie.de](mailto:thomas@webhippie.de) first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

[releases]: https://github.com/promhippie/prometheus-hcloud-sd/releases
[dockerhub]: https://hub.docker.com/r/promhippie/prometheus-hcloud-sd/tags/
[quayio]: https://quay.io/repository/promhippie/prometheus-hcloud-sd?tab=tags
[docs]: https://promhippie.github.io/prometheus-hcloud-sd/#getting-started
[golang]: http://golang.org/doc/install.html

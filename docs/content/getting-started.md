---
title: "Getting Started"
date: 2018-05-02T00:00:00+00:00
anchor: "getting-started"
weight: 10
---

## Installation

We won't cover further details how to properly setup [Prometheus](https://prometheus.io) itself, we will only cover some basic setup based on [docker-compose](https://docs.docker.com/compose/). But if you want to run this service discovery without [docker-compose](https://docs.docker.com/compose/) you should be able to adopt that to your needs.

First of all we need to prepare a configuration for [Prometheus](https://prometheus.io) that includes the service discovery which simply maps to a node exporter.

{{< gist tboerger b9c39b6571f48ce2b132de1531061531 "prometheus.yml" >}}

After preparing the configuration we need to create the `docker-compose.yml` within the same folder, this `docker-compose.yml` starts a simple [Prometheus](https://prometheus.io) instance together with the service discovery. Don't forget to update the envrionment variables with the required credentials. If you are using a different volume for the service discovery you have to make sure that the container user is allowed to write to this volume.

{{< gist tboerger b9c39b6571f48ce2b132de1531061531 "docker-compose.yml" >}}

Since our `latest` Docker tag always refers to the `master` branch of the Git repository you should always use some fixed version. You can see all available tags at our [DockerHub repository](https://hub.docker.com/r/promhippie/prometheus-hcloud-sd/tags/), there you will see that we also provide a manifest, you can easily start the exporter on various architectures without any change to the image name. You should apply a change like this to the `docker-compose.yml`:

{{< gist tboerger b9c39b6571f48ce2b132de1531061531 "tag.diff" >}}

Finally the service discovery should be configured fine, let's start this stack with [docker-compose](https://docs.docker.com/compose/), you just need to execute `docker-compose up` within the directory where you have stored `prometheus.yml` and `docker-compose.yml`.

{{< gist tboerger b9c39b6571f48ce2b132de1531061531 "output.log" >}}

That's all, the service discovery should be up and running. You can access [Prometheus](https://prometheus.io) at [http://localhost:9090](http://localhost:9090).

{{< figure src="service-discovery.png" title="Prometheus service discovery for HetznerCloud" >}}

## Kubernetes

Currently we have not prepared a deployment for Kubernetes, but this is something we will provide for sure. Most interesting will be the integration into the [Prometheus Operator](https://coreos.com/operators/prometheus/docs/latest/), so stay tuned.

## Configuration

PROMETHEUS_HCLOUD_TOKEN
: Access token for the HetznerCloud API, required for authentication

PROMETHEUS_HCLOUD_LOG_LEVEL
: Only log messages with given severity, defaults to `info`

PROMETHEUS_HCLOUD_LOG_PRETTY
: Enable pretty messages for logging, defaults to `true`

PROMETHEUS_HCLOUD_WEB_ADDRESS
: Address to bind the metrics server, defaults to `0.0.0.0:9000`

PROMETHEUS_HCLOUD_WEB_PATH
: Path to bind the metrics server, defaults to `/metrics`

PROMETHEUS_HCLOUD_OUTPUT_FILE
: Path to write the file_sd config, defaults to `/etc/prometheus/hcloud.json`

PROMETHEUS_HCLOUD_OUTPUT_REFRESH
: Discovery refresh interval in seconds, defaults to `30`

## Labels

* `__meta_hcloud_name`
* `__meta_hcloud_status`
* `__meta_hcloud_public_ipv4`
* `__meta_hcloud_public_ipv6`
* `__meta_hcloud_type`
* `__meta_hcloud_cores`
* `__meta_hcloud_memory`
* `__meta_hcloud_disk`
* `__meta_hcloud_storage`
* `__meta_hcloud_cpu`
* `__meta_hcloud_datacenter`
* `__meta_hcloud_location`
* `__meta_hcloud_city`
* `__meta_hcloud_country`
* `__meta_hcloud_image_name`
* `__meta_hcloud_image_type`
* `__meta_hcloud_os_flavor`
* `__meta_hcloud_os_version`

## Metrics

prometheus_hcloud_sd_request_duration_seconds
: Histogram of latencies for requests to the HetznerCloud API

prometheus_hcloud_sd_request_failures_total
: Total number of failed requests to the HetznerCloud API

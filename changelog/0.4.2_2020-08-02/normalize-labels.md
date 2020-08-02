Bugfix: Normalize specialchars in labels

Since you can use dots and dashes as part of your labels within Hetzner Cloud we
started to replace these characters by an underscore now, otherwise this will
result in invalid labels for Prometheus.

https://github.com/promhippie/prometheus-hcloud-sd/issues/16

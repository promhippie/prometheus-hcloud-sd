Bugfix: Define only existing image labels

It's possible that a server doesn't provide a image label, so we are setting the
right labels ony with a value if this is really available.

https://github.com/promhippie/prometheus-hcloud-sd/pull/3

# Changelog for unreleased

The following sections list the changes for unreleased.

## Summary

 * Chg #334: Switch to official logging library

## Details

 * Change #334: Switch to official logging library

   Since there have been a structured logger part of the Go standard library we
   thought it's time to replace the library with that. Be aware that log messages
   should change a little bit.

   https://github.com/promhippie/prometheus-hcloud-sd/issues/334


# Changelog for 1.1.0

The following sections list the changes for 1.1.0.

## Summary

 * Enh #261: Add private IPs as labels

## Details

 * Enhancement #261: Add private IPs as labels

   We have added a list of all attached private networks per server, that way you
   are now also able to connect optionally through the private address instead of
   using the public address.

   https://github.com/promhippie/prometheus-hcloud-sd/pull/261


# Changelog for 1.0.0

The following sections list the changes for 1.0.0.

## Summary

 * Chg #252: Read secrets form files
 * Enh #252: Update all releated dependencies

## Details

 * Change #252: Read secrets form files

   We have added proper support to load secrets like tokens from files or from
   base64-encoded strings. Just provide the flags or environment variables with a
   DSN formatted string like `file://path/to/file` or `base64://Zm9vYmFy`.

   https://github.com/promhippie/prometheus-hcloud-sd/issues/252

 * Enhancement #252: Update all releated dependencies

   We've updated all dependencies to the latest available versions, including more
   current versions of build tools and used Go version to build the binaries. It's
   time to mark a stable release.

   https://github.com/promhippie/prometheus-hcloud-sd/issues/252


# Changelog for 0.6.0

The following sections list the changes for 0.6.0.

## Summary

 * Enh #149: Use GitHub Actions onstead of Drone CI
 * Enh #149: Improve doucmentation and repo structure

## Details

 * Enhancement #149: Use GitHub Actions onstead of Drone CI

   We have replaced the previous Drone CI setup by more simple GitHub Actions since
   are anyway using GitHub for the code hosting and issue tracking. As part of that
   we are now also publishing the docker images to Quay.

   https://github.com/promhippie/prometheus-hcloud-sd/pull/149

 * Enhancement #149: Improve doucmentation and repo structure

   We have improved the available documentation pretty hard and we also added
   documentation how to install this service discovery via Helm or Kustomize on
   Kubernetes. Beside that we are testing to build the bundled Kustomize manifests
   now.

   https://github.com/promhippie/prometheus-hcloud-sd/pull/149


# Changelog for 0.5.0

The following sections list the changes for 0.5.0.

## Summary

 * Chg #19: Use bingo for development tooling
 * Chg #21: Drop dariwn/386 release builds
 * Chg #44: Update Go version and dependencies
 * Chg #44: Improvements for automated documentation
 * Chg #45: Integrate new HTTP service discovery handler
 * Chg #46: Integrate standard web config

## Details

 * Change #19: Use bingo for development tooling

   We switched to use [bingo](github.com/bwplotka/bingo) for fetching development
   and build tools based on fixed defined versions to reduce the dependencies
   listed within the regular go.mod file within this project.

   https://github.com/promhippie/prometheus-hcloud-sd/issues/19

 * Change #21: Drop dariwn/386 release builds

   We dropped the build of 386 builds on Darwin as this architecture is not
   supported by current Go versions anymore.

   https://github.com/promhippie/prometheus-hcloud-sd/issues/21

 * Change #44: Update Go version and dependencies

   We updated the Go version used to build the binaries within the CI system and
   beside that in the same step we have updated all dependencies ti keep everything
   up to date.

   https://github.com/promhippie/prometheus-hcloud-sd/pull/44

 * Change #44: Improvements for automated documentation

   We have added some simple scripts that gets executed by Drone to keep moving
   documentation parts like the available labels or the available environment
   variables always up to date. No need to update the docs related to that manually
   anymore.

   https://github.com/promhippie/prometheus-hcloud-sd/pull/44

 * Change #45: Integrate new HTTP service discovery handler

   We integrated the new HTTP service discovery which have been introduced by
   Prometheus starting with version 2.28. With this new service discovery you can
   deploy this service whereever you want and you are not tied to the Prometheus
   filesystem anymore.

   https://github.com/promhippie/prometheus-hcloud-sd/issues/45

 * Change #46: Integrate standard web config

   We integrated the new web config from the Prometheus toolkit which provides a
   configuration for TLS support and also some basic builtin authentication. For
   the detailed configuration you check out the documentation.

   https://github.com/promhippie/prometheus-hcloud-sd/issues/46


# Changelog for 0.4.2

The following sections list the changes for 0.4.2.

## Summary

 * Fix #16: Normalize specialchars in labels

## Details

 * Bugfix #16: Normalize specialchars in labels

   Since you can use dots and dashes as part of your labels within Hetzner Cloud we
   started to replace these characters by an underscore now, otherwise this will
   result in invalid labels for Prometheus.

   https://github.com/promhippie/prometheus-hcloud-sd/issues/16


# Changelog for 0.4.1

The following sections list the changes for 0.4.1.

## Summary

 * Fix #14: Binaries are not static linked

## Details

 * Bugfix #14: Binaries are not static linked

   We fixed building properly static linked binaries, since the last release and a
   major refactoring of the binaries and the CI pipeline we introduced binaries
   which had been linked to muslc by mistake. With this change applied all binaries
   will be properly static linked again.

   https://github.com/promhippie/prometheus-hcloud-sd/issues/14


# Changelog for 0.4.0

The following sections list the changes for 0.4.0.

## Summary

 * Chg #13: Code and project restructuring

## Details

 * Change #13: Code and project restructuring

   To get the project and code structure into a new shape and to get it cleaned up
   we switched to Go modules and restructured the project source in general. The
   functionality stays the same as before.

   https://github.com/promhippie/prometheus-hcloud-sd/pull/13


# Changelog for 0.3.0

The following sections list the changes for 0.3.0.

## Summary

 * Chg #4: Switch to cloud.drone.io
 * Chg #5: Add support for server labels
 * Chg #6: Support for multiple accounts
 * Chg #9: Define healthcheck command

## Details

 * Change #4: Switch to cloud.drone.io

   We don't wanted to maintain our own Drone infrastructure anymore, since there is
   cloud.drone.io available for free we switched the pipelines over to it.

   https://github.com/promhippie/prometheus-hcloud-sd/pull/4

 * Change #5: Add support for server labels

   Since Hetzner Cloud introduced labels for servers we should also map these
   labels to the exported JSON file.

   https://github.com/promhippie/prometheus-hcloud-sd/pull/5

 * Change #6: Support for multiple accounts

   Make the deployments of this service discovery easier, previously we had to
   launch one instance for every credentials we wanted to gather, with this change
   we are able to define multiple credentials for a single instance of the service
   discovery.

   https://github.com/promhippie/prometheus-hcloud-sd/pull/6

 * Change #9: Define healthcheck command

   To check the health status of the service discovery especially within Docker we
   added a simple subcommand which checks the healthz endpoint to show if the
   service is up and running.

   https://github.com/promhippie/prometheus-hcloud-sd/pull/9


# Changelog for 0.2.0

The following sections list the changes for 0.2.0.

## Summary

 * Fix #3: Define only existing image labels
 * Chg #1: Add basic documentation
 * Chg #2: Pin xgo to golang 1.10 to avoid issues
 * Chg #3: Update dependencies
 * Chg #3: Timeout for metrics handler
 * Chg #3: Panic recover within handlers

## Details

 * Bugfix #3: Define only existing image labels

   It's possible that a server doesn't provide a image label, so we are setting the
   right labels ony with a value if this is really available.

   https://github.com/promhippie/prometheus-hcloud-sd/pull/3

 * Change #1: Add basic documentation

   Add some basic documentation page which also includes build and installation
   instructions to make clear how this project can be installed and used.

   https://github.com/promhippie/prometheus-hcloud-sd/pull/1

 * Change #2: Pin xgo to golang 1.10 to avoid issues

   There had been issues while using the latest xgo version, let's pin this tag to
   1.10 to ensure the binaries are properly build.

   https://github.com/promhippie/prometheus-hcloud-sd/pull/2

 * Change #3: Update dependencies

   Just make sure to update all the build dependencies to work with the latest
   versions available.

   https://github.com/promhippie/prometheus-hcloud-sd/pull/3

 * Change #3: Timeout for metrics handler

   We added an additional middleware to properly timeout requests to the metrics
   endpoint for long running request.

   https://github.com/promhippie/prometheus-hcloud-sd/pull/3

 * Change #3: Panic recover within handlers

   To make sure panics are properly handled we added a middleware to recover
   properly from panics.

   https://github.com/promhippie/prometheus-hcloud-sd/pull/3


# Changelog for 0.1.0

The following sections list the changes for 0.1.0.

## Summary

 * Chg #12: Initial release of basic version

## Details

 * Change #12: Initial release of basic version

   Just prepared an initial basic version which could be released to the public.

   https://github.com/promhippie/prometheus-hcloud-sd/issues/12



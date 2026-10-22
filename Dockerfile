FROM --platform=$BUILDPLATFORM golang:1.27.3-alpine@sha256:f2c872bcee949b7be9c490efd02b48a8a2fbf947cdf533f1c58525ec9d8170fd AS builder

RUN apk add --no-cache -U git curl
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

WORKDIR /go/src/prometheus-hcloud-sd
COPY . /go/src/prometheus-hcloud-sd/

RUN --mount=type=cache,target=/go/pkg \
    go mod download -x

ARG TARGETOS
ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    task generate build GOOS=${TARGETOS} GOARCH=${TARGETARCH}

FROM alpine:3.24@sha256:c6b073f8e9e58d881f6ffba6ebf3b9ebd60b9e7d6e8337018c0e68bb5c356873

RUN apk add --no-cache ca-certificates mailcap && \
    addgroup -g 1337 prometheus-hcloud-sd && \
    adduser -D -u 1337 -h /var/lib/prometheus-hcloud-sd -G prometheus-hcloud-sd prometheus-hcloud-sd

EXPOSE 9000
VOLUME ["/var/lib/prometheus-hcloud-sd"]
ENTRYPOINT ["/usr/bin/prometheus-hcloud-sd"]
CMD ["server"]
HEALTHCHECK CMD ["/usr/bin/prometheus-hcloud-sd", "health"]

ENV PROMETHEUS_HCLOUD_OUTPUT_ENGINE="http"
ENV PROMETHEUS_HCLOUD_OUTPUT_FILE="/var/lib/prometheus-hcloud-sd/output.json"

COPY --from=builder /go/src/prometheus-hcloud-sd/bin/prometheus-hcloud-sd /usr/bin/prometheus-hcloud-sd
WORKDIR /var/lib/prometheus-hcloud-sd
USER prometheus-hcloud-sd

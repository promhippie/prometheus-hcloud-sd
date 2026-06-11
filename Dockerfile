FROM --platform=$BUILDPLATFORM golang:1.26.4-alpine@sha256:7a3e50096189ad57c9f9f865e7e4aa8585ed1585248513dc5cda498e2f41812c AS builder

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

FROM alpine:3.24@sha256:a2d49ea686c2adfe3c992e47dc3b5e7fa6e6b5055609400dc2acaeb241c829f4

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

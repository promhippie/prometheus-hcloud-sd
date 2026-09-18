FROM --platform=$BUILDPLATFORM golang:1.27.1-alpine@sha256:e9bbdf282b51ac8b34c46e5f31d2d56e7bad60366c35f08d2f295b921b13388b AS builder

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

FROM alpine:3.24@sha256:5b02b42e375f7426f8d65c3af331ca05d9878f9989230354504e0b9dfd431f60

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

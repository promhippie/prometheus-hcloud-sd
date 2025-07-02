FROM --platform=$BUILDPLATFORM golang:1.24.4-alpine3.21@sha256:56a23791af0f77c87b049230ead03bd8c3ad41683415ea4595e84ce7eada121a AS builder

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

FROM alpine:3.22@sha256:8a1f59ffb675680d47db6337b49d22281a139e9d709335b492be023728e11715

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

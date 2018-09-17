FROM webhippie/alpine:latest AS build
RUN apk add --no-cache ca-certificates mailcap

FROM scratch

LABEL maintainer="Thomas Boerger <thomas@webhippie.de>" \
  org.label-schema.name="Prometheus HetznerCloud SD" \
  org.label-schema.vendor="Thomas Boerger" \
  org.label-schema.schema-version="1.0"

ENTRYPOINT ["/usr/bin/prometheus-hcloud-sd"]
CMD ["server"]

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/mime.types /etc/

COPY dist/binaries/prometheus-hcloud-sd-*-linux-arm-6 /usr/bin/prometheus-hcloud-sd

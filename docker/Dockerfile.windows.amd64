# escape=`
FROM microsoft/nanoserver:10.0.14393.2430

LABEL maintainer="Thomas Boerger <thomas@webhippie.de>" `
  org.label-schema.name="Prometheus HetznerCloud SD" `
  org.label-schema.vendor="Thomas Boerger" `
  org.label-schema.schema-version="1.0"

ENTRYPOINT ["c:\\prometheus-hcloud-sd.exe"]
CMD ["server"]

COPY bin/prometheus-hcloud-sd.exe c:\prometheus-hcloud-sd.exe

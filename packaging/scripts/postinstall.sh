#!/bin/sh
set -e

chown -R prometheus-hcloud-sd:prometheus-hcloud-sd /etc/prometheus-hcloud-sd
chown -R prometheus-hcloud-sd:prometheus-hcloud-sd /var/lib/prometheus-hcloud-sd
chmod 750 /var/lib/prometheus-hcloud-sd

if [ -d /run/systemd/system ]; then
    systemctl daemon-reload

    if systemctl is-enabled --quiet prometheus-hcloud-sd.service; then
        systemctl restart prometheus-hcloud-sd.service
    fi
fi

#!/bin/sh
set -e

if ! getent group prometheus-hcloud-sd >/dev/null 2>&1; then
    groupadd --system prometheus-hcloud-sd
fi

if ! getent passwd prometheus-hcloud-sd >/dev/null 2>&1; then
    useradd --system --create-home --home-dir /var/lib/prometheus-hcloud-sd --shell /bin/bash -g prometheus-hcloud-sd prometheus-hcloud-sd
fi

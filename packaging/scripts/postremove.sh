#!/bin/sh
set -e

if [ ! -d /var/lib/prometheus-hcloud-sd ] && [ ! -d /etc/prometheus-hcloud-sd ]; then
    userdel prometheus-hcloud-sd 2>/dev/null || true
    groupdel prometheus-hcloud-sd 2>/dev/null || true
fi

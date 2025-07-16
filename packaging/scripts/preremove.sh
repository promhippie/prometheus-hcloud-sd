#!/bin/sh
set -e

systemctl stop prometheus-hcloud-sd.service || true
systemctl disable prometheus-hcloud-sd.service || true

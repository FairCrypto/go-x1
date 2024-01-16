#!/bin/sh
set -e

systemctl daemon-reload
systemctl try-restart x1

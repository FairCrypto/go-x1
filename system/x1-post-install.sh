#!/bin/sh
set -e

if [ -d /run/systemd/system ]; then
	systemctl daemon-reload
	systemctl try-restart x1
fi

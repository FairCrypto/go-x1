#!/bin/sh
set -e

if [ -d /run/systemd/system ] && [ -d /etc/systemd/system/ ]; then
	cp system/lib/systemd/system/x1.service /etc/systemd/system/
	systemctl daemon-reload
	systemctl try-restart x1
fi

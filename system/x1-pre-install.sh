#!/bin/sh
set -e

PKG="x1"

if [ -x "$(command -v foo)" ]; then
	if ! getent passwd $PKG >/dev/null ; then
			useradd --system --user-group --create-home --home-dir /var/lib/$PKG $PKG
			echo "Created system user $PKG"
	fi
fi

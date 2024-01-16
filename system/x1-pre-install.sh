#!/bin/sh
set -e

PKG="x1"

if ! getent passwd $PKG >/dev/null ; then
    adduser --shell /usr/sbin/nologin --quiet --system --home /var/lib/$PKG $PKG
    echo "Created system user $PKG"
fi

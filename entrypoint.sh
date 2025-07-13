#!/bin/sh
set -e

# Default service is api
SERVICE=${SERVICE:-api}
exec "/usr/local/bin/$SERVICE"

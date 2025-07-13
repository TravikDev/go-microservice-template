#!/bin/bash
set -euo pipefail

# Lint entrypoint script
shellcheck entrypoint.sh

# Lint all shell scripts under scripts/
if ls scripts/*.sh >/dev/null 2>&1; then
    shellcheck scripts/*.sh
fi

#!/bin/bash
set -euo pipefail

if [ -f go.mod ]; then
    fmt_out="$(find . -name '*.go' ! -path './vendor/*' -print0 | xargs -0 gofmt -l)"
    if [ -n "$fmt_out" ]; then
        echo "The following files are not gofmt'd:" >&2
        echo "$fmt_out" >&2
        exit 1
    fi
    go vet ./...
    go test ./...
else
    echo "No Go module found; skipping Go checks."
fi

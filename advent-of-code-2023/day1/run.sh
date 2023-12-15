#!/usr/bin/env bash

set -euo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
cd "$script_dir"

exec 3>&1
{
    echo '```'
    echo '$ ./run.sh'
    go test -bench=. -benchmem "$@" . | tee /dev/fd/3
    echo '```'
} > README.md

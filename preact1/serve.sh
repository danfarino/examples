#!/bin/bash

pnpm install
container_id="$(docker run --rm -d  -v "$PWD:/srv" -p 8888:8888 caddy:alpine caddy file-server --listen :8888)"
trap 'docker kill $container_id' EXIT
pnpm run build --watch

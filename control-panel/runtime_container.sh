#!/usr/bin/env bash

CONTAINER_RUNTIME=$(command -v podman &> /dev/null && echo podman || echo docker)

${CONTAINER_RUNTIME} "$@"
#!/usr/bin/env bash

cd "$(dirname "$0")/.." || exit

if [ -f .env ]; then
    source .env
fi
bin/dgdgo server
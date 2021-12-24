#!/bin/sh
set -e
if [ "$1" = './grpc-ecosystem-template' ]; then
  echo "do preset env here"
fi

# exec command
exec "$@"
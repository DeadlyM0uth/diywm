#!/bin/bash


XEPHYR=$(whereis -b Xephyr | sed -E 's/^.*: ?//')


if [ -z "$XEPHYR" ]; then
  echo "Xephyr not found"
  exit 1
fi

xinit ./xinitrc -- "$XEPHYR" :0 -ac -screen 800x600 -host-cursor



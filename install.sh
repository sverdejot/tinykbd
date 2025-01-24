#!/usr/bin/env bash

BINARY_URL="https://assets.sverdejot.dev/tinykbd/tinykbd"
BINARY_NAME="tinykbd"
INSTALL_DIR="$HOME/.tinykbd/bin"
LOG_DIR="$HOME/.tinykbd/logs"
LOG_FILE="$LOG_DIR/install.log"

mkdir -p "$INSTALL_DIR"
mkdir -p "$LOG_DIR"

exec >>"$LOG_FILE" 2>&1

curl -sSfLO "$BINARY_URL"
if [ $? -ne 0 ]; then
  exit 1
fi

chmod +x "$BINARY_NAME"

mv "$BINARY_NAME" "$INSTALL_DIR"/

nohup "$INSTALL_DIR/$BINARY_NAME" &

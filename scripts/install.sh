#!/usr/bin/env bash
set -euo pipefail

REPO="maestroi/node-manager-cli"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
BINARY_NAME="${BINARY_NAME:-node-manager-cli}"
VERSION="${VERSION:-latest}"

if [ "$VERSION" = "latest" ]; then
	DOWNLOAD_URL="https://github.com/${REPO}/releases/latest/download/${BINARY_NAME}"
else
	DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION}/${BINARY_NAME}"
fi

TMP_FILE="$(mktemp)"
cleanup() {
	rm -f "$TMP_FILE"
}
trap cleanup EXIT

echo "Downloading ${BINARY_NAME} (${VERSION})..."
curl -fsSL "$DOWNLOAD_URL" -o "$TMP_FILE"
chmod +x "$TMP_FILE"

DEST="${INSTALL_DIR%/}/${BINARY_NAME}"
if [ -w "$INSTALL_DIR" ]; then
	install -m 755 "$TMP_FILE" "$DEST"
else
	echo "Installing to ${DEST} (requires sudo)..."
	sudo install -m 755 "$TMP_FILE" "$DEST"
fi

echo "Installed ${DEST}"
"$DEST" version

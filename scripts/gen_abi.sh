#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
ACTIONS_DIR="$ROOT_DIR/otim-protocol/src/actions"
OUT_DIR="$ROOT_DIR/otim-protocol/out"
ABI_DIR="$ROOT_DIR/abi"

mkdir -p "$ABI_DIR"

for action_path in "$ACTIONS_DIR"/*.sol; do
  action_file="$(basename "$action_path")"           # e.g. UniswapV3ExactInputAction.sol
  action_name="${action_file%.sol}"                   # e.g. UniswapV3ExactInputAction
  abi_json="$OUT_DIR/$action_file/$action_name.abi.json"

  if [[ ! -f "$abi_json" ]]; then
    echo "[WARN] ABI not found for $action_name at $abi_json, skipping" >&2
    continue
  fi

  pkg_name="$(echo "$action_name" | tr 'A-Z' 'a-z')"
  pkg_dir="$ABI_DIR/$pkg_name"
  mkdir -p "$pkg_dir"

  out_file="$pkg_dir/$action_name.go"
  echo "[INFO] Generating Go binding for $action_name into package $pkg_name"
  abigen --abi="$abi_json" --pkg="$pkg_name" --out="$out_file"

done

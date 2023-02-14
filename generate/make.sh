#!/bin/bash

set -ex

WORKDIR="$(mktemp -d)"
trap 'rm -rf -- "${WORKDIR}"' EXIT

OUTDIR="../fsdata"

rm -rf "${OUTDIR}"
mkdir -p "${OUTDIR}"

for SIZE in $(seq 10 10 250); do
  truncate "${WORKDIR}/${SIZE}g" --size "${SIZE}g"
  mkfs.ext4 -F "${WORKDIR}/${SIZE}g"
  go run . --input="${WORKDIR}/${SIZE}g" --size="${SIZE}" --output-dir="${OUTDIR}"
done

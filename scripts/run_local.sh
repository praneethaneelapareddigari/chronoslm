#!/usr/bin/env bash
set -euo pipefail
go run services/go/ingestor/main.go &
go run services/go/serving/main.go &
wait

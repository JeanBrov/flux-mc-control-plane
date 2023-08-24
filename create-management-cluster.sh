#!/usr/bin/env bash

set -euo pipefail

kind create cluster --name management.9
flux install
kubectl apply -f control-plane/gitrepo.yaml
kubectl apply -f control-plane/ks.yaml

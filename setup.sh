#!/bin/bash

set -euo pipefail

pre-commit install
pre-commit install --hook-type commit-msg

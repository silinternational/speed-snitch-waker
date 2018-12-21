#!/usr/bin/env bash

# Exit script with error if any step fails.
set -e

# Build binaries
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
$DIR/build.sh

# Export env vars
export LAMBDA_FUNCTION1="${DEV_LAMBDA_FUNCTION1}"
export LAMBDA_FUNCTION2="${DEV_LAMBDA_FUNCTION2}"
export LAMBDA_FUNCTION3="${DEV_LAMBDA_FUNCTION3}"
export LAMBDA_FUNCTION4="${DEV_LAMBDA_FUNCTION4}"
export LAMBDA_FUNCTION5="${DEV_LAMBDA_FUNCTION5}"
export LAMBDA_FUNCTION6="${DEV_LAMBDA_FUNCTION6}"
export LAMBDA_FUNCTION7="${DEV_LAMBDA_FUNCTION7}"
export LAMBDA_FUNCTION8="${DEV_LAMBDA_FUNCTION8}"
export LAMBDA_FUNCTION9="${DEV_LAMBDA_FUNCTION9}"

serverless deploy -v --stage dev
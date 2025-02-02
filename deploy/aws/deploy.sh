#!/bin/bash

# Lymphly Deployer
# This deployer will deploy Lymply to a given AWS account with the provided keys
# It is expected that this script is executing in a compiled deployer 
# With all required parts in their correct locations in the deployer

APPLICATION_NAME=$1
ENVIRONMENT_NAME=$2
DEPLOYMENT_REGION=$3
STATE_REGION=$4
STATE_BUCKET=$5
STATE_TABLE=$6
export AWS_ACCESS_KEY_ID=$7
export AWS_SECRET_ACCESS_KEY=$8
RELEASES_PATH=$9

set -e
set -x

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

pushd $SCRIPT_DIR/terraform

terraform init \
    -backend-config="bucket=$STATE_BUCKET" \
    -backend-config="dynamodb_table=$STATE_TABLE" \
    -backend-config="key=terraform/$APPLICATION_NAME/$ENVIRONMENT_NAME/$ENVIRONMENT_NAME" \
    -backend-config="region=$STATE_REGION"

export TF_VAR_application_name="$APPLICATION_NAME"
export TF_VAR_environment_name="$ENVIRONMENT_NAME"
export TF_VAR_deployment_region="$DEPLOYMENT_REGION"
export TF_VAR_releases_path="$RELEASES_PATH"

terraform apply -auto-approve

popd
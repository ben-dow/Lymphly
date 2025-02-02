#!/bin/bash

# Lymphly Deployer
# This deployer will deploy Lymply to a given AWS account with the provided keys
# It is expected that this script is executing in a compiled deployer 
# With all required parts in their correct locations in the deployer

# Expects the following environment variables exported in the environment
# APPLIACTION_NAME
# ENVIRONMENT_NAME
# DEPLOYMENT_REGION
# STATE_REGION
# STATE_BUCKET
# STATE_TABLE
# AWS_ACCESS_KEY_ID
# AWS_SECRET_ACCESS_KEY
# RELEASES_PATH

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

pushd $SCRIPT_DIR/terraform
terraform init \
    -backend-config="bucket=$STATE_BUCKET" \
    -backend-config="dynamodb_table=$STATE_TABLE" \
    -backend-config="key=terraform/$APPLICATION_NAME/$ENVIRONMENT_NAME/$ENVIRONMENT_NAME" \
    -backend-config="region=$STATE_REGION"

export TF_VAR_application_name="$APPLIACTION_NAME"
export TF_VAR_environment_name="$ENVIRONMENT_NAME"
export TF_VAR_deployment_region="$DEPLOYMENT_REGION"
export TF_VAR_releases_path="$RELEASES_PATH"

terraform apply -auto-approve
popd
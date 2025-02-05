#!/bin/bash

# Lymphly Deployer

# APPLICATION_NAME=$1
# ENVIRONMENT_NAME=$2
# DEPLOYMENT_REGION=$3
# STATE_REGION=$4
# STATE_BUCKET=$5
# STATE_TABLE=$6
# export AWS_ACCESS_KEY_ID=$7
# export AWS_SECRET_ACCESS_KEY=$8
# RELEASES_PATH=$9
# RADAR_PUBLIC_KEY

set -e
set -x

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

pushd $RELEASES_PATH
unzip frontend.zip
echo "$RADAR_PUBLIC_KEY" > dist/radar_pub_key.txt
popd

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
export TF_VAR_website_build_location="$RELEASES_PATH/dist/"

terraform apply -auto-approve

popd
#!/bin/bash
pushd ../
#make build
#make install
make build13
popd
rm -rf .terraform
rm -rf terraform-provider-zscaler
rm -rf .terraform.lock.hcl
rm -rf terraform.tfstate
rm -rf terraform.tfstate.backup
terraform init && terraform apply --auto-approve

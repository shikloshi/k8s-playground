#!/bin/bash

set -ex

prefix="gcr.io"

if [[ "$#" -ne 2 ]]; then
    echo "should specify service name and tag"
    exit 1
fi

service_name=$1
tag=$2

if [[ -z $PROJECT_ID ]]; then
    p=$(gcloud projects list | tail -1 | cut -d" " -f1)
    export PROJECT_ID=$p
    image_full_name=$prefix/$PROJECT_ID/${service_name}:${tag}
else
    image_full_name=$prefix/$PROJECT_ID/${service_name}:${tag}
fi

docker build -t $image_full_name $PWD/$service_name && docker push $image_full_name

#sed -i -e "s@IMAGE@${image_full_name}@" ./k8s/$service_name-k8s.yaml

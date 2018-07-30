#!/bin/bash

set -ex

prefix="gcr.io"

if [[ "$#" -ne 1 ]]; then
    echo "should specify service tag"
    exit 1
fi

#service_name=$1
tag=$1

if [[ -z $PROJECT_ID ]]; then
    p=$(gcloud projects list | tail -1 | cut -d" " -f1)
    export PROJECT_ID=$p
    repository=$prefix/$PROJECT_ID
else
    repository=$prefix/$PROJECT_ID
fi

docker build -t $repository/meeting:$tag $PWD/meeting && docker push $repository/meeting:$tag
docker build -t $repository/worker:$tag $PWD/worker && docker push $repository/worker:$tag

#sed -i -e "s@IMAGE@${image_full_name}@" ./k8s/$service_name-k8s.yaml

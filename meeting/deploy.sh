#!/bin/bash

set -ex

image=${2:-meeting}
tag=${1:-0.0.1}
image_name=$(img-name $image $tag)
docker build -t $image_name ${3:-$PWD} && docker push $image_name

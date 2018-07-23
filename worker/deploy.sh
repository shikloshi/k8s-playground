#!/bin/bash

set -ex

name=${2:-worker}
tag=${1:-0.0.1}
image_name=$(img-name $name $tag)
docker build -t $image_name ${3:-$PWD} && docker push $image_name



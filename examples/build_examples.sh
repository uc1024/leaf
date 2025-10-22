#!/bin/bash

project=$(cd $(dirname $0);pwd)
proto="$project"
proto_lib="$HOME/go/bin/include"
echo $project

# exit 0

# $1 指定服务proto文件夹
function initProto(){
    protos=$(find $1 -type f -name '*.proto')
    echo $protos
    for item in ${protos[@]}
    do
        protoc \
        -I $project \
        -I $(dirname $proto) \
        -I ${proto_lib} \
        --go_out=$project \
        --go_opt paths=source_relative \
        --leaf_out $project \
        --leaf_opt paths=source_relative \
        $item
    done
}

initProto $project
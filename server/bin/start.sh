#!/bin/bash

# define
project_path=$(realpath $(dirname $(readlink -f $0))/../..)

# create $project_path/server/$1
function create_directory() {
    cd $project_path/server
    if [ ! -d "$1" ]; then
        mkdir $1
    fi
}

function upenv() {
    # create directory
    create_directory log
    create_directory sbin

    # go config
    go env -w GOROOT=$GOROOT
    go env -w GOPATH=$GOPATH
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.cn,direct

    # go mod tidy
    cd $project_path/server
    go mod tidy
}

function start_server() {
    cd $project_path/server

    # compile
    go build -o sbin/db_proj

    # run!
    nohup sbin/db_proj $@ 1>log/log 2>log/error &
}

upenv
start_server $@
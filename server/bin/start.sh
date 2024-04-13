#!/bin/bash

# define
project_path=$(realpath $(dirname $(readlink -f $0))/../..)
echo_prefix="[start.sh]"
file_color="\033[0;35m"
text_color="\033[0;31m"
done_color="\033[0;33m"
reset_color="\033[0m"

function echo_begin() {
    echo -e -n "$file_color$echo_prefix $reset_color"
    echo -e -n "$text_color$@...$reset_color"
}

function echo_done() {
    echo -e "$done_color ===> done!$reset_color"
}

# create $project_path/server/$1
function create_directory() {
    cd $project_path/server
    if [ ! -d "$1" ]; then
        mkdir $1
    fi
}

function upenv() {
    # create directory
    echo_begin "create directory: log, sbin"
        create_directory log
        create_directory sbin
    echo_done

    # go config
    echo_begin "set go env"
        go env -w GOROOT=$GOROOT
        go env -w GOPATH=$GOPATH
        go env -w GO111MODULE=on
        go env -w GOPROXY=https://goproxy.cn,direct
    echo_done

    # go mod tidy
    echo_begin "go mod tidy"
        cd $project_path/server
        go mod tidy
    echo_done
}

function start_server() {
    cd $project_path/server

    # compile
    echo_begin "compile"
        go build -o sbin/db_proj
    echo_done

    # run!
    echo_begin "run backend"
        nohup sbin/db_proj $@ 1>log/log 2>log/error &
    echo_done
}

upenv
start_server $@

echo -e -n "$file_color$echo_prefix $reset_color"
echo -e "start already!"

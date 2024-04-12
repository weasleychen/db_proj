#!/bin/bash

project_path=$(realpath $(dirname $(readlink -f $0))/../..)

cd $project_path/server/bin
./stop.sh
./start.sh $@
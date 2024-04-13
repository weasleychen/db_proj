#!/bin/bash

pid=$(ps -aux | grep -E "\ssbin/db_proj$" | awk '{ print $2 }')

if [ "$pid" != "" ]
then
    kill -9 $pid
fi

echo -e -n "\033[0;35m[stop.sh] \033[0m"
echo "stop already!"
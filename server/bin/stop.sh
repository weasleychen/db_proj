#!/bin/bash


if [ "$pid" != "" ]
then
    killall -9 db_proj msdbcall mstablemgr
fi

echo -e -n "\033[0;35m[stop.sh] \033[0m"
echo "stop already!"
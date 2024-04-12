#!/bin/bash

pid=$(ps -aux | grep -E "\ssbin/db_proj$" | awk '{ print $2 }')

if [ "$pid" != "" ]
then
    kill -9 $pid
fi
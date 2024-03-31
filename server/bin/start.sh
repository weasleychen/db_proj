#!/bin/bash

cd ~/db_proj/server/ && swag init && \
go build -o bin/main main.go && \
bin/main $@
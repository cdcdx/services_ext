#!/bin/bash

source .env

# GOPROXY
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
# GOLOG
export GOLOG_LOG_FMT=json    #json/nocolor
export GOLOG_LOG_LEVEL=info  #trace/debug/info/warn/error/fatal

if [ -f ./services_ext ]; then
  rm ./services_ext
fi

echo "go build"
go build

if [ -f ./services_ext ]; then
  echo "./services_ext"
  ./services_ext
fi
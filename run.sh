#!/usr/bin/env bash

export GOPATH=/c/snake/code/GoPro
DIR=`pwd`
now=`date "+%Y/%m/%d %H:%M:%s"`
echo "${now} Start Build ..."

cd $DIR && go build web

now=`date "+%Y/%m/%d %H:%M:%s"`
echo "${now} Start Server ..."

./web.exe

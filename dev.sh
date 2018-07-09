#!/usr/bin/env bash

export GOPATH=/c/snake/code/GoPro
DIR=`pwd`
now=`date "+%Y/%m/%d %H:%M:%s"`
echo "${now} Start develop server ..."

cd $DIR && /c/snake/code/GoPro/bin/bee.exe run


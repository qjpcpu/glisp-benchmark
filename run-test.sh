#!/bin/bash
if [ ! -z "$1" ];then
	go test -bench="$1"
else
        go test -bench=.
fi

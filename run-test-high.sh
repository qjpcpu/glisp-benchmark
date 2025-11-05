#!/bin/bash
if [ ! -z "$1" ];then
	go test -bench="$1"|grep Benchmark|awk '{print $1"\t"int(1000000/$3)" op/ms"}'
else
	go test -bench=. |grep Benchmark|awk '{print $1"\t"int(1000000/$3)" op/ms"}'
fi


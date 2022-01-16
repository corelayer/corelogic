#!/bin/sh
config=$1
echo "Total lines:\t\t$(wc -l $config)"
echo "Objects not replaced:\t$(cat $config | grep "<<" | wc -l)"
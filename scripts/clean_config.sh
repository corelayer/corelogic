#!/bin/sh
clean_file() {
    if [ -f $1 ]
    then
        echo "Removing $1"
        rm $1
    fi
}

clean_file config.txt
clean_file output.txt
sleep 2
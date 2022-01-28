#!/bin/sh
clean_file() {
    if [ -f $1 ]
    then
        echo "Removing $1"
        rm $1
    fi
}

clean_file config.txt
clean_file config.json
clean_file config.out
clean_file config.conf

clean_file output.txt
clean_file output.conf.out
sleep 2
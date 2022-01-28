#!/bin/sh
######################################################
### WARNING                                        ###
######################################################
### THIS SCRIPT IS INTENDED TO BE RUN THROUGH MAKE ###
### DO NOT CALL THIS SCRIPT DIRECTLY               ###
######################################################

create_protocol_ipfilter(){
    mkdir -p assets/framework/$version/packages/$package/$protocol
    cp -r assets/framework/$version/packages/$package/fake/*.yaml assets/framework/$version/packages/$package/$protocol/.
    sed -i "s/fake/$protocol/g" assets/framework/$version/packages/$package/$protocol/*
    sed -i "s/FAKE/$upperProtocol/g" assets/framework/$version/packages/$package/$protocol/*
    sed -i "s/BASEPROTOCOL/$upperBaseProtocol/g" assets/framework/$version/packages/$package/$protocol/*
}

version=$1
package=$2
protocol=$3
baseProtocol=$4

upperProtocol=$(echo $protocol | tr '[:lower:]' '[:upper:]')
upperBaseProtocol=$(echo $baseProtocol | tr '[:lower:]' '[:upper:]')

create_protocol_ipfilter
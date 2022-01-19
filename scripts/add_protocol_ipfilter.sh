#!/bin/sh
######################################################
### WARNING                                        ###
######################################################
### THIS SCRIPT IS INTENDED TO BE RUN THROUGH MAKE ###
### DO NOT CALL THIS SCRIPT DIRECTLY               ###
######################################################

create_protocol_ipfilter(){
    local package=$1
    mkdir -p assets/framework/$version/packages/$package/$protocol
    cp -r assets/framework/$version/packages/$package/fake/* assets/framework/$version/packages/$package/$protocol/.
    sed -i "s/fake/$protocol/g" assets/framework/$version/packages/$package/$protocol/*
    sed -i "s/FAKE/$upperProtocol/g" assets/framework/$version/packages/$package/$protocol/*
    sed -i "s/BASEPROTOCOL/$upperBaseProtocol/g" assets/framework/$version/packages/$package/$protocol/*
}

version=$1
protocol=$2
baseProtocol=$3

upperProtocol=$(echo $protocol | tr '[:lower:]' '[:upper:]')
upperBaseProtocol=$(echo $baseProtocol | tr '[:lower:]' '[:upper:]')


if [ $protocol != "fake" ]
then
create_protocol_ipfilter core
create_protocol_ipfilter contentswitching
create_protocol_ipfilter loadbalancers
fi
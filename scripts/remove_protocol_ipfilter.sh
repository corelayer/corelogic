#!/bin/sh
######################################################
### WARNING                                        ###
######################################################
### THIS SCRIPT IS INTENDED TO BE RUN THROUGH MAKE ###
### DO NOT CALL THIS SCRIPT DIRECTLY               ###
######################################################
version=$1
protocol=$2

if [ $protocol != "fake" ]
then
rm -rf assets/framework/$version/packages/core/$protocol
rm -rf assets/framework/$version/packages/contentswitching/$protocol
rm -rf assets/framework/$version/packages/loadbalancers/$protocol

rm -rf assets/framework/$version/packages/demo/$protocol
fi
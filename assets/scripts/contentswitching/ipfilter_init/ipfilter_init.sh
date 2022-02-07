#!/bin/bash
######################################################
### WARNING                                        ###
######################################################
### THIS SCRIPT IS INTENDED TO BE RUN THROUGH MAKE ###
### DO NOT CALL THIS SCRIPT DIRECTLY               ###
######################################################




######################################################
### GLOBAL VARIABLES && IMPORTS                    ###
######################################################
###

source assets/scripts/global.sh
source assets/scripts/headers.sh

###

inputFilename_ipfilter_init_basePath='assets/templates/framework/$version/packages/contentswitching/ipfilter_init'
inputFilename_ipfilter_packageHeader="$inputFilename_ipfilter_init_basePath"'/package_header.yaml'

inputFilename_ipfilter_init_trafficmanagement_contentswitching_policies="$inputFilename_ipfilter_init_basePath"'/init_ipfilter_trafficmanagement_contentswitching_policies.yaml'
inputFilename_ipfilter_init_trafficmanagement_contentswitching_policylabels="$inputFilename_ipfilter_init_basePath"'/init_ipfilter_trafficmanagement_contentswitching_policylabels.yaml'
inputFilename_ipfilter_init_trafficmanagement_contentswitching_policylabelbindings="$inputFilename_ipfilter_init_basePath"'/init_ipfilter_trafficmanagement_contentswitching_policylabelbindings_sequence.yaml'

###

outputFilename_ipfilter_init_basePath='assets/framework/$version/packages/contentswitching/$protocol'
outputFilename_ipfilter_init="$outputFilename_ipfilter_init_basePath"'/init_ipfilter_$ipversion.yaml'

###
######################################################




######################################################
### FUNCTIONS                                      ###
######################################################
###

create_ipfilter_init() {
  version=$1
  protocol=$2
  ipversion=$3
  bindingprefix=$4

  outputFilename=$outputFilename_ipfilter_init
  outputFilename=$(sed "s/\$version/$version/g" <<< $outputFilename)
  outputFilename=$(sed "s/\$protocol/$protocol/g" <<< $outputFilename)
  outputFilename=$(sed "s/\$ipversion/$ipversion/g" <<< $outputFilename)

  create_package_header $version $protocol $ipversion $outputFilename

  add_section_header_trafficmanagement_contentswitching_policies $outputFilename
  create_object $version $protocol $ipversion "" $inputFilename_ipfilter_init_trafficmanagement_contentswitching_policies $outputFilename

  add_section_header_trafficmanagement_contentswitching_policylabels $outputFilename
  create_object $version $protocol $ipversion $filtertype $inputFilename_ipfilter_init_trafficmanagement_contentswitching_policylabels $outputFilename

  add_section_header_trafficmanagement_contentswitching_policylabelbindings $outputFilename
  create_object_type_policybindings_sequence $version $inputFilename_ipfilter_init_trafficmanagement_contentswitching_policylabelbindings $outputFilename

  replace_ipversion $ipversion $outputFilename
  replace_protocol $protocol $outputFilename
  replace_filtertype $filtertype $outputFilename
  replace_bindingprefix $ipversion $bindingprefix $outputFilename
}

###

create_package_header() {
  version=$1
  protocol=$2
  ipversion=$3
  output=$4

  inputFilename_ipfilter_packageHeader=$(sed "s/\$version/$version/g" <<< $inputFilename_ipfilter_packageHeader)
  inputFilename_ipfilter_packageHeader=$(sed "s/\$filtertype/$filtertype/g" <<< $inputFilename_ipfilter_packageHeader)

  cat $inputFilename_ipfilter_packageHeader >> $output
}

###

create_object_type_policybindings_sequence() {
  version=$1
  input=$2
  output=$3

  for i in "${ipfilter_frontend_sequence[@]}"
  do
    echo $i
    input=$(sed "s/\$version/$version/g" <<< $input)
    input=$(sed "s/\$filtertype/$i/g" <<< $input)
    cat $input
    cat $input >> $output
  done
}

###
######################################################




######################################################
### MAIN                                           ###
######################################################

version=$1
protocol=$2

mkdir -p "assets/framework/$version/packages/contentswitching/$protocol"

create_ipfilter_init $version $protocol ipv4 101
create_ipfilter_init $version $protocol ipv6 101
######################################################
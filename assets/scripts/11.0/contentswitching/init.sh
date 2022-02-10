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

inputFilename_init_basePath='assets/templates/framework/$version/packages/contentswitching/init'
inputFilename_packageHeader="$inputFilename_init_basePath"'/package_header.yaml'

inputFilename_init_trafficmanagement_contentswitching_policies="$inputFilename_init_basePath"'/init_trafficmanagement_contentswitching_policies.yaml'
inputFilename_init_trafficmanagement_contentswitching_policylabels="$inputFilename_init_basePath"'/init_trafficmanagement_contentswitching_policylabels.yaml'
inputFilename_init_trafficmanagement_contentswitching_policylabelbindings="$inputFilename_init_basePath"'/init_trafficmanagement_contentswitching_policylabelbindings.yaml'

###

outputFilename_init_basePath='assets/framework/$version/packages/contentswitching/$protocol'
outputFilename_init="$outputFilename_init_basePath"'/init.yaml'

###
######################################################




######################################################
### FUNCTIONS                                      ###
######################################################
###

create_init() {
  version=$1
  protocol=$2

  outputFilename=$outputFilename_init
  outputFilename=$(sed "s/\$version/$version/g" <<< $outputFilename)
  outputFilename=$(sed "s/\$protocol/$protocol/g" <<< $outputFilename)

  create_package_header $version $protocol $outputFilename

  add_section_header_trafficmanagement_contentswitching_policies $outputFilename
  create_object $version $protocol "" "" $inputFilename_init_trafficmanagement_contentswitching_policies $outputFilename

  add_section_header_trafficmanagement_contentswitching_policylabels $outputFilename
  create_object $version $protocol "" "" $inputFilename_init_trafficmanagement_contentswitching_policylabels $outputFilename

  add_section_header_trafficmanagement_contentswitching_policylabelbindings $outputFilename
  create_object $version $protocol "" "" $inputFilename_init_trafficmanagement_contentswitching_policylabelbindings $outputFilename

  replace_protocol $protocol $outputFilename
}

###

create_package_header() {
  version=$1
  protocol=$2
  output=$3

  inputFilename_packageHeader=$(sed "s/\$version/$version/g" <<< $inputFilename_packageHeader)

  cat $inputFilename_packageHeader >> $output
}

###

# create_object_type_policybindings_sequence() {
#   version=$1
#   input=$2
#   output=$3

#   for i in "${ipfilter_frontend_sequence[@]}"
#   do
#     echo $i
#     input=$(sed "s/\$version/$version/g" <<< $input)
#     input=$(sed "s/\$filtertype/$i/g" <<< $input)
#     cat $input
#     cat $input >> $output
#   done
# }

###
######################################################




######################################################
### MAIN                                           ###
######################################################

version=$1
protocol=$2

mkdir -p "assets/framework/$version/packages/contentswitching/$protocol"

create_init $version $protocol

######################################################
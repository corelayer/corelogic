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

source scripts/templates/global.sh
source scripts/templates/headers.sh
source scripts/templates/contentswitching/ipfilter.sh

###

inputFilename_ipfilter_blocklist_basePath='templates/framework/$version/packages/contentswitching/fake'
inputFilename_ipfilter_blocklist_system_auditing_messageactions_sequence="$inputFilename_ipfilter_blocklist_basePath"'/ipfilter_blocklist_system_auditing_messageactions_sequence.yaml'
inputFilename_ipfilter_blocklist_trafficmanagement_contentswitching_policies_sequence="$inputFilename_ipfilter_blocklist_basePath"'/ipfilter_blocklist_trafficmanagement_contentswitching_policies_sequence.yaml'
inputFilename_ipfilter_blocklist_trafficmanagement_contentswitching_policylabels="$inputFilename_ipfilter_blocklist_basePath"'/ipfilter_blocklist_trafficmanagement_contentswitching_policylabels.yaml'
inputFilename_ipfilter_blocklist_trafficmanagement_contentswitching_policylabelbindings_sequence="$inputFilename_ipfilter_blocklist_basePath"'/ipfilter_blocklist_trafficmanagement_contentswitching_policylabelbindings_sequence.yaml'
inputFilename_ipfilter_blocklist_trafficmanagement_contentswitching_policylabelbindings="$inputFilename_ipfilter_blocklist_basePath"'/ipfilter_blocklist_trafficmanagement_contentswitching_policylabelbindings.yaml'

###

outputFilename_ipfilter_blocklist_basePath='assets/framework/$version/packages/contentswitching/$protocol'
outputFilename_ipfilter_blocklist="$outputFilename_ipfilter_blocklist_basePath"'/$filtertype_$ipversion_ipfilter_blocklist.yaml'

###
######################################################




######################################################
### FUNCTIONS                                      ###
######################################################
###

create_ipfilter_blocklist() {
  version=$1
  protocol=$2
  ipversion=$3
  filtertype=$4
  bindingprefix=$5

  outputFilename=$outputFilename_ipfilter_blocklist
  outputFilename=$(sed "s/\$version/$version/g" <<< $outputFilename)
  outputFilename=$(sed "s/\$protocol/$protocol/g" <<< $outputFilename)
  outputFilename=$(sed "s/\$ipversion/$ipversion/g" <<< $outputFilename)
  outputFilename=$(sed "s/\$filtertype/$filtertype/g" <<< $outputFilename)

  create_package_header $version $protocol $ipversion $filtertype $outputFilename

  add_section_header_system_auditing_messageactions $outputFilename
  create_object_sequence $version $protocol $ipversion $filtertype $inputFilename_ipfilter_blocklist_system_auditing_messageactions_sequence $outputFilename

  add_section_header_trafficmanagement_contentswitching_policies $outputFilename
  create_object_sequence $version $protocol $ipversion $filtertype $inputFilename_ipfilter_blocklist_trafficmanagement_contentswitching_policies_sequence $outputFilename

  add_section_header_trafficmanagement_contentswitching_policylabels $outputFilename
  create_object $version $protocol $ipversion $filtertype $inputFilename_ipfilter_blocklist_trafficmanagement_contentswitching_policylabels $outputFilename

  add_section_header_trafficmanagement_contentswitching_policylabelbindings $outputFilename
  create_objectbindings_sequence $version $protocol $ipversion $filtertype $inputFilename_ipfilter_blocklist_trafficmanagement_contentswitching_policylabelbindings_sequence $outputFilename
  create_object $version $protocol $ipversion $filtertype $inputFilename_ipfilter_blocklist_trafficmanagement_contentswitching_policylabelbindings $outputFilename

  replace_ipfilter_next_label $filtertype $outputFilename

  replace_ipversion $ipversion $outputFilename
  replace_protocol $protocol $outputFilename
  replace_filtertype $filtertype $outputFilename
  replace_bindingprefix $ipversion $bindingprefix $outputFilename
}

###
######################################################


######################################################
### MAIN                                           ###
######################################################

version=$1
protocol=$2

create_ipfilter_blocklist $version $protocol ipv4 csv 109
create_ipfilter_blocklist $version $protocol ipv6 csv 109

create_ipfilter_blocklist $version $protocol ipv4 csvgroup 107
create_ipfilter_blocklist $version $protocol ipv6 csvgroup 107

create_ipfilter_blocklist $version $protocol ipv4 tenant 105
create_ipfilter_blocklist $version $protocol ipv6 tenant 105

create_ipfilter_blocklist $version $protocol ipv4 endpoint 103
create_ipfilter_blocklist $version $protocol ipv6 endpoint 103

######################################################
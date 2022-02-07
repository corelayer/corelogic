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
source assets/scripts/contentswitching/ipfilter_frontend/ipfilter_frontend.sh

###

inputFilename_ipfilter_allowlist_basePath='assets/templates/framework/$version/packages/contentswitching/ipfilter_frontend'
inputFilename_ipfilter_allowlist_system_auditing_messageactions="$inputFilename_ipfilter_allowlist_basePath"'/ipfilter_allowlist_system_auditing_messageactions.yaml'
inputFilename_ipfilter_allowlist_trafficmanagement_contentswitching_policies_sequence="$inputFilename_ipfilter_allowlist_basePath"'/ipfilter_allowlist_trafficmanagement_contentswitching_policies_sequence.yaml'
inputFilename_ipfilter_allowlist_trafficmanagement_contentswitching_policies="$inputFilename_ipfilter_allowlist_basePath"'/ipfilter_allowlist_trafficmanagement_contentswitching_policies.yaml'
inputFilename_ipfilter_allowlist_trafficmanagement_contentswitching_policylabels="$inputFilename_ipfilter_allowlist_basePath"'/ipfilter_allowlist_trafficmanagement_contentswitching_policylabels.yaml'
inputFilename_ipfilter_allowlist_trafficmanagement_contentswitching_policylabelbindings_sequence="$inputFilename_ipfilter_allowlist_basePath"'/ipfilter_allowlist_trafficmanagement_contentswitching_policylabelbindings_sequence.yaml'
inputFilename_ipfilter_allowlist_trafficmanagement_contentswitching_policylabelbindings="$inputFilename_ipfilter_allowlist_basePath"'/ipfilter_allowlist_trafficmanagement_contentswitching_policylabelbindings.yaml'

###

outputFilename_ipfilter_allowlist_basePath='assets/framework/$version/packages/contentswitching/$protocol'
outputFilename_ipfilter_allowlist="$outputFilename_ipfilter_allowlist_basePath"'/$filtertype_$ipversion_ipfilter_allowlist.yaml'

###
######################################################




######################################################
### FUNCTIONS                                      ###
######################################################
###

create_ipfilter_allowlist() {
  version=$1
  protocol=$2
  ipversion=$3
  filtertype=$4
  bindingprefix=$5

  outputFilename=$outputFilename_ipfilter_allowlist
  outputFilename=$(sed "s/\$version/$version/g" <<< $outputFilename)
  outputFilename=$(sed "s/\$protocol/$protocol/g" <<< $outputFilename)
  outputFilename=$(sed "s/\$ipversion/$ipversion/g" <<< $outputFilename)
  outputFilename=$(sed "s/\$filtertype/$filtertype/g" <<< $outputFilename)

  create_package_header $version $protocol $ipversion $filtertype $outputFilename

  add_section_header_system_auditing_messageactions $outputFilename
  create_object $version $protocol $ipversion $filtertype $inputFilename_ipfilter_allowlist_system_auditing_messageactions $outputFilename

  add_section_header_trafficmanagement_contentswitching_policies $outputFilename
  create_object_ipversion_sequence $version $protocol $ipversion $filtertype $inputFilename_ipfilter_allowlist_trafficmanagement_contentswitching_policies_sequence $outputFilename
  create_object $version $protocol $ipversion $filtertype $inputFilename_ipfilter_allowlist_trafficmanagement_contentswitching_policies $outputFilename

  add_section_header_trafficmanagement_contentswitching_policylabels $outputFilename
  create_object $version $protocol $ipversion $filtertype $inputFilename_ipfilter_allowlist_trafficmanagement_contentswitching_policylabels $outputFilename

  add_section_header_trafficmanagement_contentswitching_policylabelbindings $outputFilename
  create_objectbindings_ipversion_sequence $version $protocol $ipversion $filtertype $inputFilename_ipfilter_allowlist_trafficmanagement_contentswitching_policylabelbindings_sequence $outputFilename
  create_object $version $protocol $ipversion $filtertype $inputFilename_ipfilter_allowlist_trafficmanagement_contentswitching_policylabelbindings $outputFilename

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

create_ipfilter_allowlist $version $protocol ipv4 endpoint 103
create_ipfilter_allowlist $version $protocol ipv6 endpoint 103

create_ipfilter_allowlist $version $protocol ipv4 tenant 105
create_ipfilter_allowlist $version $protocol ipv6 tenant 105

create_ipfilter_allowlist $version $protocol ipv4 csvgroup 107
create_ipfilter_allowlist $version $protocol ipv6 csvgroup 107

create_ipfilter_allowlist $version $protocol ipv4 csv 109
create_ipfilter_allowlist $version $protocol ipv6 csv 109

######################################################
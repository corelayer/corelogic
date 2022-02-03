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

###

inputFilename_ipfilter_init_basePath='templates/framework/$version/packages/contentswitching/fake'
inputFilename_ipfilter_packageHeader="$inputFilename_ipfilter_init_basePath"'/ipfilter_package_header.yaml'

inputFilename_ipfilter_init_system_auditing_messageactions="$inputFilename_ipfilter_init_basePath"'/ipfilter_system_auditing_messageactions.yaml'
inputFilename_ipfilter_init_trafficmanagement_contentswitching_actions="$inputFilename_ipfilter_init_basePath"'/ipfilter_trafficmanagement_contentswitching_actions.yaml'
inputFilename_ipfilter_init_trafficmanagement_contentswitching_policies="$inputFilename_ipfilter_init_basePath"'/ipfilter_trafficmanagement_contentswitching_policies.yaml'
inputFilename_ipfilter_init_trafficmanagement_contentswitching_policylabels="$inputFilename_ipfilter_init_basePath"'/ipfilter_trafficmanagement_contentswitching_policylabels.yaml'
inputFilename_ipfilter_init_trafficmanagement_contentswitching_policylabelbindings="$inputFilename_ipfilter_init_basePath"'/ipfilter_trafficmanagement_contentswitching_policylabelbindings.yaml'

endpoint_ipfilter_next_label='<<contentswitching.tenant_$ipversion_ipfilter.trafficmanagement.contentswitching.policylabels.TENANT_$IPVERSION_IPFILTER_$PROTOCOL/name>>'
tenant_ipfilter_next_label='<<contentswitching.csvgroup_$ipversion_ipfilter.trafficmanagement.contentswitching.policylabels.CSVGROUP_$IPVERSION_IPFILTER_$PROTOCOL/name>>'
csvgroup_ipfilter_next_label='<<contentswitching.csv_$ipversion_ipfilter.trafficmanagement.contentswitching.policylabels.CSV_$IPVERSION_IPFILTER_$PROTOCOL/name>>'
csv_ipfilter_next_label='CSL_CL1100_ZONE_$IPVERSION_FAKE'

###

outputFilename_ipfilter_init_basePath='assets/framework/$version/packages/contentswitching/$protocol'
outputFilename_ipfilter_init="$outputFilename_ipfilter_init_basePath"'/$filtertype_$ipversion_ipfilter.yaml'

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
  filtertype=$4
  bindingprefix=$5

  outputFilename=$outputFilename_ipfilter_init
  outputFilename=$(sed "s/\$version/$version/g" <<< $outputFilename)
  outputFilename=$(sed "s/\$protocol/$protocol/g" <<< $outputFilename)
  outputFilename=$(sed "s/\$ipversion/$ipversion/g" <<< $outputFilename)
  outputFilename=$(sed "s/\$filtertype/$filtertype/g" <<< $outputFilename)

  create_package_header $version $protocol $ipversion $filtertype $outputFilename

  add_section_header_system_auditing_messageactions $outputFilename
  create_object $version $protocol $ipversion $filtertype $inputFilename_ipfilter_init_system_auditing_messageactions $outputFilename

  add_section_header_trafficmanagement_contentswitching_actions $outputFilename
  create_object $version $protocol $ipversion $filtertype $inputFilename_ipfilter_init_trafficmanagement_contentswitching_actions $outputFilename

  add_section_header_trafficmanagement_contentswitching_policies $outputFilename
  create_object $version $protocol $ipversion $filtertype $inputFilename_ipfilter_init_trafficmanagement_contentswitching_policies $outputFilename

  add_section_header_trafficmanagement_contentswitching_policylabels $outputFilename
  create_object $version $protocol $ipversion $filtertype $inputFilename_ipfilter_init_trafficmanagement_contentswitching_policylabels $outputFilename

  add_section_header_trafficmanagement_contentswitching_policylabelbindings $outputFilename
  create_object $version $protocol $ipversion $filtertype $inputFilename_ipfilter_init_trafficmanagement_contentswitching_policylabelbindings $outputFilename

  replace_ipfilter_next_label $filtertype $outputFilename

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
  filtertype=$4
  output=$5

  inputFilename_ipfilter_packageHeader=$(sed "s/\$version/$version/g" <<< $inputFilename_ipfilter_packageHeader)
  inputFilename_ipfilter_packageHeader=$(sed "s/\$filtertype/$filtertype/g" <<< $inputFilename_ipfilter_packageHeader)

  cat $inputFilename_ipfilter_packageHeader >> $output
}

###

replace_ipfilter_next_label() {
  filtertype=$1
  output=$2

  case $filtertype in
    endpoint)
      ipfilter_next_label="$endpoint_ipfilter_next_label"
      ;;
    
    tenant)
      ipfilter_next_label="$tenant_ipfilter_next_label"
      ;;

    csvgroup)
      ipfilter_next_label="$csvgroup_ipfilter_next_label"
      ;;

    csv)
      ipfilter_next_label="$csv_ipfilter_next_label"
      ;;
  esac

  sed -i "s!\$ipfilter_next_label!$ipfilter_next_label!g" $output
}

###
######################################################




######################################################
### MAIN                                           ###
######################################################

version=$1
protocol=$2

mkdir -p "assets/framework/$version/packages/contentswitching/$protocol"

create_ipfilter_init $version $protocol ipv4 endpoint 102
create_ipfilter_init $version $protocol ipv6 endpoint 102

create_ipfilter_init $version $protocol ipv4 tenant 104
create_ipfilter_init $version $protocol ipv6 tenant 104

create_ipfilter_init $version $protocol ipv4 csvgroup 106
create_ipfilter_init $version $protocol ipv6 csvgroup 106

create_ipfilter_init $version $protocol ipv4 csv 108
create_ipfilter_init $version $protocol ipv6 csv 108

######################################################
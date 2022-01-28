#!/bin/bash
######################################################
### WARNING                                        ###
######################################################
### THIS SCRIPT IS INTENDED TO BE RUN THROUGH MAKE ###
### DO NOT CALL THIS SCRIPT DIRECTLY               ###
######################################################

create_ipfilter_blocklist() {
  version=$1
  protocol=$2
  ipversion=$3

  outputFilename='assets/framework/$version/packages/contentswitching/$protocol/csv_$ipversion_ipfilter_blocklist.yaml'
  outputFilename=$(sed "s/\$version/$version/g" <<< $outputFilename)
  outputFilename=$(sed "s/\$protocol/$protocol/g" <<< $outputFilename)
  outputFilename=$(sed "s/\$ipversion/$ipversion/g" <<< $outputFilename)

  # touch $outputFilename


  create_package_header $version $protocol $ipversion $outputFilename

  create_section_header_system_auditing_messageactions $version $protocol $ipversion $outputFilename
  create_elements_system_auditing_messageactions_sequence $version $protocol $ipversion $outputFilename

  create_section_header_trafficmanagement_contentswitching_policies $version $protocol $ipversion $outputFilename
  create_elements_trafficmanagement_contentswitching_policies_sequence $version $protocol $ipversion $outputFilename

  create_section_header_trafficmanagement_contentswitching_policylabels $version $protocol $ipversion $outputFilename
  create_elements_trafficmanagement_contentswitching_policylabels $version $protocol $ipversion $outputFilename

  create_section_header_trafficmanagement_contentswitching_policylabelbindings $version $protocol $ipversion $outputFilename
  create_elements_trafficmanagement_contentswitching_policylabelbindings_sequence $version $protocol $ipversion $outputFilename

  replace_ipversion $version $protocol $ipversion $outputFilename
  replace_protocol $version $protocol $ipversion $outputFilename
}


create_package_header() {
  version=$1
  protocol=$2
  ipversion=$3
  output=$4

  inputFilename='assets/framework/$version/packages/contentswitching/fake/templates/csv_ipfilter_blocklist_package_header.yaml'
  inputFilename=$(sed "s/\$version/$version/g" <<< $inputFilename)

  cat $inputFilename >> $output
}

create_section_header_system_auditing_messageactions() {
  version=$1
  protocol=$2
  ipversion=$3
  output=$4

  inputFilename='assets/framework/$version/packages/contentswitching/fake/templates/csv_ipfilter_blocklist__system__auditing__messageactions_section_header.yaml'
  inputFilename=$(sed "s/\$version/$version/g" <<< $inputFilename)
  cat $inputFilename >> $output
}

create_elements_system_auditing_messageactions_sequence() {
  version=$1
  protocol=$2
  ipversion=$3
  output=$4

  maxCount=0
  if [[ $ipversion == "ipv4" ]]; then
    maxCount=32
  else
    maxCount=128
  fi

  for ((i="$maxCount";i>0;i--))
  do
      sequence=$(printf "%03d" $i)
      inputFilename='assets/framework/$version/packages/contentswitching/fake/templates/csv_ipfilter_blocklist__system__auditing__messageactions_sequence.yaml'
      inputFilename=$(sed "s/\$version/$version/g" <<< $inputFilename)

      cat $inputFilename | sed "s/\$sequence/$sequence/g" >> $output
  done
}


create_section_header_trafficmanagement_contentswitching_policies() {
  version=$1
  protocol=$2
  ipversion=$3
  output=$4

  inputFilename='assets/framework/$version/packages/contentswitching/fake/templates/csv_ipfilter_blocklist__trafficmanagement__contentswitching__policies_section_header.yaml'
  inputFilename=$(sed "s/\$version/$version/g" <<< $inputFilename)
  cat $inputFilename >> $output
}


create_elements_trafficmanagement_contentswitching_policies_sequence() {
  version=$1
  protocol=$2
  ipversion=$3
  output=$4

  maxCount=0
  if [[ $ipversion == "ipv4" ]]; then
    maxCount=32
  else
    maxCount=128
  fi

  for ((i="$maxCount";i>0;i--))
  do
      sequence=$(printf "%03d" $i)
      inputFilename='assets/framework/$version/packages/contentswitching/fake/templates/csv_ipfilter_blocklist__trafficmanagement__contentswitching__policies_sequence.yaml'
      inputFilename=$(sed "s/\$version/$version/g" <<< $inputFilename)

      cat $inputFilename | sed "s/\$sequence/$sequence/g" >> $output
  done
}


create_section_header_trafficmanagement_contentswitching_policylabels() {
  version=$1
  protocol=$2
  ipversion=$3
  output=$4

  inputFilename='assets/framework/$version/packages/contentswitching/fake/templates/csv_ipfilter_blocklist__trafficmanagement__contentswitching__policylabels_section_header.yaml'
  inputFilename=$(sed "s/\$version/$version/g" <<< $inputFilename)
  cat $inputFilename >> $output
}

create_elements_trafficmanagement_contentswitching_policylabels() {
  version=$1
  protocol=$2
  ipversion=$3
  output=$4

  inputFilename='assets/framework/$version/packages/contentswitching/fake/templates/csv_ipfilter_blocklist__trafficmanagement__contentswitching__policylabels.yaml'
  inputFilename=$(sed "s/\$version/$version/g" <<< $inputFilename)
  cat $inputFilename >> $output
}


create_section_header_trafficmanagement_contentswitching_policylabelbindings() {
  version=$1
  protocol=$2
  ipversion=$3
  output=$4

  inputFilename='assets/framework/$version/packages/contentswitching/fake/templates/csv_ipfilter_blocklist__trafficmanagement__contentswitching__policylabelbindings_section_header.yaml'
  inputFilename=$(sed "s/\$version/$version/g" <<< $inputFilename)
  cat $inputFilename >> $output
}

create_elements_trafficmanagement_contentswitching_policylabelbindings_sequence() {
  version=$1
  protocol=$2
  ipversion=$3
  output=$4

  maxCount=0
  if [[ $ipversion == "ipv4" ]]; then
    maxCount=32
  else
    maxCount=128
  fi


  policyid_sequence=1
  for ((i="$maxCount";i>0;i--))
  do
      sequence=$(printf "%03d" $i)
      policyid=$(printf "%03d" $policyid_sequence)
      inputfile="assets/framework/$version/packages/contentswitching/fake/templates/csv_ipfilter_blocklist__trafficmanagement__contentswitching__policylabelbindings_sequence.yaml"
      tempfile="assets/framework/$version/packages/contentswitching/fake/templates/csv_ipfilter_blocklist_policylabelbindings_sequence_$policyid.yaml"

      cat $inputfile | sed "s/\$sequence/$sequence/g" >> $tempfile
      cat $tempfile | sed "s/\$policyid/$policyid/g" >> $output
        
      rm $tempfile
      ((policyid_sequence++))
  done
}

replace_ipversion() {
  version=$1
  protocol=$2
  ipversion=$3
  output=$4

  upperIpversion=$(echo $ipversion | tr '[:lower:]' '[:upper:]')

  sed -i "s/\$ipversion/$ipversion/g" $output
  sed -i "s/\$IPVERSION/$upperIpversion/g" $output
}

replace_protocol() {
  version=$1
  protocol=$2
  ipversion=$3
  output=$4

  upperProtocol=$(echo $protocol | tr '[:lower:]' '[:upper:]')

  sed -i "s/\$protocol/$protocol/g" $output
  sed -i "s/\$PROTOCOL/$upperProtocol/g" $output
}


version=$1
protocol=$2

mkdir -p "assets/framework/$version/packages/contentswitching/$protocol"

create_ipfilter_blocklist $version $protocol ipv4
create_ipfilter_blocklist $version $protocol ipv6
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


###
######################################################




######################################################
### FUNCTIONS                                      ###
######################################################
###

create_object() {
  version=$1
  protocol=$2
  ipversion=$3
  filtertype=$4
  input=$5
  output=$6

  input=$(sed "s/\$version/$version/g" <<< $input)
  input=$(sed "s/\$filtertype/$filtertype/g" <<< $input)

  cat $input >> $output
}

create_object_sequence() {
  version=$1
  protocol=$2
  ipversion=$3
  filtertype=$4
  input=$5
  output=$6

  maxCount=0
  if [[ $ipversion == "ipv4" ]]; then
    maxCount=32
  else
    maxCount=128
  fi

  for ((i="$maxCount";i>0;i--))
  do
      sequence=$(printf "%03d" $i)
      input=$(sed "s/\$version/$version/g" <<< $input)
      input=$(sed "s/\$filtertype/$filtertype/g" <<< $input)

      cat $input | sed "s/\$sequence/$sequence/g" >> $output
  done
}

create_objectbindings_sequence() {
  version=$1
  protocol=$2
  ipversion=$3
  filtertype=$4
  input=$5
  output=$6

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

      input=$(sed "s/\$version/$version/g" <<< $input)
      input=$(sed "s/\$filtertype/$filtertype/g" <<< $input)

      tempFilename="$input"'_$policyid'
      tempFilename=$(sed "s/\$version/$version/g" <<< $tempFilename)
      tempFilename=$(sed "s/\$filtertype/$filtertype/g" <<< $tempFilename)
      tempFilename=$(sed "s/\$policyid/$policyid/g" <<< $tempFilename)

      cat $input | sed "s/\$sequence/$sequence/g" >> $tempFilename
      cat $tempFilename | sed "s/\$policyid/$policyid/g" >> $output
        
      rm $tempFilename
      ((policyid_sequence++))
  done
}

###

replace_ipversion() {
  ipversion=$1
  output=$2

  upperIpversion=$(echo $ipversion | tr '[:lower:]' '[:upper:]')

  sed -i "s/\$ipversion/$ipversion/g" $output
  sed -i "s/\$IPVERSION/$upperIpversion/g" $output
}

###

replace_protocol() {
  protocol=$1
  output=$2

  upperProtocol=$(echo $protocol | tr '[:lower:]' '[:upper:]')

  sed -i "s/\$protocol/$protocol/g" $output
  sed -i "s/\$PROTOCOL/$upperProtocol/g" $output
}

###

replace_filtertype() {
  filtertype=$1
  output=$2

  upperFiltertype=$(echo $filtertype | tr '[:lower:]' '[:upper:]')

  sed -i "s/\$filtertype/$filtertype/g" $output
  sed -i "s/\$FILTERTYPE/$upperFiltertype/g" $output
}

###

replace_bindingprefix() {
  ipversion=$1
  bindingprefix=$2
  output=$3

  bindingprefix="$bindingprefix${ipversion: -1}"

  sed -i "s/\$bindingprefix/$bindingprefix/g" $output
}

###
######################################################

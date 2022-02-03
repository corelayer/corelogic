#!/bin/sh
clean_file() {
    if [ -f $1 ]
    then
        echo "Removing $1"
        rm $1
    fi
}

clean_file config.txt
clean_file config.json
clean_file config.out
clean_file config.conf

clean_file output.txt
clean_file output.conf.out


clean_file assets/framework/11.0/packages/contentswitching/fake/csv_ipv4_ipfilter_blocklist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/csv_ipv6_ipfilter_blocklist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/csvgroup_ipv4_ipfilter_blocklist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/csvgroup_ipv6_ipfilter_blocklist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/tenant_ipv4_ipfilter_blocklist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/tenant_ipv6_ipfilter_blocklist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/endpoint_ipv4_ipfilter_blocklist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/endpoint_ipv6_ipfilter_blocklist.yaml

clean_file assets/framework/11.0/packages/contentswitching/fake/csv_ipv4_ipfilter.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/csv_ipv6_ipfilter.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/csvgroup_ipv4_ipfilter.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/csvgroup_ipv6_ipfilter.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/tenant_ipv4_ipfilter.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/tenant_ipv6_ipfilter.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/endpoint_ipv4_ipfilter.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/endpoint_ipv6_ipfilter.yaml

sleep 5
#!/bin/sh
clean_file() {
    if [ -f $1 ]
    then
        echo "Removing $1"
        rm $1
    fi

    if [ -d $1 ]
    then
        echo "Removing $1"
        rm -rf $1
    fi
}

clean_file output
#clean_file output/config.conf
#clean_file output/output.txt
#clean_file output/output_docker.txt

clean_file assets/framework/11.0/packages/contentswitching/fake/csv_ipv4_ipfilter_blocklist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/csv_ipv6_ipfilter_blocklist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/csvgroup_ipv4_ipfilter_blocklist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/csvgroup_ipv6_ipfilter_blocklist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/tenant_ipv4_ipfilter_blocklist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/tenant_ipv6_ipfilter_blocklist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/endpoint_ipv4_ipfilter_blocklist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/endpoint_ipv6_ipfilter_blocklist.yaml

clean_file assets/framework/11.0/packages/contentswitching/fake/csv_ipv4_ipfilter_allowlist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/csv_ipv6_ipfilter_allowlist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/csvgroup_ipv4_ipfilter_allowlist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/csvgroup_ipv6_ipfilter_allowlist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/tenant_ipv4_ipfilter_allowlist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/tenant_ipv6_ipfilter_allowlist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/endpoint_ipv4_ipfilter_allowlist.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/endpoint_ipv6_ipfilter_allowlist.yaml

clean_file assets/framework/11.0/packages/contentswitching/fake/csv_ipv4_ipfilter.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/csv_ipv6_ipfilter.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/csvgroup_ipv4_ipfilter.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/csvgroup_ipv6_ipfilter.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/tenant_ipv4_ipfilter.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/tenant_ipv6_ipfilter.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/endpoint_ipv4_ipfilter.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/endpoint_ipv6_ipfilter.yaml


clean_file assets/framework/11.0/packages/contentswitching/fake/init_ipfilter_ipv4.yaml
clean_file assets/framework/11.0/packages/contentswitching/fake/init_ipfilter_ipv6.yaml

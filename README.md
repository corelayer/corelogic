corelogic list
corelogic compile

corelogic install
corelogic uninstall

corelogic gui




CS
id                  name

tenant              name
csgroup             name
status              enabled | disabled | maintenance
ipfilter            disabled | allow | block


CSGROUP
id                  name

tenant              name
status              enabled | disabled | maintenance
ipfilter            disabled | allow | block


TENANT
id                  name

status              enabled | disabled | maintenance
ipfilter            disabled | allow | block


ENDPOINT
ip                  1.2.3.4
port                80
protocol            http | ssl | tcp | udp | ssl_tcp | ssl_bridge

tenant              name
state               enabled | disabled | maintenance
cs



TENANT_IPFILTER
tenant              name
cidr                10.0.0.0/8

description


CSVGROUP_IPFILTER
csgroup             name
cidr                10.0.0.0/8

description


CSV_IPFILTER
cs                  name
cidr                10.0.0.0/8

description


LB_IPFILTER
lb                  name
cidr                10.0.0.0/8

description
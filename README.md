# CoreLogic

## Introduction
CoreLogic is a configuration framework for Citrix ADC/NetScaler, compatible with all editions (Standard, Advanced/Enterprise, Premium/Platinum), aimed at standardising the application flow through the packet engine. More specifically, controlling access applications from L3 to L7 by extensive use of content-switching virtual servers.

## Goals
- Standarization:
  - Multiple engineers working on one configuration are forced to configure policies in the same way
  - You can quickly get your bearing at someone else's configuration without having to rely heavily on documentation
- Documentation:
  - It is easier to document and graph how an application is configured
  - Focus on application-specific documentation
- Simplification:
  - All "moving parts" are defined in a set of tables (key-value pairs)
- Security:
  - All application (modules) must be configured explicitly. If the application (or application module) is not defined, it will not be allowed to pass through.
  - Block access to application (modules) at L4-L7 based on client IP address

---

## Features
### Access Zones
CoreLogic has knowledge of two zones, LAN and ANY, allowing you to define how a request will be handled if you're accessing the service from an internal network (LAN) or from the outside (ANY). This decision is based on the entries in a table, defining which networks are to be considered an "internal network".

### IP Whitelist/Blacklist
CoreLogic also implements a basic ACL by allowing you to define whether content-switching virtual servers and load-balancing virtual servers have an IP whitelist or blacklist, thus restricting access to specific services. This decision is based on the entries in a table, defining which networks belong the the whitelist/blacklist.

Using the IP whitelist/blacklist functionality in combination with the content-switching flow control enables you to define complex access scenarios based on IP address in combination with L4-L7 data.

### Web Application (module) granularity
Given a URI `https://www.netscalerrocks.com/packetengine/internals/documentation?section=rewrite`, we can identify specific components:

- Protocol: https
- FQDN (hostname): www.netscalerrocks.com
- URL: /packetengine/internals/documentation.php?section=rewrite
  - URL Path: /packetengine/internals/documentation.php
  - URL Query: section=rewrite

CoreLogic can differentiate following scenarios:

| Scenario | Elements | Example |
|-|-|-|
| **FULL** | FQDN + URL Path | www.netscalerrocks.com/packetengine/internals/documentation.php |
| **SCND** | FQDN + URL Path (first two elements) | www.netscalerrocks.com/packetengine/internals |
| **FRST** | FQDN + URL Path (first element only) | www.netscalerrocks.com/packetengine |
| **FQDN** | FQDN | www.netscalerrocks.com |
| **WILD** | Wildcarded domain | *.netscalerrocks.com

For each of the scenario's above, you can target different load-balancing virtual servers to access, redirect requests or block access.

| Target Scenario | Action |
|-|-|
| **Load-balancing virtual server** | Pass the request to a load-balancing virtual server |
| **VS_REDIR_301** | 301 Redirect to a defined location |
| **VS_REDIR_302** | 302 Redirect to a defined location |
| **VS_REDIR_302_SWITCH** | 302 Redirect from http to https (or vice versa) on the same URI |
| **VS_BLOCKED** | Show a message that the request is blocked |
| **VS_DROP** | Drop the request silently at the TCP-connection |
| **VS_RESET** | Do a hard reset on the TCP-connection |

*Notes:*
- *In case of a FQDN like customer.api.netscalerrocks.com, the wildcarded domain becomes \*.api.netscalerrocks.com*
- *CoreLogic does not process URL queries or HTTP headers such as User-Agent*

---

## Install Procedure

Open up a SSH connection to your Citrix ADC/NetScaler appliance, and copy-paste the contents of install.conf.

Alternatively you might want to copy the contents of install.conf into a configuration job template on Citrix ADM, and execute it from there.

## Uninstall Procedure

None so far, it's a manual action due to dependencies.
If any policy is still in use, it will fail to uninstall.

---

## Content-Switching Virtual Servers
### Initialization
To create a new set of content-switching virtual servers, use the script in **create_cs.conf**.

1. Replace `$TENANT` with a name of your choice, e.g. PUB012.
2. Replace `$IPADDRESS` with the actual IP address for the virtual server.

**Note:** It is important that the key for SM_IP_CONTROL is in **LOWERCASE**: cs_$tenant_http --> cs_pub012_http.


For example:
```
 add cs vserver CS_PUB012_HTTP HTTP 192.168.0.12 80 -cltTimeout 180
bind cs vserver CS_PUB012_HTTP -policyName NOPOLICY-REWRITE -priority 10601 -gotoPriorityExpression END -type REQUEST -invoke policylabel RWPL_CL10_6_CS_REQ_CORE
bind cs vserver CS_PUB012_HTTP -policyName NOPOLICY-REWRITE -priority 10601 -gotoPriorityExpression END -type RESPONSE -invoke policylabel RWPL_CL10_6_CS_RES_CORE
bind cs vserver CS_PUB012_HTTP -policyName RSP_CL10_6_CSTCP_BLOCKED_NOT_LISTED -priority 10601 -gotoPriorityExpression END -type REQUEST
bind cs vserver CS_PUB012_HTTP -policyName RSP_CL10_6_VSTCP_BLOCKED_NOT_LISTED -priority 10602 -gotoPriorityExpression END -type REQUEST
bind cs vserver CS_PUB012_HTTP -policyName RSP_CL10_6_CSTCP_BLOCKED_WHITE -priority 10603 -gotoPriorityExpression END -type REQUEST
bind cs vserver CS_PUB012_HTTP -policyName RSP_CL10_6_CSTCP_BLOCKED_BLACK -priority 10604 -gotoPriorityExpression END -type REQUEST
bind cs vserver CS_PUB012_HTTP -policyName RSP_CL10_6_VSTCP_BLOCKED_WHITE -priority 10605 -gotoPriorityExpression END -type REQUEST
bind cs vserver CS_PUB012_HTTP -policyName RSP_CL10_6_VSTCP_BLOCKED_BLACK -priority 10606 -gotoPriorityExpression END -type REQUEST
bind cs vserver CS_PUB012_HTTP -policyName CSP_CL10_6_FULL_LAN -priority 10601
bind cs vserver CS_PUB012_HTTP -policyName CSP_CL10_6_SCND_LAN -priority 10602
bind cs vserver CS_PUB012_HTTP -policyName CSP_CL10_6_FRST_LAN -priority 10603
bind cs vserver CS_PUB012_HTTP -policyName CSP_CL10_6_FQDN_LAN -priority 10604
bind cs vserver CS_PUB012_HTTP -policyName CSP_CL10_6_WILD_LAN -priority 10605
bind cs vserver CS_PUB012_HTTP -policyName CSP_CL10_6_FULL_ANY -priority 10611
bind cs vserver CS_PUB012_HTTP -policyName CSP_CL10_6_SCND_ANY -priority 10612
bind cs vserver CS_PUB012_HTTP -policyName CSP_CL10_6_FRST_ANY -priority 10613
bind cs vserver CS_PUB012_HTTP -policyName CSP_CL10_6_FQDN_ANY -priority 10614
bind cs vserver CS_PUB012_HTTP -policyName CSP_CL10_6_WILD_ANY -priority 10615
bind cs vserver CS_PUB012_HTTP -lbvserver VS_NO_SERVICE_HTTP
bind policy stringmap SM_IP_CONTROL cs_pub012_http "list=blacklist;"
```

### Configuration
#### Whitelist/Blacklist
By default, all content-switching virtual servers are set to have a blacklist of IP addresses.
This means that all source IP addresses are allowed to access the content-switching virtual server, unless there is an entry in `SM_IP_CONTROL`.
To change the behavior of the content-switching virtual to be a whitelist, all you need to do is change the entry.

If you change the behavior to be a `whitelist`, it results in all client IP addresses on the list to be blocked!

For example:
```
bind policy stringmap SM_IP_CONTROL cs_pub012_http "list=whitelist;"
```

**Notes**:
- As stated before, it is important that the key for SM_IP_CONTROL is in lowercase: **cs_pub012_http**.
- It is equally important not to omit the semicolon at the end as policies are looking for the value between `=` and `;` to determine their action.

#### IP addresses on the whitelist/blacklist
To add IP addresses or complete networks to the list, we need to provide additional entries in `SM_IP_CONTROL`.

Example:
- Content-switching virtual server `CS_PUB012_HTTP` can only provide access to a specified list IP addresses.
- All clients from the sales network `192.168.0.0/24` must be allowed.
- All clients from the development network `172.16.0.0/16` must be allowed.
- An administrator with IP address `10.0.0.1` must be allowed.

```
bind policy stringmap SM_IP_CONTROL cs_pub012_http "list=whitelist;"
bind policy stringmap SM_IP_CONTROL "cs_pub012_http;any;192.168.0.0/24" "Sales"
bind policy stringmap SM_IP_CONTROL "cs_pub012_http;any;172.16.0.0/16" "Development"
bind policy stringmap SM_IP_CONTROL "cs_pub012_http;any;10.0.0.1/32" "Administrator"
```

Example:
- Content-switching virtual server `CS_PUB012_HTTP` was defined to have a blacklist.
- Public IP addresses from Google DNS must be blocked.
- All clients from the sales network `192.168.0.0/24` must be blocked.

```
bind policy stringmap SM_IP_CONTROL cs_pub012_http "list=blacklist;"
bind policy stringmap SM_IP_CONTROL "cs_pub012_http;any;8.8.8.8/32" "Google DNS"
bind policy stringmap SM_IP_CONTROL "cs_pub012_http;any;8.8.4.4/32" "Google DNS"
bind policy stringmap SM_IP_CONTROL "cs_pub012_http;any;192.168.0.0/24" "Sales"
```

#### LAN Networks


### Flow
Module processing on the Request:
1. Content-Switching policies: determine which load-balancing virtual server to use
  1. Policies for clients on LAN networks (LAN)
  2. Policies for clients not on LAN networks (ANY)
2. Responder policies: based on the current content-switching virtual server and selected load-balancing virtual server, check the whitelist/blacklist.
3. Rewrite policies: adds/removes some headers to be used by the backend server, such as X-Forwarded-For and X-Forwarded-Proto.
4. Pass the request to the selected load-balancing virtual server.

Module processing on the Response:
1. Rewrite policies: adds/remove some headers for improved security.

#### Content-Switching Policies

---
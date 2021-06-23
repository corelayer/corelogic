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

### IP Allowlist/Blocklist
CoreLogic also implements a basic ACL by allowing you to define whether content-switching virtual servers and load-balancing virtual servers have an IP allowlist or blocklist, thus restricting access to specific services. This decision is based on the entries in a table, defining which networks belong the the allowlist/blocklist.

Using the IP allowlist/blocklist functionality in combination with the content-switching flow control enables you to define complex access scenarios based on IP address in combination with L4-L7 data.

### Web Application (module) granularity
Given a URI `https://www.netscalerrocks.com/packetengine/internals/documentation.php?section=rewrite`, we can identify specific components:

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
| **VS_REDIR_301_KEEPPATH** | 301 Redirect to a defined location while appending the original path to the redirect destination |
| **VS_REDIR_301_SWITH** | 301 Redirect from http to https (or vice versa) on the same URI |
| **VS_REDIR_302** | 302 Redirect to a defined location |
| **VS_REDIR_302_KEEPPATH** | 301 Redirect to a defined location while appending the original path to the redirect destination |
| **VS_REDIR_302_SWITCH** | 302 Redirect from http to https (or vice versa) on the same URI |
| **VS_REDIR_307** | 307 Redirect to a defined location |
| **VS_REDIR_307_KEEPPATH** | 307 Redirect to a defined location while appending the original path to the redirect destination |
| **VS_REDIR_307_SWITH** | 307 Redirect from http to https (or vice versa) on the same URI |
| **VS_REDIR_308** | 308 Redirect to a defined location |
| **VS_REDIR_308_KEEPPATH** | 308 Redirect to a defined location while appending the original path to the redirect destination |
| **VS_REDIR_308_SWITCH** | 308 Redirect from http to https (or vice versa) on the same URI |
| **VS_NOTFOUND_HTTP** | Show a message that the page cannot be found |
| **VS_BLOCKED_HTTP** | Show a message that the request is blocked |
| **VS_DROP_HTTP** | Drop the request silently on L4 |
| **VS_DROP_TCP** | Drop the request silently on L4 |
| **VS_RESET_HTTP** | Do a hard reset on L4 |
| **VS_RESET_TCP** | Do a hard reset on L4 |
| **VS_ACME_HTTP** | Respond to ACME-challenge |


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
To create a new set of content-switching virtual servers, use the script in [create_cs.conf](https://github.com/CoreLayer/CoreLogic/blob/master/create_cs.conf).

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
bind cs vserver CS_PUB012_HTTP -policyName RSP_CL10_6_CSTCP_BLOCKED_ALLOW -priority 10603 -gotoPriorityExpression END -type REQUEST
bind cs vserver CS_PUB012_HTTP -policyName RSP_CL10_6_CSTCP_BLOCKED_BLOCK -priority 10604 -gotoPriorityExpression END -type REQUEST
bind cs vserver CS_PUB012_HTTP -policyName RSP_CL10_6_VSTCP_BLOCKED_ALLOW -priority 10605 -gotoPriorityExpression END -type REQUEST
bind cs vserver CS_PUB012_HTTP -policyName RSP_CL10_6_VSTCP_BLOCKED_BLOCK -priority 10606 -gotoPriorityExpression END -type REQUEST
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
bind policy stringmap SM_IP_CONTROL cs_pub012_http "list=blocklist;"
```

### Configuration
#### Allowlist/Blocklist
By default, all content-switching virtual servers are set to have a blocklist of IP addresses.
This means that all source IP addresses are allowed to access the content-switching virtual server, unless there is an entry in `SM_IP_CONTROL`.
To change the behavior of the content-switching virtual to be a allowlist, all you need to do is change the entry.

If you change the behavior to be a `allowlist`, it results in all client IP addresses on the list to be blocked!

For example:
```
bind policy stringmap SM_IP_CONTROL cs_pub012_http "list=allowlist;"
```

*Notes:*
- *As stated before, it is important that the key for SM_IP_CONTROL is in lowercase: **cs_pub012_http**.*
- *It is equally important not to omit the semicolon at the end as policies are looking for the value between `=` and `;` to determine their action.*

#### IP addresses on the allowlist/blocklist
To add IP addresses or complete networks to the list, we need to provide additional entries in `SM_IP_CONTROL`.

Example:
- Content-switching virtual server `CS_PUB012_HTTP` can only provide access to a specified list IP addresses.
- All clients from the sales network `192.168.0.0/24` must be allowed.
- All clients from the development network `172.16.0.0/16` must be allowed.
- An administrator with IP address `10.0.0.1` must be allowed.

```
bind policy stringmap SM_IP_CONTROL cs_pub012_http "list=allowlist;"
bind policy stringmap SM_IP_CONTROL "cs_pub012_http;any;192.168.0.0/24" "Sales"
bind policy stringmap SM_IP_CONTROL "cs_pub012_http;any;172.16.0.0/16" "Development"
bind policy stringmap SM_IP_CONTROL "cs_pub012_http;any;10.0.0.1/32" "Administrator"
```

Example:
- Content-switching virtual server `CS_PUB012_HTTP` was defined to have a blocklist.
- Public IP addresses from Google DNS must be blocked.
- All clients from the sales network `192.168.0.0/24` must be blocked.

```
bind policy stringmap SM_IP_CONTROL cs_pub012_http "list=blocklist;"
bind policy stringmap SM_IP_CONTROL "cs_pub012_http;any;8.8.8.8/32" "Google DNS"
bind policy stringmap SM_IP_CONTROL "cs_pub012_http;any;8.8.4.4/32" "Google DNS"
bind policy stringmap SM_IP_CONTROL "cs_pub012_http;any;192.168.0.0/24" "Sales"
```

#### LAN Networks
Configuring LAN networks is a similar procedure, but there are some differences to allowlisting/blocklisting.
Whilst allowlisting/blocklisting happens for each and every specific content-switching virtual server, LAN networks apply to a collection of content-switching virtual servers with the same name.

Assume we have executed the complete create_cs.conf script to create the following content-switching virtual servers:
- CS_PUB012_HTTP
- CS_PUB012_SSL
- CS_PUB012_TCP
- CS_PUB012_STCP
- CS_PUB012_UDP

If we want to define a LAN network for CS_PUB012_HTTP, it will be considered a LAN network for all content-switching virtual servers of PUB012.

Example:
- All clients from the sales network `192.168.0.0/24` are considered to be internal clients on the LAN.
- All clients from the development network `172.16.0.0/16` are considered to be internal clients on the LAN.
- An administrator with IP address `10.0.0.1` is considered to be an internal client on the LAN.

```
bind policy stringmap SM_IP_CONTROL "cs_pub012;lan;192.168.0.0/24" "Sales"
bind policy stringmap SM_IP_CONTROL "cs_pub012;lan;172.16.0.0/16" "Development"
bind policy stringmap SM_IP_CONTROL "cs_pub012;lan;10.0.0.1/32" "Administrator"
```

As you can see from the example above, the protocol of the content-switching virtual servers is omitted in the key of the entry.

### Execution flow
Module processing on the Request:
1. Content-Switching policies to determine which load-balancing virtual server to use
  1. Policies for clients on LAN networks (LAN)
  2. Policies for clients not on LAN networks (ANY)
2. Responder policies: based on the current content-switching virtual server and selected load-balancing virtual server, check the allowlist/blocklist.
3. Rewrite policies: adds/removes some headers to be used by the backend server, such as `X-Forwarded-For` and `X-Forwarded-Proto`.
4. Pass the request to the selected load-balancing virtual server.

Module processing on the Response:
1. Rewrite policies: adds/remove some headers for improved security.

#### Fallback flows
##### HTTP/SSL
If the selected load-balacing virtual server is `OUT OF SERVICE` or `DOWN`, content-switching policy processing will stop and the default load-balancing virtual server will be selected immediately.

- For HTTP, `VS_NO_SERVICE_HTTP` will be used
- For SSL, `VS_NO_SERVICE_SSL` will be used.

Both have the same functionality built-in:
- If an entry in `SM_CS_CONTROL` should have been used for the current protocol (HTTP/SSL), we know that the target load-balancing virtual server is down and a `NO SERVICE` message will be shown. The HTTP response code is 503.
- If an entry in `SM_CS_CONTROL` is not found for the current request on HTTP, we know that the request is not allowed. We will lookup if there is an entry for SSL, and redirect if and entry is found.
- If an entry in `SM_CS_CONTROL` is not found for the current request on HTTP, nor for SSL traffic, we will respond with a `NOT FOUND` message. and HTTP response code 404.

##### TCP/SSL_TCP/UDP
- For general TCP/UDP content-switching virtual servers, the connection will time out.
---

## Applications
Having a fancy framework is one thing, but what it's all about are the applications and their modules.

Although you are here to use CoreLogic (or at least get to know it), you still have to complete some necessary steps which are not related to using CoreLogic:
- Create server objects
- (Create a monitor)
- Create a service group
- Create a load-balancing virtual server

The only object which is addressed by CoreLogic, is the load-balancing virtual server.

To illustrate the process, we are going to deploy Wordpress through Citrix NetScaler with CoreLogic.

### Example application
#### Server
Wordpress will be running a server with IP address `192.168.1.10`.

```
add server SRV_WORDPRESS 192.168.1.10
```

#### Service Group
Wordpress is deployed on the server without SSL, so we can use `HTTP` as the Service Group Type and `tcp/80` as the destination port.

```
 add servicegroup SG_WORDPRESS HTTP
bind servicegroup SG_WORDPRESS SRV_WORDPRESS 80
```

#### Load-Balancing Virtual Server
Now things get interesting.

First of all, let's create the load-balancing virtual server. As the load-balancing virtual server will only be addressed through a content-switching virtual server, it doesn't need to have an IP address (non-addressable). Also, if you are running your site on SSL, the encryption will be taken off by the content-switching virtual server (SSL Offloading), so the load-balancing virtual server can be of type HTTP.

```
 add lb vserver VS_WORDPRESS HTTP 0.0.0.0 0
bind lb vserver VS_WORDPRESS SG_WORDPRESS
```

Now that we have created the load-balancing virtual server for our application, we have to configure it for allowlisting/blocklisting. Remember, the execution flow on the content-switching virtual server checks whether the selected load-balancing virtual server has a allowlist or blocklist configured. If it is not configured, you will get a connection reset.

```
bind policy stringmap SM_IP_CONTROL vs_wordpress "list=blocklist;"
```

Good, almost there!
The only thing left is configuring when the content-switching virtual server has to forward our request to the load-balancing virtual server.
For the purpose of this example, we will deploy our Wordpress application on the hostname `www.netscalerrocks.com`.

We need to add entry to `SM_CS_CONTROL` to specify the content-switching scenarios.
Format for these entries is as follows:
| Key (all in lowercase) | Value |
| - | - |
| "<cs_name>;<lan\|any>;<request>" | "vs=<lb_name>;dst=<redirect_location_if_applicable>;" |

```
bind policy stringmap SM_CS_CONTROL "cs_pub012_ssl;any;www.netscalerrocks.com" "vs=VS_WORDPRESS;"
```

To illustrate a redirect, let's say we have another domain `netscalerrocks.org` we'd like to redirect to `www.netscalerrocks.com`.However `netscalerrocks.org/nitro` needs to go to `api.netscalerrocks.com/nitro`, which means we want to keep the full path of the original request if it starts with `/nitro`.

```
bind policy stringmap SM_CS_CONTROL "cs_pub012_ssl;any;netscalerrocks.org" "vs=VS_REDIR_301;dst=//www.netscalerrocks.com;"
bind policy stringmap SM_CS_CONTROL "cs_pub012_ssl;any;netscalerrocks.org/nitro" "vs=VS_REDIR_302_KEEPPATH;dst=//api.netscalerrocks.com"
```

As a final step, we only want to allow the administration pages from the internal network. This would translate in the following commands:

```
bind policy stringmap SM_CS_CONTROL "cs_pub012_ssl;lan;www.netscalerrocks.com" "vs=VS_WORDPRESS;"
bind policy stringmap SM_CS_CONTROL "cs_pub012_ssl;any;www.netscalerrocks.com" "vs=VS_WORDPRESS;"
bind policy stringmap SM_CS_CONTROL "cs_pub012_ssl;any;www.netscalerrocks.com/wp-admin" "vs=VS_RESET;"
```

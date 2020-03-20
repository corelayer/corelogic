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
Given a URI https://www.netscalerrocks.com/packetengine/internals/documentation?section=rewrite, we can identify specific components:

- Protocol: https
- FQDN (hostname): www.netscalerrocks.com
- URL: /packetengine/internals/documentation.php?section=rewrite
  - URL Path: /packetengine/internals/documentation.php
  - URL Query: section=rewrite

CoreLogic can differentiate following scenarios:

| Scenario | Elements | Example |
|-|-|-|
| FULL | FQDN + URL Path | www.netscalerrocks.com/packetengine/internals/documentation.php |
| SCND | FQDN + URL Path (first two elements) | www.netscalerrocks.com/packetengine/internals |
| FRST | FQDN + URL Path (first element only) | www.netscalerrocks.com/packetengine |
| FQDN | FQDN | www.netscalerrocks.com |
| WILD | Wildcarded domain | *.netscalerrocks.com

**Notes**:
- In case of a FQDN like customer.api.netscalerrocks.com, the wildcarded domain becomes *.api.netscalerrocks.com
- CoreLogic does not process URL queries or HTTP headers such as User-Agent

---

## Install Procedure

Open up a SSH connection to your Citrix ADC/NetScaler appliance, and copy-paste the contents of install.conf.

Alternatively you might want to copy the contents of install.conf into a configuration job template on Citrix ADM, and execute it from there.

## Uninstall Procedure

None so far, it's a manual action due to dependencies.
If any policy is still in use, it will fail to uninstall.

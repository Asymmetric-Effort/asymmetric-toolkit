DnsProxy
========

## Purpose:
To create a DNS Proxy which can distribute DNS queries from a given
source using a weighted round-robin algorithm to a set of target DNS 
servers.

## Use Case:
When using dnsEnum, many DNS queries will be generated.  In fact this 
could overwhelm a target DNS server and cause an accidental outage.  To
mitigate that possibility, the DNS Proxy allows the queries to be 
distributed across a fleet of target DNS servers.

## Usage:
```bash
dnsProxy --listen <listenerConfig> --dnsServers <serverList>

    listenerConfig = <protocol>:(<ipv4addr>|<ipv6addr>):<portNumber>
    
    serverList=targetDnsServer(0),targetDnsServer(1),...,targetDnsServer(n)
    
    targetDnsServer = <weight>:<protocol>:(<ipv4addr>|<ipv6addr>):<portNumber>
    
    weight=0-9
    
    protocol=tcp|udp

    portNumber={1...65535}
```
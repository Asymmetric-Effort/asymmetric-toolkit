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


#!/bin/bash

# The task is to find out which domains are on Cloudflare nameservers. This is
# the first thought how to do it.

head testdata/domains.txt | xargs -I{} dig NS {} | grep -E 'IN\s+NS.*\.$'
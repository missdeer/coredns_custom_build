#!/bin/bash
if [[ $EUID -ne 0 ]]; then
   echo "This script must be run as root"
   exit 1
fi

setcap CAP_NET_RAW,CAP_NET_ADMIN,CAP_NET_BIND_SERVICE+ep ./coredns

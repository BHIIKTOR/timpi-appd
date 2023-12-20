#!/bin/sh

if test -n "$1"; then
    # need -R not -r to copy hidden files
    cp -R "$1/.timpid" /root
fi

mkdir -p /root/log
timpid start --rpc.laddr tcp://0.0.0.0:26657 --trace

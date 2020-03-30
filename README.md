# CoreDNS custom build

[![Build status](https://ci.appveyor.com/api/projects/status/e2y1n3k3wwiei0bs?svg=true)](https://ci.appveyor.com/project/missdeer/coredns-custom-build)

# Added plugins

## official plugins

1. [proxy](https://github.com/coredns/proxy) 
2. [fallback](https://github.com/coredns/fallback) with some modification for build error
3. [forward](https://github.com/coredns/coredns/tree/master/plugin/forward) with some modification for speed up except list lookup

## third party plugins

1. [ads](https://github.com/c-mueller/ads) 
2. [bogus](https://github.com/missdeer/bogus)
3. [ipset](https://github.com/missdeer/ipset)
4. [redisc](https://github.com/miekg/redis)

# Download Prebuilt Binaries

Download from [Appveyor artifacts](https://ci.appveyor.com/project/missdeer/coredns-custom-build).

# Configuration

## GUI Configuration tool for GUI environment

[CoreDNS Home](https://github.com/missdeer/corednshome)  screenshot:

![main window](https://raw.githubusercontent.com/missdeer/corednshome/master/screenshots/mainwindow.png)

## Shell script for UNIX-like environment

Check [bash script](https://gist.github.com/missdeer/5c7c82b5b67f8afb41cfd43d51b82c2d) for generating configuration file. 

# More Information

1. [增强版CoreDNS，上网更科学](https://blog.minidump.info/2019/12/enhanced-coredns/)
2. [CoreDNS搭建无污染DNS](https://blog.minidump.info/2019/07/coredns-no-dns-poisoning/)
3. [使用Prometheus观察CoreDNS运行状况](https://blog.minidump.info/2020/03/prometheus-for-coredns/)

# Donate

![微信扫一扫](https://raw.githubusercontent.com/missdeer/corednshome/master/src/res/wepay.jpg)  ![支付宝扫一扫](https://raw.githubusercontent.com/missdeer/corednshome/master/src/res/alipay.jpg)

[![paypal donate](https://raw.githubusercontent.com/missdeer/corednshome/master/paypal-donate.png)](https://www.paypal.me/dfordsoft/)

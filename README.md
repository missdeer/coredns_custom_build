# CoreDNS custom build

[![Build status](https://ci.appveyor.com/api/projects/status/e2y1n3k3wwiei0bs?svg=true)](https://ci.appveyor.com/project/missdeer/coredns-custom-build)

# Added plugins

## official plugins

1. [proxy](https://github.com/missdeer/proxy) 
2. [fallback](https://github.com/missdeer/fallback) with some modification for build error
3. [forward](https://github.com/coredns/coredns/tree/master/plugin/forward) with some modification for speed up except list lookup

## third party plugins

1. [bogus](https://github.com/missdeer/bogus)
2. [ipset](https://github.com/missdeer/ipset)
3. [blocklist](https://github.com/relekang/coredns-blocklist)
4. [dnsredir](https://github.com/leiless/dnsredir)
5. [ads](https://github.com/c-mueller/ads) 

# Download Prebuilt Binaries

All prebuilt binaries are built by [Appveyor service](https://ci.appveyor.com/project/missdeer/coredns-custom-build), you may choose the right binary for your platform from the list shown below:

| OS           | Arch    | Option     | Link                                                                |
|--------------|---------|--------    |---------------------------------------------------------------------|
| Windows      | x86_64  |            | https://coredns.minidump.info/dl/coredns-windows-amd64.zip             |
| Windows      | x86     |            | https://coredns.minidump.info/dl/coredns-windows-386.zip               |
| Windows      | arm     |            | https://coredns.minidump.info/dl/coredns-windows-arm.zip               |
| Windows      | arm64   |            | https://coredns.minidump.info/dl/coredns-windows-arm64.zip             |
| macOS        | x86_64  |            | https://coredns.minidump.info/dl/coredns-darwin-amd64.zip              |
| macOS        | arm64   |            | https://coredns.minidump.info/dl/coredns-darwin-arm64.zip              |
| Linux        | x86     |            | https://coredns.minidump.info/dl/coredns-linux-386.zip                 |
| Linux        | x86_64  |            | https://coredns.minidump.info/dl/coredns-linux-amd64.zip               |
| Linux        | arm     |  5         | https://coredns.minidump.info/dl/coredns-linux-armv5.zip               |
| Linux        | arm     |  6         | https://coredns.minidump.info/dl/coredns-linux-armv6.zip               |
| Linux        | arm     |  7         | https://coredns.minidump.info/dl/coredns-linux-armv7.zip               |
| Linux        | arm64   |            | https://coredns.minidump.info/dl/coredns-linux-arm64.zip               |
| Linux        | ppc64   |            | https://coredns.minidump.info/dl/coredns-linux-ppc64.zip               |
| Linux        | ppc64le |            | https://coredns.minidump.info/dl/coredns-linux-ppc64le.zip             |
| Linux        | mips64  |  hardfloat | https://coredns.minidump.info/dl/coredns-linux-mips64-hardfloat.zip    |
| Linux        | mips64  |  softfloat | https://coredns.minidump.info/dl/coredns-linux-mips64-softfloat.zip    |
| Linux        | mips64le|  hardfloat | https://coredns.minidump.info/dl/coredns-linux-mips64le-hardfloat.zip  |
| Linux        | mips64le|  softfloat | https://coredns.minidump.info/dl/coredns-linux-mips64le-softfloat.zip  |
| Linux        | mips    |  hardfloat | https://coredns.minidump.info/dl/coredns-linux-mips-hardfloat.zip      |
| Linux        | mips    |  softfloat | https://coredns.minidump.info/dl/coredns-linux-mips-softfloat.zip      |
| Linux        | mipsle  |  hardfloat | https://coredns.minidump.info/dl/coredns-linux-mipsle-hardfloat.zip    |
| Linux        | mipsle  |  softfloat | https://coredns.minidump.info/dl/coredns-linux-mipsle-softfloat.zip    |
| Linux        | s390x   |            | https://coredns.minidump.info/dl/coredns-linux-s390x.zip               |
| FreeBSD      | x86_64  |            | https://coredns.minidump.info/dl/coredns-freebsd-amd64.zip             |
| FreeBSD      | x86     |            | https://coredns.minidump.info/dl/coredns-freebsd-386.zip               |
| FreeBSD      | arm     |            | https://coredns.minidump.info/dl/coredns-freebsd-arm.zip               |
| FreeBSD      | arm64   |            | https://coredns.minidump.info/dl/coredns-freebsd-arm64.zip             |
| NetBSD       | x86_64  |            | https://coredns.minidump.info/dl/coredns-netbsd-amd64.zip              |
| NetBSD       | x86     |            | https://coredns.minidump.info/dl/coredns-netbsd-386.zip                |
| NetBSD       | arm     |            | https://coredns.minidump.info/dl/coredns-netbsd-arm.zip                |
| NetBSD       | arm64   |            | https://coredns.minidump.info/dl/coredns-netbsd-arm64.zip              |
| OpenBSD      | x86_64  |            | https://coredns.minidump.info/dl/coredns-openbsd-amd64.zip             |
| OpenBSD      | x86     |            | https://coredns.minidump.info/dl/coredns-openbsd-386.zip               |
| OpenBSD      | arm     |            | https://coredns.minidump.info/dl/coredns-openbsd-arm.zip               |
| OpenBSD      | arm64   |            | https://coredns.minidump.info/dl/coredns-openbsd-arm64.zip             |
| DragonflyBSD | x86_64  |            | https://coredns.minidump.info/dl/coredns-dragonfly-amd64.zip           |
| Solaris      | x86_64  |            | https://coredns.minidump.info/dl/coredns-solaris-amd64.zip             |
| illumos      | x86_64  |            | https://coredns.minidump.info/dl/coredns-illumos-amd64.zip             |
| Android      | x86_64  |            | https://coredns.minidump.info/dl/coredns-android-amd64.zip             |
| Android      | x86     |            | https://coredns.minidump.info/dl/coredns-android-386.zip               |
| Android      | arm     |            | https://coredns.minidump.info/dl/coredns-android-arm.zip               |
| Android      | arm64   |            | https://coredns.minidump.info/dl/coredns-android-aarch64.zip           |


# Configuration

## Web edition GUI configuration generator, strongly recommended

[https://coredns.minidump.info](https://coredns.minidump.info) screenshot

![web configuration generator](https://raw.githubusercontent.com/missdeer/coredns_custom_build/master/screenshots/web-configuration-generator.png)

## Desktop edition GUI Configuration generator

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

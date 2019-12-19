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

Check [bash script](https://gist.github.com/missdeer/5c7c82b5b67f8afb41cfd43d51b82c2d) for generating configuration file. 

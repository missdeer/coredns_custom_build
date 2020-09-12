#!/bin/bash
pwd
APPVEYOR_BUILD_FOLDER=../coredns_custom_build
cd ..
rm -rf fallback
git clone --depth=1 https://github.com/missdeer/fallback.git
cd fallback
pwd
go mod init github.com/missdeer/fallback
cd ..
rm -rf redis
git clone --depth=1 https://github.com/missdeer/redis.git
cd redis
pwd
go mod init github.com/missdeer/redis
cd ..
rm -rf proxy
git clone --depth=1 https://github.com/missdeer/proxy.git
cd proxy
pwd
go mod init github.com/missdeer/proxy
cd ..
rm -rf ads
git clone --depth=1 https://github.com/c-mueller/ads.git
cd ads
pwd
rm -f go.mod go.sum
go mod init github.com/c-mueller/ads
cd ..
rm -rf dnsredir
git clone --depth=1 https://github.com/leiless/dnsredir.git
cd dnsredir
pwd
rm -f go.mod go.sum
go mod init github.com/leiless/dnsredir
cd ../coredns
pwd
git checkout .
git apply "$APPVEYOR_BUILD_FOLDER/forward.go.patch"
git apply "$APPVEYOR_BUILD_FOLDER/forward-setup.go.patch"
sed -i 's|forward:forward|fallback:github.com/missdeer/fallback\ndnsredir:github.com/leiless/dnsredir\nforward:forward\nproxy:github.com/missdeer/proxy|g' plugin.cfg
sed -i 's|hosts:hosts|ads:github.com/c-mueller/ads\nhosts:hosts|g' plugin.cfg
sed -i 's|cache:cache|cache:cache\nredisc:github.com/missdeer/redis|g' plugin.cfg
sed -i 's|rewrite:rewrite|rewrite:rewrite\nbogus:github.com/missdeer/bogus\nipset:github.com/missdeer/ipset|g' plugin.cfg
echo "replace (" >> go.mod
echo "    github.com/missdeer/fallback => ../fallback" >> go.mod
echo "    github.com/missdeer/redis => ../redis" >> go.mod
echo "    github.com/missdeer/proxy => ../proxy" >> go.mod
echo "    github.com/c-mueller/ads => ../ads" >> go.mod
echo "    github.com/leiless/dnsredir => ../dnsredir" >> go.mod
echo ")" >> go.mod
sed -i '/azure/d' plugin.cfg
sed -i '/route53/d' plugin.cfg

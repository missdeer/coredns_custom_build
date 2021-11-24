#!/bin/bash
cd coredns
git checkout .
git clean -dfx .
git pull
cd ../dnsredir
git pull
cd ../coredns_custom_build
git pull
cd ..
git apply ../coredns_custom_build/forward.go.patch
git apply ../coredns_custom_build/forward-setup.go.patch
sed -i.bak 's|forward:forward|fallback:github.com/missdeer/fallback\ndnsredir:github.com/leiless/dnsredir\nforward:forward\nproxy:github.com/missdeer/proxy|g' plugin.cfg
sed -i.bak 's|hosts:hosts|ads:github.com/c-mueller/ads\nblocklist:github.com/relekang/coredns-blocklist\nhosts:hosts|g' plugin.cfg
sed -i.bak 's|rewrite:rewrite|rewrite:rewrite\nbogus:github.com/missdeer/bogus\nipset:github.com/missdeer/ipset|g' plugin.cfg
sed -i.bak 's|cache:cache|cache:cache\nredisc:github.com/missdeer/redis|g' plugin.cfg
echo "replace (" >> go.mod
echo "    github.com/leiless/dnsredir => ../dnsredir" >> go.mod
echo ")" >> go.mod
sed -i.bak '/azure/d' plugin.cfg
sed -i.bak '/route53/d' plugin.cfg
sed -i.bak '/trace/d' plugin.cfg
sed -i.bak '/etcd/d' plugin.cfg
sed -i.bak '/clouddns/d' plugin.cfg
sed -i.bak '/k8s_external/d' plugin.cfg
sed -i.bak '/kubernetes/d' plugin.cfg
sed -i.bak '/geoip/d' plugin.cfg
sed -i.bak '/nsid/d' plugin.cfg
sed -i.bak '/debug/d' plugin.cfg
sed -i.bak '/trace/d' plugin.cfg
sed -i.bak '/autopath/d' plugin.cfg
sed -i.bak '/erratic/d' plugin.cfg
sed -i.bak '/metadata/d' plugin.cfg
sed -i.bak '/cancel/d' plugin.cfg
sed -i.bak '/pprof/d' plugin.cfg
sed -i.bak '/dnstap/d' plugin.cfg
sed -i.bak '/dns64/d' plugin.cfg
sed -i.bak '/acl/d' plugin.cfg
sed -i.bak '/chaos/d' plugin.cfg
sed -i.bak '/dnssec/d' plugin.cfg
sed -i.bak '/secondary/d' plugin.cfg
sed -i.bak '/loop/d' plugin.cfg
sed -i.bak '/grpc/d' plugin.cfg
sed -i.bak '/sign/d' plugin.cfg

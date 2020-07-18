package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/missdeer/golib/fsutil"
)

var (
	chinaDNSServers = map[string]string{
		"114.114.114.114": "114 DNS",
		"114.114.115.115": "114 DNS",
		"223.5.5.5":       "阿里 AliDNS",
		"223.6.6.6":       "阿里 AliDNS",
		"180.76.76.76":    "百度 BaiduDNS",
		"119.29.29.29":    "DNSPod DNS+",
		"182.254.116.116": "DNSPod DNS+",
		"1.2.4.8":         "CNNIC SDNS",
		"210.2.4.8":       "CNNIC SDNS",
	}

	abroadDNSServers = map[string]string{
		"dns://208.67.222.222:443":  "OpenDNS",
		"dns://208.67.222.222:5353": "OpenDNS",
		"dns://208.67.220.220:443":  "OpenDNS",
		"dns://208.67.220.220:5353": "OpenDNS",
		"tcp://8.8.8.8":             "Google DNS",
		"tcp://8.8.4.4":             "Google DNS",
		"udp://8.8.8.8:443":         "Google DNS",
		"udp://8.8.4.4:443":         "Google DNS",
		"tls://1.1.1.1:853":         "Cloudflare DNS",
		"tls://1.0.0.1:853":         "Cloudflare DNS",
		"tls://8.8.8.8:853":         "Google DNS",
		"tls://8.8.4.4:853":         "Google DNS",
		"tls://9.9.9.9:853":         "IBM Quad9",
		"tls://9.9.9.10:853":        "IBM Quad9",
		"tls://145.100.185.15:853":  "dnsovertls.sinodun.com",
		"tls://145.100.185.16:853":  "dnsovertls1.sinodun.com",
		"tls://145.100.185.17:853":  "dnsovertls2.sinodun.com",
		"tls://145.100.185.18:853":  "dnsovertls3.sinodun.com",
		"tls://185.49.141.37:853":   "getdnsapi.net",
		"tls://89.233.43.71:853":    "unicast.censurfridns.dk",
		"tls://158.64.1.29:853":     "kaitain.restena.lu",
		"tls://176.103.130.130:853": "dns.adguard.com",
		"tls://176.103.130.131:853": "dns.adguard.com",
		"tls://176.103.130.132:853": "dns-family.adguard.com",
		"tls://176.103.130.134:853": "dns-family.adguard.com",
		"tls://199.58.81.218:853":   "dns.cmrg.net",
		"tls://51.15.70.167:853":    "dns.larsdebruin.net",
		"tls://146.185.167.43:853":  "dot.securedns.eu",
		"tls://81.187.221.24:853":   "dns-tls.bitwiseshift.net",
		"tls://94.130.110.185:853":  "ns1.dnsprivacy.at",
		"tls://94.130.110.178:853":  "ns2.dnsprivacy.at",
		"tls://139.59.51.46:853":    "dns.bitgeek.in",
		"tls://89.234.186.112:853":  "dns.neutopia.org",
		"tls://93.177.65.183:853":   "dot1.applied-privacy.net",
		"tls://184.105.193.78:853":  "tls-dns-u.odvr.dns-oarc.net",
		// ipv6
		"tls://[2001:610:1:40ba:145:100:185:15]:853": "dnsovertls.sinodun.com",
		"tls://[2001:610:1:40ba:145:100:185:16]:853": "dnsovertls1.sinodun.com",
		"tls://[2a04:b900:0:100::38]:853":            "getdnsapi.net",
		"tls://[2620:fe::fe]:853":                    "dns.quad9.net",
		"tls://[2620:fe::10]:853":                    "dns.quad9.net",
		"tls://[2606:4700:4700::1111]:853":           "cloudflare-dns.com",
		"tls://[2606:4700:4700::1001]:853":           "cloudflare-dns.com",
		"tls://[2001:4860:4860::8888]:853":           "dns.google",
		"tls://[2001:4860:4860::8844]:853":           "dns.google",
		"tls://[2a00:5a60::ad1:0ff]:853":             "dns.adguard.com",
		"tls://[2a00:5a60::ad2:0ff]:853":             "dns.adguard.com",
		"tls://[2a00:5a60::bad1:0ff]:853":            "dns-family.adguard.com",
		"tls://[2a00:5a60::bad2:0ff]:853":            "dns-family.adguard.com",
		"tls://[2a01:3a0:53:53::0]:853":              "unicast.censurfridns.dk",
		"tls://[2001:a18:1::29]:853":                 "kaitain.restena.lu",
		"tls://[2001:610:1:40ba:145:100:185:18]:853": "dnsovertls3.sinodun.com",
		"tls://[2001:610:1:40ba:145:100:185:17]:853": "dnsovertls2.sinodun.com",
		"tls://[2001:470:1c:76d::53]:853":            "dns.cmrg.net",
		"tls://[2a03:b0c0:0:1010::e9a:3001]:853":     "dot.securedns.eu",
		"tls://[2001:8b0:24:24::24]:853":             "dns-tls.bitwiseshift.net",
		"tls://[2a01:4f8:c0c:3c03::2]:853":           "ns1.dnsprivacy.at",
		"tls://[2a01:4f8:c0c:3bfc::2]:853":           "ns2.dnsprivacy.at",
		"tls://[2001:67c:27e4::35]:853":              "privacydns.go6lab.si",
		"tls://[2a00:5884:8209::2]:853":              "dns.neutopia.org",
		"tls://[2a03:4000:38:53c::2]:853":            "dot1.applied-privacy.net",
		"tls://[2620:ff:c000:0:1::64:25]:853":        "tls-dns-u.odvr.dns-oarc.net",
	}

	// data from https://github.com/getdnsapi/stubby/blob/develop/stubby.yml.example
	tlsNameMap = map[string]string{
		// ipv4
		"8.8.8.8":         "dns.google",
		"8.8.4.4":         "dns.google",
		"1.1.1.1":         "cloudflare-dns.com",
		"1.0.0.1":         "cloudflare-dns.com",
		"9.9.9.9":         "dns.quad9.net",
		"9.9.9.10":        "dns.quad9.net",
		"145.100.185.15":  "dnsovertls.sinodun.com",
		"145.100.185.16":  "dnsovertls1.sinodun.com",
		"145.100.185.17":  "dnsovertls2.sinodun.com",
		"145.100.185.18":  "dnsovertls3.sinodun.com",
		"185.49.141.37":   "getdnsapi.net",
		"89.233.43.71":    "unicast.censurfridns.dk",
		"158.64.1.29":     "kaitain.restena.lu",
		"176.103.130.130": "dns.adguard.com",
		"176.103.130.131": "dns.adguard.com",
		"176.103.130.132": "dns-family.adguard.com",
		"176.103.130.134": "dns-family.adguard.com",
		"199.58.81.218":   "dns.cmrg.net",
		"51.15.70.167":    "dns.larsdebruin.net",
		"146.185.167.43":  "dot.securedns.eu",
		"81.187.221.24":   "dns-tls.bitwiseshift.net",
		"94.130.110.185":  "ns1.dnsprivacy.at",
		"94.130.110.178":  "ns2.dnsprivacy.at",
		"139.59.51.46":    "dns.bitgeek.in",
		"89.234.186.112":  "dns.neutopia.org",
		"93.177.65.183":   "dot1.applied-privacy.net",
		"184.105.193.78":  "tls-dns-u.odvr.dns-oarc.net",
		// ipv6
		"2001:610:1:40ba:145:100:185:15": "dnsovertls.sinodun.com",
		"2001:610:1:40ba:145:100:185:16": "dnsovertls1.sinodun.com",
		"2a04:b900:0:100::38":            "getdnsapi.net",
		"2620:fe::fe":                    "dns.quad9.net",
		"2620:fe::10":                    "dns.quad9.net",
		"2606:4700:4700::1111":           "cloudflare-dns.com",
		"2606:4700:4700::1001":           "cloudflare-dns.com",
		"2001:4860:4860::8888":           "dns.google",
		"2001:4860:4860::8844":           "dns.google",
		"2a00:5a60::ad1:0ff":             "dns.adguard.com",
		"2a00:5a60::ad2:0ff":             "dns.adguard.com",
		"2a00:5a60::bad1:0ff":            "dns-family.adguard.com",
		"2a00:5a60::bad2:0ff":            "dns-family.adguard.com",
		"2a01:3a0:53:53::0":              "unicast.censurfridns.dk",
		"2001:a18:1::29":                 "kaitain.restena.lu",
		"2001:610:1:40ba:145:100:185:18": "dnsovertls3.sinodun.com",
		"2001:610:1:40ba:145:100:185:17": "dnsovertls2.sinodun.com",
		"2001:470:1c:76d::53":            "dns.cmrg.net",
		"2a03:b0c0:0:1010::e9a:3001":     "dot.securedns.eu",
		"2001:8b0:24:24::24":             "dns-tls.bitwiseshift.net",
		"2a01:4f8:c0c:3c03::2":           "ns1.dnsprivacy.at",
		"2a01:4f8:c0c:3bfc::2":           "ns2.dnsprivacy.at",
		"2001:67c:27e4::35":              "privacydns.go6lab.si",
		"2a00:5884:8209::2":              "dns.neutopia.org",
		"2a03:4000:38:53c::2":            "dot1.applied-privacy.net",
		"2620:ff:c000:0:1::64:25":        "tls-dns-u.odvr.dns-oarc.net",
	}

	chinaDomainList           string
	withAppleDomainList       string
	withGoogleDomainList      string
	withAppleGoogleDomainList string
	bogusIPList               string
)

const (
	chinaDomainListURL   = `https://raw.githubusercontent.com/felixonmars/dnsmasq-china-list/master/accelerated-domains.china.conf`
	appleDomainListURL   = `https://raw.githubusercontent.com/felixonmars/dnsmasq-china-list/master/apple.china.conf`
	googleDomainListURL  = `https://raw.githubusercontent.com/felixonmars/dnsmasq-china-list/master/google.china.conf`
	bogusIPListURL       = `https://raw.githubusercontent.com/felixonmars/dnsmasq-china-list/master/bogus-nxdomain.china.conf`
	chinaDomainListFile  = `chinaDomain.conf`
	appleDomainListFile  = `apple.conf`
	googleDomainListFile = `google.conf`
	bogusIPListFile      = `bogus.conf`
)

func init() {
	// load China domain names
	chinaDomainList = loadDomainList(chinaDomainListFile, chinaDomainListURL)
	// load Apple domain names
	appleDomainList := loadDomainList(appleDomainListFile, appleDomainListURL)
	// load Google domain names
	googleDomainList := loadDomainList(googleDomainListFile, googleDomainListURL)

	withAppleGoogleDomainList = strings.Join([]string{chinaDomainList, appleDomainList, googleDomainList}, " ")
	withAppleDomainList = strings.Join([]string{chinaDomainList, appleDomainList}, " ")
	withGoogleDomainList = strings.Join([]string{chinaDomainList, googleDomainList}, " ")
	// load bogus IP list
	loadBogusIPList()
}

func loadBogusIPList() {
	c, err := getFileContent(bogusIPListFile, bogusIPListURL)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(c))
	scanner.Split(bufio.ScanLines)
	r := regexp.MustCompile(`^bogus-nxdomain=(.+)$`)
	var ips []string
	for scanner.Scan() {
		ss := r.FindStringSubmatch(scanner.Text())
		if len(ss) > 1 {
			ips = append(ips, ss[1])
		}
	}
	bogusIPList = strings.Join(ips, " ")
}

func loadDomainList(filename string, u string) string {
	c, err := getFileContent(filename, u)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(c))
	scanner.Split(bufio.ScanLines)
	r := regexp.MustCompile(`^server=\/([^\/]+)\/114\.114\.114\.114$`)
	var domains []string
	for scanner.Scan() {
		ss := r.FindStringSubmatch(scanner.Text())
		if len(ss) > 1 {
			domains = append(domains, ss[1])
		}
	}
	return strings.Join(domains, " ")
}

func getFileContent(filename string, u string) ([]byte, error) {
	var c []byte
	if b, e := fsutil.FileExists(filename); e != nil || !b {
		// download it
		client := &http.Client{}
		req, err := http.NewRequest("GET", u, nil)
		if err != nil {
			return nil, err
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return nil, err
		}

		c, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		// save to local file
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		f.Write(c)
	} else {
		f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		c, err = ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func getChinaDomainList(includeApple bool, includeGoogle bool) string {
	if includeApple && includeGoogle {
		return withAppleGoogleDomainList
	}
	if includeApple {
		return withAppleDomainList
	}
	if includeGoogle {
		return withGoogleDomainList
	}
	return chinaDomainList
}

// code is copy from https://github.com/xjdrew/kone, Thanks xjdrew
package geoip

import (
	"net"
	"sort"
	"strings"
)

var (
	geoIPLen = len(GeoIP)
)

func QueryCountry(ip uint32) string {
	i := sort.Search(geoIPLen, func(i int) bool {
		n := GeoIP[i]
		return n.End >= ip
	})

	var country string
	if i < geoIPLen {
		n := GeoIP[i]
		if n.Start <= ip {
			country = n.Name
		}
	}
	return country
}

func QueryCountryByIP(ip net.IP) string {
	ip = ip.To4()
	if ip == nil {
		return ""
	}

	v := uint32(ip[0]) << 24
	v += uint32(ip[1]) << 16
	v += uint32(ip[2]) << 8
	v += uint32(ip[3])
	return QueryCountry(v)
}

func QueryCountryByString(v string) string {
	ip := net.ParseIP(v)
	if ip == nil {
		return ""
	}
	return QueryCountryByIP(ip)
}

func IsDirectDomain(v string) bool {
	index := sort.SearchStrings(Domains, v)
	return Domains[index] == v
}

func IsDirectDomainRecursive(v string) bool {
	for {
		indexd := sort.SearchStrings(Domains, v)
		if Domains[indexd] == v {
			return true
		}
		index := strings.Index(v, ".")
		if index == -1 {
			return false
		}
		v = v[index+1:]

		index = strings.Index(v, ".")
		if index == -1 {
			return false
		}
	}

}

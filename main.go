package main

import (
	"flag"
	"fmt"
	"net"
)

var SiteUrl string

func main() {
	flag.StringVar(&SiteUrl, "d", "", "Set domain name")
	flag.Parse()
	fmt.Println("Domain name: ", SiteUrl)

	dnslist := GetDNS(SiteUrl)
	if dnslist == nil {
		fmt.Println("Error!!!")
	} else {
		fmt.Println("<================================DNS Servers================================>")
		for i := 0; i < len(dnslist); i++ {
			fmt.Println(dnslist[i])
		}
	}

	iplist := GetIP(SiteUrl)
	if iplist == nil {
		fmt.Println("Error!!!")
	} else {
		fmt.Println("<================================IP=========================================>")
		for i := 0; i < len(iplist); i++ {
			fmt.Println(iplist[i])
		}
	}

	txtList := GetTxtList(SiteUrl)
	if txtList == nil {
		fmt.Println("Error!!!")
	} else {
		fmt.Println("<================================TXT========================================>")
		for i := 0; i < len(txtList); i++ {
			fmt.Println(txtList[i])
		}
	}

	fmt.Println("<================================CName======================================>")
	fmt.Println(GetCName(SiteUrl))

}

func GetDNS(siteUrl string) []string {
	var servers []string

	nameserver, _ := net.LookupNS(siteUrl)
	for _, ns := range nameserver {
		servers = append(servers, ns.Host)
	}
	return servers
}

func GetIP(siteUrl string) []string {
	var ips []string

	addr, _ := net.LookupIP(siteUrl)
	for _, ip := range addr {
		ips = append(ips, ip.String())
	}
	return ips
}

func GetCName(siteUrl string) string {
	cname, _ := net.LookupCNAME(siteUrl)
	return cname
}

func GetTxtList(siteUrl string) []string {
	var txtlist []string

	txt, _ := net.LookupTXT(siteUrl)
	for _, txt := range txt {
		txtlist = append(txtlist, txt)
	}
	return txtlist
}

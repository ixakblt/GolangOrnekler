/*
------------------------------------
ixakblt - ibrahim AKBULUT
------------------------------------
Web Site :ixakblt
------------------------------------
All Sites : @ixakblt
------------------------------------
*/

package main

import (
	"fmt"
	"net"
)

//main fonk
func main() {
	hatz := Elit{}
	hatz.link = "fb.com"
	hatz.getCname()
	hatz.getIP()
	hatz.getNs()
	hatz.getHost()
	hatz.getTxt()
	for _, i := range hatz.ns {
		fmt.Printf("%v\n", i[:len(i)-1])
	}

}

type Elit struct {
	link    string
	cname   string
	ip      []string
	ns      []string
	hostlar []string
	txts    []string
}

//getTxt :D
func (e *Elit) getTxt() {
	txtler, _ := net.LookupTXT(e.link)
	for _, i := range txtler {
		e.txts = append(e.txts, i)
	}
}

//getHost test
func (e *Elit) getHost() {
	hostlist, _ := net.LookupHost(e.link)
	for _, i := range hostlist {
		e.hostlar = append(e.hostlar, i)
	}
}

//GetCnmae :D
func (e *Elit) getCname() {
	e.cname, _ = net.LookupCNAME(e.link)
}

func (e *Elit) getNs() {
	nsler, _ := net.LookupNS(e.link)
	for _, i := range nsler {
		e.ns = append(e.ns, i.Host)

	}
}

//getIP fonl
func (e *Elit) getIP() {
	e.ip, _ = net.LookupHost(e.link)

}

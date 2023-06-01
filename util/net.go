package util

import (
	"errors"
	"net"
	"regexp"

	externalip "github.com/glendc/go-external-ip"
)

// GetInterface returns a `name` string of the suitable interface found
func GetInterface(listAll bool) (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	var re = regexp.MustCompile(`^(veth|br\-|docker|lo|EHC|XHC|bridge|gif|stf|p2p|awdl|utun|tun|tap)`)
	for _, iface := range ifaces {
		if !listAll && re.MatchString(iface.Name) {
			continue
		}
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		ip, err := FindIP(iface)
		if err != nil {
			continue
		}
		if ip == "127.0.0.1" {
			continue
		}
		return iface.Name, nil
	}
	return "", errors.New("no interfaces found")
}

// GetExernalIP of this host
func GetExernalIP() (net.IP, error) {
	consensus := externalip.DefaultConsensus(nil, nil)
	// Get your IP, which is never <nil> when err is <nil>
	ip, err := consensus.ExternalIP()
	if err != nil {
		return nil, err
	}
	return ip, nil
}

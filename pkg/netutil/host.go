package netutil

import (
	. "main/pkg/logging"
	"net"
	"strings"
)

const (
	IP_FLAGS = net.FlagUp | net.FlagBroadcast | net.FlagMulticast
)

// Find external IP address
func Ip() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		Logger.Fatal(err)
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			Logger.Fatal(err)
		}

		if iface.Flags&IP_FLAGS == IP_FLAGS && len(addrs) > 0 {
			return strings.Split(addrs[0].String(), "/")[0]
		}
	}

	Logger.Fatal("could not find ip")
	return ""
}

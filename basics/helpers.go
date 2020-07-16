package main

import (
	"net"
	"strings"
)

func GetHostIp() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	ips := make([]string, 0)
	for _, i := range ifaces {

		if !strings.Contains(i.Name, "docker0") && !strings.Contains(i.Name, "lo") &&
			!strings.Contains(i.Name, "vboxnet") && !strings.Contains(i.Name, "bridge") {
			{
				addrs, err := i.Addrs()
				if err != nil {
					return nil, err
				}
				// handle err
				for _, addr := range addrs {
					var ip string
					switch v := addr.(type) {
					case *net.IPNet:
						ip = v.IP.String()
						ips = append(ips, ip)
					case *net.IPAddr:
						ip = v.IP.String()
						ips = append(ips, ip)
					}
				}
			}

		}
	}
	return ips, nil
}

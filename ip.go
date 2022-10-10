package ipv4

import "net"

// LocalIP返回所有非环回IPv4地址
func LocalIPv4s() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	return ips, nil
}

//返回ipv4地址
func LocalIP() string {
	ips, _ := LocalIPv4s()
	return ips[0]
}

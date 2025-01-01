package randomNetIP

import (
	"math/rand"
	"net"
)

func RandomNetIPv4() string {
	privateRanges := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"169.254.0.0/16",
		"100.64.0.0/10",
		"127.0.0.0/8",
	}

	for {
		firstOctet := rand.Intn(223-1) + 1
		secondOctet := rand.Intn(256)
		thirdOctet := rand.Intn(256)
		fourthOctet := rand.Intn(256)

		ip := net.IPv4(byte(firstOctet), byte(secondOctet), byte(thirdOctet), byte(fourthOctet)).String()

		isPrivate := false
		for _, cidr := range privateRanges {
			_, ipNet, _ := net.ParseCIDR(cidr)
			if ipNet.Contains(net.ParseIP(ip)) {
				isPrivate = true
				break
			}
		}
		if !isPrivate {
			return ip
		}
	}
}

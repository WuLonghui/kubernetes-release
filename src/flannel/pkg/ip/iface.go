package ip

import (
	"errors"
	"net"

	"github.com/coreos/flannel/Godeps/_workspace/src/github.com/docker/libcontainer/netlink"
)

func GetIfaceIP4Addr(iface *net.Interface) (net.IP, error) {
	addrs, err := iface.Addrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		// Attempt to parse the address in CIDR notation
		// and assert it is IPv4
		ip, _, err := net.ParseCIDR(addr.String())
		if err == nil && ip.To4() != nil {
			return ip.To4(), nil
		}
	}

	return nil, errors.New("No IPv4 address found for given interface")
}

func GetIfaceIP4AddrMatch(iface *net.Interface, matchAddr net.IP) error {
	addrs, err := iface.Addrs()
	if err != nil {
		return err
	}

	for _, addr := range addrs {
		// Attempt to parse the address in CIDR notation
		// and assert it is IPv4
		ip, _, err := net.ParseCIDR(addr.String())
		if err == nil && ip.To4() != nil {
			if ip.To4().Equal(matchAddr) {
				return nil
			}
		}
	}

	return errors.New("No IPv4 address found for given interface")
}

func GetDefaultGatewayIface() (*net.Interface, error) {
	routes, err := netlink.NetworkGetRoutes()
	if err != nil {
		return nil, err
	}

	for _, route := range routes {
		if route.Default {
			if route.Iface == nil {
				return nil, errors.New("Found default route but could not determine interface")
			}
			return route.Iface, nil
		}
	}

	return nil, errors.New("Unable to find default route")
}

func GetInterfaceByIP(ip net.IP) (*net.Interface, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range ifaces {
		err := GetIfaceIP4AddrMatch(&iface, ip)
		if err == nil {
			return &iface, nil
		}
	}

	return nil, errors.New("No interface with given IP found")
}

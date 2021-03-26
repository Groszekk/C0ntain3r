package models

import "github.com/vishvananda/netlink"

type VethPair struct {
	Veth netlink.Link
	Name string
	Peer netlink.Link
	IP string
	PeerIP string
	PeerName string
}
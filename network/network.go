package network

import (
	"Container/models"
	"github.com/vishvananda/netlink"
	"Container/iptables"
)

// todo: error handling
func SetNetworkInterfaces(pid int, config models.Config) error {
	rules := []models.Rules{
		{Table:"nat", Chain:"POSTROUTING", Rp:[]string{"--source", config.BridgeIP, "--out-interface", config.HostInterface, "--jump", "MASQUERADE"}},
		{Table:"filter", Chain:"FORWARD", Rp:[]string{"--in-interface", config.HostInterface, "--out-interface", config.BridgeName, "--jump", "ACCEPT"}},
		{Table:"filter", Chain:"FORWARD", Rp:[]string{"--out-interface", config.HostInterface, "--in-interface", config.BridgeName, "--jump", "ACCEPT"}},

	}

	iptables.SetITables(rules)

	bridge := createBridge(config)
	veth := createVirtualEthDevice(123, config)
	netlink.LinkSetMaster(veth, bridge.(*netlink.Bridge)) // sets the master of the link device

	return nil
}

func createBridge(config models.Config) netlink.Link {
	nLink, _ := netlink.LinkByName(config.BridgeName)
	link := netlink.NewLinkAttrs()

	link.Name = config.BridgeName
	nLink = &netlink.Bridge{LinkAttrs: link}

	netlink.LinkAdd(nLink)
	addr, _ := netlink.ParseAddr(config.BridgeIP)
	netlink.AddrAdd(nLink, addr)
	netlink.LinkSetUp(nLink)

	return nLink
}

func createVirtualEthDevice(pid int, config models.Config) netlink.Link {
	nLink, _ := netlink.LinkByName(config.BridgeName)
	link := netlink.NewLinkAttrs()
	link.Name = config.VethPair.Name
	link.MasterIndex = nLink.Attrs().Index

	vethPair := &netlink.Veth{LinkAttrs: link, PeerName: config.VethPair.PeerName}
	netlink.LinkDel(vethPair)

	peer, _ := netlink.LinkByName(config.VethPair.PeerName)
	netlink.LinkSetNsPid(peer, pid)

	addr, _ := netlink.ParseAddr(config.VethPair.PeerIP)

	netlink.AddrAdd(vethPair, addr)
	netlink.LinkSetUp(vethPair)

	return nLink
}
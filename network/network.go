package network

import (
	"Container/models"
	// "github.com/vishvananda/netlink"
	"Container/iptables"
)

func SetNetwork(pid int, config models.Config) error {
	rules := []models.Rules{
		{Table:"nat", Chain:"POSTROUTING", Rp:[]string{"--source", config.BridgeIP, "--out-interface", config.HostInterface, "--jump", "MASQUERADE"}},
		{Table:"filter", Chain:"FORWARD", Rp:[]string{"--in-interface", config.HostInterface, "--out-interface", config.BridgeName, "--jump", "ACCEPT"}},
		{Table:"filter", Chain:"FORWARD", Rp:[]string{"--out-interface", config.HostInterface, "--in-interface", config.BridgeName, "--jump", "ACCEPT"}},

	}

	iptables.SetITables(rules)

	return nil
}

// in progress...
// func createBridge(nLink *netlink.Link, bridgeName string) error {
// 	bridge, err := netlink.LinkByName(bridgeName)
// 	if err != nil {
// 		return err
// 	}
// 	l := netlink.NewLinkAttrs()
// }
package iptables

import (
	"Container/models"
	"github.com/coreos/go-iptables/iptables"
)

func SetITables(itablesRules []models.Rules) error {
	itables, err := iptables.New()
	if err != nil {
		return err
	}

	for _, r := range itablesRules {
		if err := itables.AppendUnique(r.Table, r.Chain, r.Rp...); err != nil {
			return err
		}
	}

	return nil
}
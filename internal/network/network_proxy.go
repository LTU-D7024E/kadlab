package network

import (
	"context"
	"d7024e_group04/env"
)

type ProxyNetwork struct {
	publicNetwork *PublicNetwork
}

func NewProxyNetwork() *ProxyNetwork {
	return &ProxyNetwork{
		publicNetwork: &PublicNetwork{},
	}
}

func (network *ProxyNetwork) ResolveDNS(ctx context.Context, domain string) ([]string, error) {
	switch domain {
	case env.NodesProxyDomain:
		return []string{"127.0.0.1"}, nil
	default:
		return network.publicNetwork.ResolveDNS(ctx, domain)
	}
}

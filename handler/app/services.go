package app

import "cacher/pkg/hashing"

func (a *application) InitCacheService() *hashing.HashRing {
	hr := hashing.NewHashRing(a.config.System.VNodeCount)
	nodes := a.config.Nodes
	for _, node := range nodes {
		hr.AddNode(node.Name, node.Address)
	}
	return hr
}

package hashing

import (
	"cacher/pkg/cache/rds"
	"fmt"
	"hash/crc32"
	"sort"
	"sync"
)

type Node struct {
	Title   string
	Address string
	Hash    uint32
}

type HashRing struct {
	sync.RWMutex
	Nodes        []*Node
	Clients      map[string]rds.CacheClient
	vnodePerNode int
}

func NewHashRing(vnode int) *HashRing {
	hr := &HashRing{
		Nodes:        make([]*Node, 0),
		Clients:      make(map[string]rds.CacheClient),
		vnodePerNode: vnode,
	}
	return hr
}

func (hr *HashRing) Len() int {
	return len(hr.Nodes)
}

func (hr *HashRing) AddNode(title string, address string) {
	hr.Lock()
	defer hr.Unlock()

	if _, ok := hr.Clients[address]; ok {
		return // node already exists
	}
	for i := range hr.vnodePerNode {
		t := fmt.Sprintf("%s#%d", title, i)
		hash := getHashKey(fmt.Sprintf("%s:%s", t, address))
		node := Node{
			Title:   t,
			Address: address,
			Hash:    hash,
		}
		hr.Nodes = append(hr.Nodes, &node)
		hr.Clients[address] = rds.NewClient(node.Address)
	}

	sort.Slice(hr.Nodes, func(i, j int) bool {
		return hr.Nodes[i].Hash < hr.Nodes[j].Hash
	})
}

func (hr *HashRing) RemoveNode(address string) {
	hr.Lock()
	defer hr.Unlock()

	if _, ok := hr.Clients[address]; !ok {
		return // node does not exist
	}

	delete(hr.Clients, address)
	for i := hr.Len() - 1; i >= 0; i-- {
		if hr.Nodes[i].Address == address {
			hr.Nodes = append(hr.Nodes[:i], hr.Nodes[i+1:]...)
		}
	}
}

func (hr *HashRing) GetNode(key string) *Node {
	hr.RLock()
	defer hr.RUnlock()

	hash := getHashKey(key)
	index := hr.findIndex(hash)
	if index == -1 || index >= len(hr.Nodes) {
		index = 0
	}

	return hr.Nodes[index]
}

func (hr *HashRing) findIndex(hash uint32) int {
	for i, node := range hr.Nodes {
		if hash <= node.Hash {
			return i
		}
	}
	return -1
}

func getHashKey(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

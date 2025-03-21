package cache

import (
	"maps"

	"github.com/go-orb/go-orb/registry"
)

func serviceToRecord(service *registry.Service) record {
	nodes := make(map[string]node, len(service.Nodes))

	for _, n := range service.Nodes {
		mnode := registry.Node{
			ID:        n.ID,
			Address:   n.Address,
			Metadata:  maps.Clone(n.Metadata),
			Transport: n.Transport,
		}

		nodes[n.ID] = node{
			Node: mnode,
		}
	}

	endpoints := make([]registry.Endpoint, len(service.Endpoints))
	for i, e := range service.Endpoints {
		endpoints[i] = registry.Endpoint{
			Name:     e.Name,
			Request:  e.Request,
			Response: e.Response,
			Metadata: maps.Clone(e.Metadata),
		}
	}

	return record{
		Name:      service.Name,
		Version:   service.Version,
		Metadata:  maps.Clone(service.Metadata),
		Nodes:     nodes,
		Endpoints: endpoints,
	}
}

func recordToService(inRecord record) *registry.Service {
	endpoints := make([]*registry.Endpoint, len(inRecord.Endpoints))

	for i, e := range inRecord.Endpoints {
		request := new(registry.Value)
		if e.Request != nil {
			*request = *e.Request
		}

		response := new(registry.Value)
		if e.Response != nil {
			*response = *e.Response
		}

		endpoints[i] = &registry.Endpoint{
			Name:     e.Name,
			Request:  request,
			Response: response,
			Metadata: maps.Clone(e.Metadata),
		}
	}

	nodes := make([]*registry.Node, len(inRecord.Nodes))
	i := 0

	for _, n := range inRecord.Nodes {
		nodes[i] = &registry.Node{
			ID:        n.ID,
			Address:   n.Address,
			Transport: n.Transport,
			Metadata:  maps.Clone(n.Metadata),
		}
		i++
	}

	return &registry.Service{
		Name:      inRecord.Name,
		Version:   inRecord.Version,
		Metadata:  maps.Clone(inRecord.Metadata),
		Endpoints: endpoints,
		Nodes:     nodes,
	}
}

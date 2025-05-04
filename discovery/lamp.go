package discovery

import (
	"context"

	"github.com/nexitf/lamp"
	"github.com/nexitf/unit/discovery"
)

type LampServiceDiscovery struct {
	lc *lamp.Client
}

// NewLampServiceDiscovery
func NewLampServiceDiscovery(lc *lamp.Client) (sd *LampServiceDiscovery) {
	return &LampServiceDiscovery{lc: lc}
}

// Watch
func (sd *LampServiceDiscovery) Watch(ctx context.Context, serviceName, tag string, update func(endpoints []discovery.Endpoint, closed bool)) (cancel func() error, err error) {
	return sd.lc.WatchWithContext(ctx, serviceName, tag, func(endpoints []lamp.Endpoint, closed bool) {
		var tmpEndpoints []discovery.Endpoint
		for _, endpoint := range endpoints {
			tmpEndpoints = append(tmpEndpoints, discovery.Endpoint{
				ID:     endpoint.ID,
				Addr:   endpoint.Addr,
				Weight: endpoint.Weight,
				Meta:   endpoint.Meta,
			})
		}
		update(tmpEndpoints, closed)
	})
}

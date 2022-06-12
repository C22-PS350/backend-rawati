package server

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/C22-PS350/backend-rawati/internal/server/apiv1"
)

func (srv *Server) setupGCPClients(h *apiv1.Handler) (*apiv1.Handler, error) {
	client, err := pubsub.NewClient(context.Background(), srv.Config.GCPProject)
	if err != nil {
		return nil, err
	}

	gcpClients := &apiv1.GcpClients{
		PubSub: client,
	}

	h.GcpClients = gcpClients
	return h, nil
}

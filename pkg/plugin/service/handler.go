package service

import (
	"context"
	"sync"

	ocontext "github.com/vmware-tanzu/octant/internal/context"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/vmware-tanzu/octant/pkg/action"
	"github.com/vmware-tanzu/octant/pkg/navigation"
	"github.com/vmware-tanzu/octant/pkg/plugin"
	"github.com/vmware-tanzu/octant/pkg/view/component"
)

// Handler is the plugin service helper handler. Functions on this struct are called from Octant.
type Handler struct {
	HandlerFuncs

	mu sync.Mutex

	name         string
	description  string
	capabilities *plugin.Capabilities

	dashboardFactory func(dashboardAPIAddress string) (Dashboard, error)
	dashboardClient  Dashboard
	router           *Router
}

var _ plugin.Service = (*Handler)(nil)

// Validate validates Handler.
func (p *Handler) Validate() error {
	if p.dashboardFactory == nil {
		return errors.New("plugin handler doesn't know how to create a dashboard client")
	}

	return nil
}

// Register registers a plugin with Octant.
func (p *Handler) Register(ctx context.Context, dashboardAPIAddress string) (plugin.Metadata, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	client, err := p.dashboardFactory(dashboardAPIAddress)
	if err != nil {
		return plugin.Metadata{}, errors.Wrap(err, "create api client")
	}

	p.dashboardClient = client

	return plugin.Metadata{
		Name:         p.name,
		Description:  p.description,
		Capabilities: *p.capabilities,
	}, nil
}

// Print prints components for an object.
func (p *Handler) Print(ctx context.Context, object runtime.Object) (plugin.PrintResponse, error) {
	clientID := ocontext.WebsocketClientIDFrom(ctx)

	if p.HandlerFuncs.Print == nil {
		return plugin.PrintResponse{}, nil
	}

	request := &PrintRequest{
		baseRequest:     newBaseRequest(ctx, p.name),
		DashboardClient: p.dashboardClient,
		Object:          object,
		ClientID:        clientID,
	}

	return p.HandlerFuncs.Print(request)
}

// PrintTabs prints one or more tabs for an object.
func (p *Handler) PrintTabs(ctx context.Context, object runtime.Object) ([]plugin.TabResponse, error) {
	if p.HandlerFuncs.PrintTabs == nil {
		return []plugin.TabResponse{}, nil
	}

	request := &PrintRequest{
		baseRequest:     newBaseRequest(ctx, p.name),
		DashboardClient: p.dashboardClient,
		Object:          object,
		ClientID:        ocontext.WebsocketClientIDFrom(ctx),
	}

	var tabResponses []plugin.TabResponse
	for _, handlerFunc := range p.HandlerFuncs.PrintTabs {
		resp, err := handlerFunc(request)
		if err != nil {
			return []plugin.TabResponse{}, err
		}
		tabResponses = append(tabResponses, resp)
	}
	return tabResponses, nil
}

// ObjectStatus creates status for an object.
func (p *Handler) ObjectStatus(ctx context.Context, object runtime.Object) (plugin.ObjectStatusResponse, error) {
	if p.HandlerFuncs.ObjectStatus == nil {
		return plugin.ObjectStatusResponse{}, nil
	}

	request := &PrintRequest{
		baseRequest:     newBaseRequest(ctx, p.name),
		DashboardClient: p.dashboardClient,
		Object:          object,
		ClientID:        ocontext.WebsocketClientIDFrom(ctx),
	}

	return p.HandlerFuncs.ObjectStatus(request)
}

// HandleAction handles actions given a payload.
func (p *Handler) HandleAction(ctx context.Context, actionName string, payload action.Payload) error {
	if p.HandlerFuncs.HandleAction == nil {
		return nil
	}

	request := &ActionRequest{
		baseRequest:     newBaseRequest(ctx, p.name),
		DashboardClient: p.dashboardClient,
		ActionName:      actionName,
		Payload:         payload,
		ClientID:        ocontext.WebsocketClientIDFrom(ctx),
	}

	return p.HandlerFuncs.HandleAction(request)
}

// Navigation creates navigation.
func (p *Handler) Navigation(ctx context.Context) (navigation.Navigation, error) {
	if p.HandlerFuncs.Navigation == nil {
		return navigation.Navigation{}, nil
	}

	request := &NavigationRequest{
		baseRequest:     newBaseRequest(ctx, p.name),
		DashboardClient: p.dashboardClient,
		ClientID:        ocontext.WebsocketClientIDFrom(ctx),
	}

	return p.HandlerFuncs.Navigation(request)
}

// Content creates content for a given content path.
func (p *Handler) Content(ctx context.Context, contentPath string) (component.ContentResponse, error) {
	handlerFunc, ok := p.router.Match(contentPath)
	if !ok {
		return component.ContentResponse{}, nil
	}

	request := &request{
		baseRequest:     newBaseRequest(ctx, p.name),
		dashboardClient: p.dashboardClient,
		path:            contentPath,
	}

	return handlerFunc(request)
}

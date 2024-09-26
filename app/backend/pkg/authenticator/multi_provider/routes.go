package multiprovider

import (
	"context"
	"fmt"
	"net/http"

	"trigger.com/trigger/internal/google"
)

var providersMap = map[string]Service{
	"google": google.Model{},
}

func GothRouter(ctx context.Context) (*http.ServeMux, error) {
	router := http.NewServeMux()
	provider, ok := ctx.Value(providerKey).(string)
	if !ok || provider == "" {
		return nil, fmt.Errorf("provider not found")
	}

	service, exists := providersMap[provider]
	if !exists {
		return nil, fmt.Errorf("service for provider %s not found", provider)
	}

	handler := Handler{
		Service: service,
	}

	router.Handle(fmt.Sprintf("/auth/%s/auth", provider), http.HandlerFunc(handler.Auth))
	router.Handle(fmt.Sprintf("/auth/%s/callback", provider), http.HandlerFunc(handler.Callback))
	router.Handle(fmt.Sprintf("/auth/%s/logout", provider), http.HandlerFunc(handler.Logout))

	return router, nil
}

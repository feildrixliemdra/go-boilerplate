package appcontext

import "go-boilerplate/internal/config"

// AppContext the app context struct
type AppContext struct {
	config config.Config
}

// NewAppContext initiate appcontext object
func NewAppContext(config *config.Config) *AppContext {
	return &AppContext{
		config: *config,
	}
}

package bootstrap

import (
	"github.com/tianrosandhy/goconfigloader"
)

type Application struct {
	Config *goconfigloader.Config
}

func NewApplication() *Application {
	cfg := goconfigloader.NewConfigLoader()
	application := Application{
		Config: cfg,
	}

	return &application
}

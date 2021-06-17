package mini

import (
	"mini/config"
	"mini/provider"
	"mini/req"
)

type Application struct {
	client req.HttpClient
	conf   *config.Config
}

func (a *Application) Auth() provider.Auth {
	return provider.NewAuth(a.conf, a.client)
}

func (a *Application) Analysis() provider.Analysis {
	return provider.NewAnalysis(a.conf, a.client)
}

// any other providers...

func New(c *config.Config) *Application {
	return &Application{
		client: req.NewRequest(),
		conf:   c,
	}
}

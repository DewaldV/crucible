package crucible

import (
	"code.google.com/p/gorest"
	"fmt"
	"github.com/DewaldV/crucible/config"
	"github.com/DewaldV/crucible/filters"
	"github.com/DewaldV/crucible/filters/cors"
	"net/http"
)

func Start(config config.CrucibleConfiguration, services ...interface{}) {
	registerGoRestServices(services)
	registerDefaultServices()

	http.Handle(config.RootContext, filters.Handle())
	http.ListenAndServe(fmt.Sprintf(":%d", config.HttpPort), nil)
}

func registerDefaultServices() {
	filters.AddFilter(cors.CorsFilter())
	filters.AddHandler(gorest.Handle())
}

func registerGoRestServices(services ...interface{}) {
	for _, serv := range services {
		gorest.RegisterService(serv)
	}
}

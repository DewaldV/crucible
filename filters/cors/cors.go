package cors

import (
	"net/http"
	"strings"
	"github.com/DewaldV/crucible/config"
	"github.com/DewaldV/crucible/filters"
)

func CorsFilter() *corsFilter {
	return corsFilterHandler
}

var corsFilterHandler *corsFilter = &corsFilter{}

type corsFilter struct {
}

func (f *corsFilter) DoFilter(writer http.ResponseWriter, request *http.Request, filterChain chan filters.Filter) {
	sc := config.Conf.Services[request.RequestURI]

	if request.Method == "OPTIONS" {
		f.addOptionsCors(writer, request, sc)
	} else {
		f.addAllowOriginsCors(writer, request, sc.AllowedOrigins)
	}
}

func (f *corsFilter) addOptionsCors(writer http.ResponseWriter, request *http.Request, sc *config.ServiceConfiguration) error {
	f.addAllowOriginsCors(writer, request, sc.AllowedOrigins)
	f.addAllowMethodCors(writer, request, sc.AllowedMethods)
	f.addAllowHeaderCors(writer)
	return nil
}

func (f *corsFilter) addAllowOriginsCors(writer http.ResponseWriter, request *http.Request, allowedOrigins map[string]bool) error {
	origin := strings.ToLower(request.Header.Get("origin"))
	originAllowed := allowedOrigins[origin]

	if originAllowed {
		writer.Header().Add("Access-Control-Allow-Origin", origin)
	}

	return nil
}

func (f *corsFilter) addAllowMethodCors(writer http.ResponseWriter, request *http.Request, allowedMethods map[string]bool) error {
	method := strings.ToUpper(request.Header.Get("access-control-allow-method"))
	methodAllowed := allowedMethods[method]

	if methodAllowed {
		writer.Header().Add("Access-Control-Allow-Method", method)
	}
	return nil
}

func (f *corsFilter) addAllowHeaderCors(writer http.ResponseWriter) {
	writer.Header().Add("Access-Control-Allow-Headers", "accept, origin, x-requested-with, content-type")
}

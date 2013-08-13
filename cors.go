package crucible

import (
	"code.google.com/p/gorest"
	"fmt"
	"net/http"
)

func AddOptionsCors(rb *gorest.ResponseBuilder, request *http.Request, sc *ServiceConfigStruct) error {
	AddAllowOriginsCors(rb, request, sc.AllowedOrigins)
	AddAllowMethodCors(rb, request, sc.AllowedMethods)
	AddAllowHeaderCors(rb)
	return nil
}

func AddAllowOriginsCors(rb *gorest.ResponseBuilder, request *http.Request, allowedOrigins map[string]bool) error {
	origin := request.Header.Get("origin")
	originAllowed := allowedOrigins[origin]

	if originAllowed {
		rb.AddHeader("Access-Control-Allow-Origin", origin)
	}

	return nil
}

func AddAllowMethodCors(rb *gorest.ResponseBuilder, request *http.Request, allowedMethods map[string]bool) error {
	//method := request.Header.Get("access-control-allow-method")
	var allowedMethodsStr string
	for key := range allowedMethods {
		allowedMethodsStr += fmt.Sprintf("%s,", key)
	}
	allowedMethodsStr = allowedMethodsStr[:len(allowedMethodsStr)-1]

	rb.AddHeader("Access-Control-Allow-Method", allowedMethodsStr)
	return nil
}

func AddAllowHeaderCors(rb *gorest.ResponseBuilder) {
	rb.AddHeader("Access-Control-Allow-Headers", "accept, origin, x-requested-with, content-type")
}

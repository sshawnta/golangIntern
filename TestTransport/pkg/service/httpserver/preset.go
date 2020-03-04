//Package service http client
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpserver


const (
	URIPrefix = "/api/v1"

	URIPathClientGetClientId = URIPrefix + "/user"
	URIPathClientPostOrder = URIPrefix + "/orders"
	URIPathClientGetCount = URIPrefix + "/user/:id/count"
	URIPathClientGetOrder = URIPrefix + "/orders"

	URIPathClientGetCountClient = URIPrefix + "/user/%s/count"

	URIPathGetClientId = URIPrefix + "/some/:name"
	URIPathPostOrder = URIPrefix + "/some/:name"
	URIPathGetCount = URIPrefix + "/some/:name"
	URIPathGetOrder = URIPrefix + "/some/:name"

	HTTPMethodGetClientId ="GET"
	HTTPMethodPostOrder = "POST"
	HTTPMethodGetCount = "GET"
	HTTPMethodGetOrder = "GET"
)

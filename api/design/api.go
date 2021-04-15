package design

import (
	"net/http"

	. "goa.design/goa/v3/dsl"
)

// API describes the global properties of the API server.
var _ = API("walletnode", func() {
	Title("WalletNode REST API")
	Version("1.0")

	Server("walletnode", func() {
		Services("artworks", "swagger")

		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})

var InnerError = func(code int) {
	Attribute("code", Int, func() {
		Description("Code refers to a code number in the response header that indicates the general classification of the response.")
		Example(code)
		Default(code)
	})
	Attribute("message", String, func() {
		Description("Message is a human-readable explanation specific to this occurrence of the problem.")
		Example(http.StatusText(code))
		Default(http.StatusText(code))
	})
	Required("code")
}

var BadRequest = Type("BadRequest", func() {
	Attribute("error", func() {
		InnerError(http.StatusBadRequest)
		Meta("struct:field:name", "InnerError")
	})
	Required("error")

})

var InternalServerError = Type("InternalServerError", func() {
	Attribute("error", func() {
		InnerError(http.StatusInternalServerError)
		Meta("struct:field:name", "InnerError")
	})
	Required("error")

})

package http_transport

import "net/http"

type DecoderFunc func(r *http.Request) (interface{}, error)

type EncoderFunc func(d interface{}) (interface{}, error)

type ResponderFunc func(w http.ResponseWriter, d interface{})

type Endpoint interface {
	Handle(d interface{}) (interface{}, error)
}

// This code is auto generated. DO NOT EDIT.

package v0

import (
	"net/http"

	"github.com/metalmatze/signal/server/signalhttp"
	"github.com/prometheus/client_golang/prometheus"
)

type InstrumentedServerInterface struct {
	ServerInterface
	signalhttp.HandlerInstrumenter
}

func NewInstrumentedServerInterface(impl ServerInterface, r prometheus.Registerer) *InstrumentedServerInterface {
	i := signalhttp.NewHandlerInstrumenter(r, []string{"handler"})
	return &InstrumentedServerInterface{impl, i}
}

func (i *InstrumentedServerInterface) GetUser(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		i.ServerInterface.GetUser(w, r)
	}
	i.NewHandler(prometheus.Labels{"handler": "GetUser"}, http.HandlerFunc(handler))(w, r)
}

func (i *InstrumentedServerInterface) UpdatePin(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		i.ServerInterface.UpdatePin(w, r)
	}
	i.NewHandler(prometheus.Labels{"handler": "UpdatePin"}, http.HandlerFunc(handler))(w, r)
}

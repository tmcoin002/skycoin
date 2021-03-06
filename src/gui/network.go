// Network-related information for the GUI
package gui

import (
	"net/http"

	"github.com/tmcoin002/tmcoin/src/daemon"
)

func connectionHandler(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if addr := r.FormValue("addr"); addr == "" {
			Error404(w)
		} else {
			SendOr404(w, gateway.GetConnection(addr))
		}
	}
}

func connectionsHandler(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		SendOr404(w, gateway.GetConnections())
	}
}

func defaultConnectionsHandler(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		SendOr404(w, gateway.GetDefaultConnections())
	}
}

func RegisterNetworkHandlers(mux *http.ServeMux, gateway *daemon.Gateway) {
	mux.HandleFunc("/network/connection", connectionHandler(gateway))
	mux.HandleFunc("/network/connections", connectionsHandler(gateway))
	mux.HandleFunc("/network/defaultConnections", defaultConnectionsHandler(gateway))
}

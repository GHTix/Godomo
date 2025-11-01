package probeService

import "net/http"

type ProbeService struct {
	Ready bool
}

func New() *ProbeService {
	return &ProbeService{
		Ready: true,
	}
}

func (probe *ProbeService) Unready() {
	probe.Ready = false
}

func (probe *ProbeService) ProbeReadyHandler(w http.ResponseWriter, r *http.Request) {
	if probe.Ready {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}

func (probe *ProbeService) ProbeAliveHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

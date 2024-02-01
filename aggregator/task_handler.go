package aggregator

import (
	"net/http"
)

type TaskHandler struct {
	aggregator *Aggregator
}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cid := r.URL.Query().Get("cid")

	s1Client := &System1Client{
		aggregator: h.aggregator,
	}
	s1 := s1Client.getContractHash(cid)

	s2Client := &System2Client{
		aggregator: h.aggregator,
	}
	s2 := s2Client.getContractHash(cid)

	dkgClient := &DkgClient{
		aggregator: h.aggregator,
	}
	dkgHash := dkgClient.getContractHash(cid)

	task := MfssiaTask{
		system1Value: s1,
		system2Value: s2,
		dkgValue:     dkgHash,
	}

	h.aggregator.sendNewTask(task)
}

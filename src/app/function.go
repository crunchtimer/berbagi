package main

import "net/http"

func TallyCount(w http.ResponseWriter, r *http.Request) {
	ctrls, err := NewControllers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = ctrls.App.TallyCount()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

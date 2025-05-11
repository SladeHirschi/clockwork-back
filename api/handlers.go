package api

import (
	"net/http"
)

func (c *ClockworkApi) ClockIn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Clocked in"))
}

func (c *ClockworkApi) Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged in"))
}

func (c *ClockworkApi) Signup(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Signed up"))
}

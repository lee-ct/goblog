package api

import "net/http"

func (*ApiHandle) SaveAndUpdate(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("ok"))
}

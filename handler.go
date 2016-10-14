package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync/atomic"
	"time"
)

const ServerHeader = "ident v" + Version
const ContentTypeHeader = "application/json; charset=utf-8"

type Handler struct {
	Id  uint16
	seq int32
}

func (h *Handler) doResetSeq() {
	ch := time.Tick(1 * time.Second)

	for _ = range ch {
		atomic.StoreInt32(&h.seq, 0)
	}
}

func (h *Handler) generateToken() string {
	v := atomic.AddInt32(&h.seq, 1)
	return Encode(h.Id, uint32(time.Now().Unix()), uint16(v), rand.Uint32())
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ContentTypeHeader)
	w.Header().Set("Server", ServerHeader)

	fmt.Fprintf(w, `{"identifier":"%s"}`, h.generateToken())
}

func NewHandler(id uint16) *Handler {
	h := &Handler{id, 0}
	go h.doResetSeq()
	return h
}

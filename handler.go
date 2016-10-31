package main

import (
	"github.com/reflect/ident/idgen"

	"fmt"
	"net/http"
)

type Handler struct {
	Provider *idgen.Provider
}

const ServerHeader = "ident v" + Version
const ContentTypeHeader = "application/json; charset=utf-8"

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ContentTypeHeader)
	w.Header().Set("Server", ServerHeader)

	fmt.Fprintf(w, `{"identifier":"%s"}`, h.Provider.Next())
}

func (h *Handler) Close() error {
	return h.Provider.Close()
}

func NewHandler(provider *idgen.Provider) *Handler {
	return &Handler{provider}
}

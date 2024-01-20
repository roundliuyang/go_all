package api

import (
	"fmt"
	"net/http"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/julienschmidt/httprouter"
)

func NewHttpServer() *khttp.Server {
	r := httprouter.New()
	r.HandlerFunc(http.MethodGet, "/test", hello)

	srv := khttp.NewServer(
		khttp.Address(":8888"),
	)
	srv.HandlePrefix("/", r)
	return srv
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

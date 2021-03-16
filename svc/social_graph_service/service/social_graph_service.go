package service

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	sgm "github.com/willing8310/delinkcious/pkg/social_graph_manager"
)

func Run() {
	store, err := sgm.NewDbSocialGraphStore("localhost", 5432, "postgres", "postgres")
	if err != nil {
		log.Fatal(err)
	}
	svc, err := sgm.NewSocialGraphManager(store)
	if err != nil {
		log.Fatal((err))
	}

	followHandler := httptransport.NewServer(
		makeFollowEndpoint(svc),
		decodeFollowRequest,
		encodeResponse,
	)

	r := mux.NewRouter()
	r.Methods("POST").Path("/follow").Handler(followHandler)

	log.Println("Listening on port 9090...")
	log.Fatal(http.ListenAndServe(":9090", r))
}

package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"

	"github.com/KedroBoss/gowebdev_TIL/grpc/pb"
)

var gcdClient pb.GCDServiceClient

func main() {

	conn, err := grpc.Dial("gcd-service:3000", grpc.WithInsecure())
	if err != nil {
		log.Println("Error while Dial")
		log.Fatal(err)
	}
	gcdClient = pb.NewGCDServiceClient(conn)
	log.Printf("Connection: %v; Client: %v", conn, gcdClient)

	hr := httprouter.New()

	hr.GET("/gcd/:a/:b", computeGET)

	if err := http.ListenAndServe(":3000", hr); err != nil {
		log.Println("Error while ListenAndServe")
		log.Fatal(err)
	}
}

func computeGET(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	a, err := strconv.ParseUint(ps.ByName("a"), 10, 64)
	log.Printf("A: %v", a)
	if err != nil {
		log.Println("Error while parsing A")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := strconv.ParseUint(ps.ByName("b"), 10, 64)
	log.Printf("B: %v", b)
	if err != nil {
		log.Println("Error while parsing B")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req := &pb.GCDRequest{
		A: a,
		B: b,
	}
	log.Printf("Req: %v", req)
	res, err := gcdClient.Compute(context.Background(), req)
	log.Printf("Res: %v", res)
	if err != nil {
		log.Printf("Error: %v\n", err)
		log.Printf("GRPC Code: %v", grpc.Code(err))

		log.Println("Error while computing")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	xb, err := json.Marshal(res)
	if err != nil {
		log.Println("Error while marshaling")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(xb); err != nil {
		log.Println("Error while writing")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// js, err := json.Marshal(a)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

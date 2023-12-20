package main

import {
	"log"
	"nest/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm/dialects/mysql"
    "github.com/atharvaverma12/book/pkg/routes" //using absolute path from infolder files
}

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localost:9010",r))
}
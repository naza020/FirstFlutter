package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const port = ":5500"

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/products/{fetchCountPercentage}", products).Methods("GET")

	fmt.Println("Serving @ http://127.0.0.1" + port)
	log.Fatal(http.ListenAndServe(port, router))
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is root page"))
}

func products(w http.ResponseWriter, r *http.Request) {
	fetchCountPercentage, errInput := strconv.ParseFloat(mux.Vars(r)["fetchCountPercentage"], 64)
	fetchCount := 0

	if errInput != nil {
		fmt.Println(errInput.Error())
	} else {
		//calculate something
		//input >=100 = all product
		fetchCount = int(float64(len(productList)) * fetchCountPercentage / 100)
		if fetchCount > len(productList) {
			fetchCount = len(productList)
		}
	}

	//write to response
	jsonList, err := json.Marshal(productList[0:fetchCount])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("content-type", "application/json")
		w.Write(jsonList)
	}
}

type product struct {
	Name  string
	Price float64
	Count int
}

var productList = []product{
	product{"p1", 25.0, 30},
	product{"p2", 20.0, 10},
	product{"p3", 250.0, 320},
	product{"p4", 256.0, 730},
	product{"p5", 24.0, 340},
	product{"p6", 10.0, 300},
	product{"a1", 25.0, 30},
	product{"a2", 20.0, 10},
	product{"a3", 250.0, 320},
	product{"a4", 256.0, 730},
	product{"a5", 24.0, 340},
	product{"a6", 10.0, 300},
	product{"x1", 25.0, 30},
	product{"x2", 20.0, 10},
	product{"x3", 250.0, 320},
	product{"x4", 256.0, 730},
	product{"x5", 24.0, 340},
	product{"x6", 10.0, 300},
}

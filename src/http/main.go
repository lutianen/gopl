package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var (
	mutex sync.Mutex
)

type dollars float32
type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// for item, price := range db {
	// 	fmt.Fprintf(w, "%s: %s\n", item, price)
	// }

	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}

	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	price, err := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if err != nil {
		fmt.Fprintf(w, "price is wrong: %f\n", price)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	db[item] = dollars(price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		fmt.Fprintf(w, "%q has been exsisted.", item)
		return
	}
	temp := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(temp, 32)
	if err != nil {
		fmt.Fprintf(w, "The input price is wrong: %q", temp)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	db[item] = dollars(price)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		fmt.Fprintf(w, "%q is not existed.\n", item)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	delete(db, item)
}

func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}
	// log.Fatal(http.ListenAndServe("localhost:5882", db))

	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	mux.Handle("/update", http.HandlerFunc(db.update))
	mux.Handle("/create", http.HandlerFunc(db.create))
	mux.Handle("/delete", http.HandlerFunc(db.delete))

	log.Fatal(http.ListenAndServe("localhost:5882", mux))
}

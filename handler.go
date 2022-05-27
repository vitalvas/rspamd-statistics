package main

import (
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
	"github.com/vmihailenco/msgpack/v5"
)

func requestHandler(w http.ResponseWriter, req bunrouter.Request) error {
	var items []uint64

	decoder, err := compressionReader(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))

		log.Println("bad request", err)

		return nil
	}

	if err := msgpack.NewDecoder(decoder).Decode(&items); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))

		log.Println("bad request", err)

		return nil
	}

	log.Println("items:", items)

	w.WriteHeader(http.StatusNoContent)

	return nil
}

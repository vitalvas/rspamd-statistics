package app

import (
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
	"github.com/vmihailenco/msgpack/v5"
)

func (app *App) requestFindHandler(w http.ResponseWriter, req bunrouter.Request) error {
	var requestItems []uint64

	if err := msgpack.NewDecoder(req.Body).Decode(&requestItems); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))

		log.Println("bad request", err)

		return nil
	}

	found, err := app.storage.Find(requestItems)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))

		return err
	}

	w.Header().Set("content-type", "application/msgpack")

	if err := msgpack.NewEncoder(w).Encode(found); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))

		return err
	}

	return nil
}

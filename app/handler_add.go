package app

import (
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
	"github.com/vmihailenco/msgpack/v5"
)

func (app *App) requestAddHandler(w http.ResponseWriter, req bunrouter.Request) error {
	var requestItems []uint64

	if err := msgpack.NewDecoder(req.Body).Decode(&requestItems); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))

		log.Println("bad request", err)

		return nil
	}

	if err := app.storage.Add(requestItems); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))

		return err
	}

	w.WriteHeader(http.StatusNoContent)

	return nil
}

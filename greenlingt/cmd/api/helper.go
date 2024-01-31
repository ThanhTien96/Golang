package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	// convert type to int64
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

// Define a writeJson() helper for sending responses.
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, header http.Header) error {
	// Encode the data to JSON, returning the error if there was one
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// append a newline to make it easer to view in terminal application.
	js = append(js, '\n')
	fmt.Println(string(js))
	/**
	At this point, we know that we won't encounter any more error before writing the response
	**/

	for key, val := range header {
		w.Header()[key] = val
	}

	// add the Content-Type: application haeder, then write the status code and JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

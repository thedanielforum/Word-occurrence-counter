package app

import (
    "net/http"
    "encoding/json"
    "log"
    "strconv"
    "io/ioutil"
)

/*
 * Server config and http helpers
 */
func init() {
    http.HandleFunc("/", countOccurrencesHandler)
}

func responseEncoder(w http.ResponseWriter, status int, data interface{}) {
    w.WriteHeader(status)

    if err := json.NewEncoder(w).Encode(data); err != nil {
        log.Printf("Failed to encode json result %v", err)
    }
}

/*
 * URL/Request Handlers
 */
func countOccurrencesHandler(w http.ResponseWriter, r *http.Request) {
    var err error
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)
    res, err := CountOccurrences(string(body))
    if err != nil {
        responseEncoder(w, http.StatusInternalServerError, err.Error())
        return
    }

    // Limit results
    limitHeader := r.Header.Get("limit")
    limit := 10
    if limitHeader != "" {
        limit, err = strconv.Atoi(limitHeader)
        if err != nil {
            responseEncoder(w, http.StatusBadRequest, err.Error())
            return
        }
    }

    // To prevent "index out of bound"
    if len(res) > limit {
        responseEncoder(w, http.StatusOK, res[0:limit])
        return
    }
    responseEncoder(w, http.StatusOK, res)
}

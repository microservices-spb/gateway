package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	srv Srv
}

func New(srv Srv) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) Do(x, y int64) int64 {
	if x == y {
		return 0
	}
	fmt.Println("api handler: ", x+y)
	return h.srv.Mulity(x, y)
}

type RequestData struct {
	DataQwerty string `json:"data"`
	Number     int    `json:"number"`
}

type ResponseData struct {
	Text  string `json:"text"`
	Count int    `json:"count"`
}

func (h *Handler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//q := r.PathValue("q")
		aString := r.URL.Query().Get("a")
		bString := r.URL.Query().Get("b")

		aInt, err := strconv.ParseInt(aString, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		bInt, err := strconv.ParseInt(bString, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		res := h.srv.Mulity(aInt, bInt)

		log.Println("[GET]")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("result %d", res)))
	case http.MethodPost:
		data, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		defer r.Body.Close()

		var reqData RequestData
		err = json.Unmarshal(data, &reqData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		log.Println(reqData, reqData.DataQwerty, reqData.Number)

		resData := ResponseData{
			Text:  reqData.DataQwerty + " after request",
			Count: reqData.Number * 2,
		}

		resJson, err := json.Marshal(resData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		log.Println("[POST]")
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("x-gateway", "qwerty")
		w.Write(resJson)
	case http.MethodPut:
	case http.MethodDelete:
	}
}

// http://localhost:3111/?a=6&b=2
// http://localhost:3111/1234

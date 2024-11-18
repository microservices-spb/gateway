package api

import (
	"encoding/json"
	"fmt"
	"github.com/microservices-spb/gateway/internal/model"
	"io"
	"log"
	"net/http"
)

type Handler struct {
	srv Srv
	aC  AuthClient
}

func New(srv Srv, aC AuthClient) *Handler {
	return &Handler{
		srv: srv,
		aC:  aC,
	}
}

func (h *Handler) Do(x, y int64) int64 {
	if x == y {
		return 0
	}
	fmt.Println("api handler: ", x+y)
	return h.srv.Mulity(x, y)
}

func (h *Handler) Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("api handler: ", r.URL.Path)
	defer log.Println("finish handle: ", r.URL.Path)
	switch r.Method {
	case http.MethodPost:
		data, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		defer r.Body.Close()

		var reqData model.RequestData
		err = json.Unmarshal(data, &reqData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		token, err := h.aC.DoLogin(r.Context(), reqData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		resData := model.ResponseData{
			Token: token,
		}

		resJson, err := json.Marshal(resData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		log.Println("[POST]")
		w.WriteHeader(http.StatusOK)
		w.Write(resJson)
	}
}

// http://localhost:3111/?a=6&b=2
// http://localhost:3111/1234

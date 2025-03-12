package api

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/microservices-spb/gateway/internal/model"
)

type Handler struct {
	//srv Srv
	aC AuthClient
	Db UserRepository
}

func New(Db UserRepository, aC AuthClient) *Handler {
	return &Handler{
		Db: Db,
		aC: aC,
	}
}

/*func (h *Handler) Do(x, y int64) int64 {
	if x == y {
		return 0
	}
	fmt.Println("api handler: ", x+y)
	return h.srv.Mulity(x, y)
}*/

func (h *Handler) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Println("api handler: ", r.URL.Path)
	defer log.Println("finish handle: ", r.URL.Path)
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var user model.User

	err = json.Unmarshal(data, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	id, err := h.Db.Save(context.Background(), &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	log.Println(id)

	token, err := h.aC.DoLogin(r.Context(), model.RequestData{})
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

// http://localhost:3111/?a=6&b=2
// http://localhost:3111/1234

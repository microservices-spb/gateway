package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/microservices-spb/gateway/internal/model"
)

type Handler struct {
	//srv Srv
	aC AuthClient
	Db *Repository.Conn
}

func New(Db *sqlx.DB, aC AuthClient) *Handler {
	return &Handler{
		Db: *sqlx.DB,
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

	var reqData model.RequestData

	err = json.Unmarshal(data, &reqData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var reqId int32
	query := "INSERT INTO usersInfo (username, password) VALUES ($1, $2) RETURNING id"
	err = h.Db.QueryRow(query, reqData.Username, reqData.Password).Scan(&reqId)
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

// http://localhost:3111/?a=6&b=2
// http://localhost:3111/1234

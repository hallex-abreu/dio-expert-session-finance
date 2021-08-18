package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hallex-abreu/dio-expert-session-finance/model/transaction"
)

func GetTransations(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var layout = "2021-11-20T15:04:05"
	salaryReceveide, _ := time.Parse(layout, "2022-11-20T11:04:05")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1300.0,
			Type:      0,
			CreatedAt: salaryReceveide,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

func CreateTransations(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}

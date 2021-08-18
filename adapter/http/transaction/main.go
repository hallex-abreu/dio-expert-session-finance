package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hallex-abreu/dio-expert-session-finance/model/transaction"
	"github.com/hallex-abreu/dio-expert-session-finance/util"
)

func GetTransations(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1300.0,
			Type:      0,
			CreatedAt: util.StringToTime("2021-02-12T10:10:10"),
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

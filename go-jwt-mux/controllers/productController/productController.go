package productController

import (
	"net/http"

	"github.com/gagassurya19/go-basic-project/go-jwt-mux/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"id": 1,
			"nama_product": "kemeja",
			"stok": 1000,
		},
		{
			"id": 2,
			"nama_product": "celana",
			"stok": 5000,
		},
		{
			"id": 3,
			"nama_product": "kaos",
			"stok": 11000,
		},
	}

	helper.ResponseJSON(w,http.StatusOK, data)
}
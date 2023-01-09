package productcontroller

import (
	"net/http"

	"github.com/pjw1702/go-jwt-mux/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {

	data := []map[string]interface{}{
		{
			"id":           1,
			"name_product": "PjwProduct",
			"stock":        1000,
		},
		{
			"id":           2,
			"name_product": "PjwProduct2",
			"stock":        1000,
		},
		{
			"id":           3,
			"name_product": "PjwProduct3",
			"stock":        500,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)
}

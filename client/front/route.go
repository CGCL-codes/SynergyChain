package front

import (
	"ShareChain-Client/client"
	"ShareChain-Client/vrf"
	"github.com/gorilla/mux"
	"net/http"
)

// 路由结构
type Route struct {
	routeName   string
	Method      string
	HandlerFunc http.HandlerFunc
}
// 路由表结构
type Routes []Route
var ALLRoutes = Routes{
	Route{
		routeName:   "/GetData/{key}",
		Method:      "GET",
		HandlerFunc: HandleGetData,
	},
}

func HandleGetData(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["key"]

	data := client.GetSharedData("http://127.0.0.1:7999/GetSharingData", key, "Address-Client-A")
	data.Value = "#####  经过篡改后的数据  #####"
	isVrf := vrf.VerifyData(*data)

	type response struct {
		PassVrf bool `json:"pass_vrf"`
		Data    interface{} `json:"data"`
	}

	resp := response{
		PassVrf: isVrf,
		Data:    data,
	}

	RespondWithJSON(w, http.StatusOK, resp)
}
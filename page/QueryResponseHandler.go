package page

import (
	"log"
	"net/http"

	"go_mongo/mongo"
)

func QueryResponseHandler(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	andOr := req.FormValue("anyOr")
	queryDevice := mongo.Device{
		Hostname:   req.FormValue("hostname"),
		IPAddress:  req.FormValue("ipAddress"),
		DeviceType: req.FormValue("deviceType"),
	}
	//
	deviceList := mongo.ListDevice(&queryDevice, &andOr, res, req)
	err = tpl.ExecuteTemplate(res, "queryResponse.gohtml", deviceList)
	if err != nil {
		log.Fatalln(err)
	}
}
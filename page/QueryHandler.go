package page

import (
	"log"
	"net/http"

	"go_mongo/mongo"
)

//QueryHandler calls the query.gohtml.  URL is localhost/query.
func QueryHandler(res http.ResponseWriter, req *http.Request) {
	//fmt.Println("DeviceHandler called")
	rh := req.Method
	if rh == "GET" {

		deviceList := mongo.ListDevice("", res, req)
		//fmt.Println(deviceList)
		//fmt.Printf("deviceList is of type %T\n", deviceList)
		err := tpl.ExecuteTemplate(res, "query.gohtml", deviceList)
		if err != nil {
			log.Fatalln(err)
		}

	} else {
		err := req.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}

		queryString := req.FormValue("hostname")

		deviceList := mongo.ListDevice(queryString, res, req)
		err = tpl.ExecuteTemplate(res, "query.gohtml", deviceList)
		if err != nil {
			log.Fatalln(err)
		}
	}

}
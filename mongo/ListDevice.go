package mongo

import (
	"fmt"
	"log"

	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type dInt interface {
	makeMap() map[string]string
}

//MdbDevice struct
type MdbDevice struct {
	Hostname   string `bson:"hostname"`
	IPAddress  string `bson:"ipaddress"`
	DeviceType string `bson:"devicetype"`
}

//ListDevice returns all hostnames from the Device Collection
func ListDevice(qy *MdbDevice, w http.ResponseWriter, r *http.Request) []Device {

	session, err := mgo.Dial("10.132.0.5")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	deviceCollect := session.DB("test").C("device")

	session.SetMode(mgo.Monotonic, true)
	search, _ := bson.Marshal(qy)
	result := []Device{}
	//mongoMap := qy.makeMongoString()
	if qy.Hostname+qy.DeviceType+qy.IPAddress == "" {
		err = deviceCollect.Find(bson.M{}).All(&result)
	} else {
		fmt.Println("This is with the string convert", string(search))
		fmt.Println("This is without the convert", search)
		err = deviceCollect.Find(string(search)).All(&result)
		//fmt.Println(qy)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result

}

/*func (d *Device) makeMongoString() mongoDevice {
	newDev := mongoDevice{
		Hostname:   map[string]string{"hostname": d.Hostname},
		IPAddress:  map[string]string{"ipaddress": d.IPAddress},
		DeviceType: map[string]string{"devicetype": d.DeviceType},
	}

	return newDev
}

type mongoDevice struct {
	Hostname   map[string]string `bson:"hostname"`
	IPAddress  map[string]string `bson:"ipaddress"`
	DeviceType map[string]string `bson:"devicetype"`
}*/

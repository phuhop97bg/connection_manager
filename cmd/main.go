package main

import (
	"hrm-system/internal/model"
	"hrm-system/internal/repository"
	"log"
)

func init() {
	log.Println("Initializing HRM System...")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	cm := repository.NewConnectionManager()
	// connection03 := cm.GetConnectionByPartnerID("03")
	// connection02 := cm.GetConnectionByPartnerID("02")
	// connection04 := cm.GetConnectionByPartnerID("04")
	// log.Println(connection02.GetProfileByID(1))
	// log.Println(connection03.GetProfileByID(1))
	// log.Println(connection04.GetProfileByID(1))
	cm.GetCommonDB().Create(&model.Partner{
		Id:          6,
		PartnerId:   "06",
		PartnerName: "Partner 06",
	})
	cm.OnNewPartner("06")
	log.Println(cm.GetConnectionByPartnerID("06").CreateProfile(&model.Profile{
		ID:        6,
		ProfileId: "06",
		FirstName: "Profile 06",
		LastName:  "Profile 06",
	}))
	log.Println(cm.GetConnectionByPartnerID("06").GetProfileByID(6))
}

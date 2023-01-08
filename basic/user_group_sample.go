package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUserGroup() {
	batman := basic.User{
		Id:       97,
		Username: "Batman",
		IsActive: true,
		Password: []byte("batmanpassword"),
	}

	nightwing := basic.User{
		Id:       96,
		Username: "Nightwing",
		IsActive: true,
		Password: []byte("nightwingpassword"),
	}

	robin := basic.User{
		Id:       95,
		Username: "Robin",
		IsActive: true,
		Password: []byte("robinpassword"),
	}

	batFamily := basic.UserGroup{
		GroupId:   999,
		GroupName: "Bat Family",
		Users:     []*basic.User{&batman, &nightwing, &robin},
	}

	jsonBytes, _ := protojson.Marshal(&batFamily)
	log.Println(string(jsonBytes))
}

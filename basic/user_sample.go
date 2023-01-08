package basic

import (
	"log"
	"math/rand"
	"my-protobuf/protogen/basic"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/anypb"
)

func BasicUser() {
	addr := basic.Addresss{
		Street:  "Daily Planet",
		City:    "Metropolis",
		Country: "US",
		Coordinate: &basic.Addresss_Coordinate{
			Latitude:  40.70797893425118,
			Longitude: -74.01163838107261,
		},
	}

	comm := randomCommunicationChannel()

	u := basic.User{
		Id:                   99,
		Username:             "superman",
		IsActive:             true,
		Password:             []byte("supermanpassword"),
		Gender:               basic.Gender_GENDER_MALE,
		Emails:               []string{"superman@movie.com", "superman@dc.com"},
		Addresss:             &addr,
		CommunicationChannel: &comm,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}

func BasicUnmarshalAnyKnown() {
	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "my-social-media-platform",
		SocialMediaUsername: "my-social-media-username",
	}

	var a anypb.Any
	anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})

	// known type (Social Media)
	sm := new(basic.SocialMedia)

	if err := a.UnmarshalTo(sm); err != nil {
		return
	}

	jsonBytes, _ := protojson.Marshal(sm)
	log.Println(string(jsonBytes))
}

func BasicUnmarshalAnyNotKnown() {
	a := randomCommunicationChannel()

	var unmarshaled protoreflect.ProtoMessage

	unmarshaled, err := a.UnmarshalNew()
	if err != nil {
		return
	}

	log.Println("Unmarshal as", unmarshaled.ProtoReflect().Descriptor().Name())

	jsonBytes, _ := protojson.Marshal(unmarshaled)
	log.Println(string(jsonBytes))
}

func BasicUnmarshalAnyIs() {
	a := randomCommunicationChannel()

	pm := new(basic.PaperMail)

	if a.MessageIs(pm) {
		if err := a.UnmarshalTo(pm); err != nil {
			return
		}

		jsonBytes, _ := protojson.Marshal(pm)
		log.Println(string(jsonBytes))
	} else {
		log.Println("Not PaperMail, do nothing :", a.TypeUrl)
	}
}

func randomCommunicationChannel() anypb.Any {
	rand.Seed(time.Now().UnixNano())

	paperMail := basic.PaperMail{
		PaperMailAddress: "Some paper mail address",
	}

	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "byteMe",
		SocialMediaUsername: "krypton.man",
	}

	instantMessaging := basic.InstantMessaging{
		InstantMessagingProduct:  "whatsup",
		InstantMessagingUsername: "krypton.last",
	}

	var a anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		anypb.MarshalFrom(&a, &paperMail, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})
	default:
		anypb.MarshalFrom(&a, &instantMessaging, proto.MarshalOptions{})
	}

	return a
}

func ProtoToJsonUser() {
	p := basic.User{
		Id:       98,
		Username: "wonderwoman",
		IsActive: true,
		Password: []byte("wonderwomanpassword"),
		Gender:   basic.Gender_GENDER_FEMALE,
		Emails:   []string{"wonderwoman@movie.com", "wonderwoman@dc.com"},
	}

	jsonBytes, _ := protojson.Marshal(&p)
	log.Println(string(jsonBytes))
}

func JsonToProtoUser() {
	json := `
	{
		"id":97,
		"username":"batman",
		"is_active":true,
		"password":"YmF0bWFucGFzc3dvcmQ=",
		"emails":[
		   "batman@movie.com",
		   "batman@dc.com"
		],
		"gender":"GENDER_MALE"
	 }
	`
	var p basic.User

	err := protojson.Unmarshal([]byte(json), &p)

	if err != nil {
		return
	}

	log.Println(&p)
}

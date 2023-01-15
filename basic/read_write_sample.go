package basic

import (
	"io/ioutil"
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/proto"
)

func WriteProtoToFile(msg proto.Message, fname string) {
	b, err := proto.Marshal(msg)

	if err != nil {
		log.Fatalln("Can't marshal message", err)
	}

	if err := ioutil.WriteFile(fname, b, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
	}
}

func ReadProtoFromFile(fname string, dest proto.Message) {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can't read file", fname, err)
	}

	if err := proto.Unmarshal(in, dest); err != nil {
		log.Fatalln("Can't unmarshal", err)
	}
}

func dummyUser() basic.User {
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

	sr := map[string]uint32{"fly": 5, "speed": 5, "durability": 4}

	return basic.User{
		Id:                   99,
		Username:             "superman",
		IsActive:             true,
		Password:             []byte("supermanpassword"),
		Gender:               basic.Gender_GENDER_MALE,
		Emails:               []string{"superman@movie.com", "superman@dc.com"},
		Addresss:             &addr,
		CommunicationChannel: &comm,
		SkillRating:          sr,
	}
}

func WriteToFileSample() {
	u := dummyUser()

	WriteProtoToFile(&u, "superman.bin")
}

func ReadFromFileSample() {
	dest := new(basic.User)

	ReadProtoFromFile("superman.bin", dest)

	log.Println(dest)
}

package basic

import (
	"io/ioutil"
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteProtoToJSON(msg proto.Message, fname string) {
	b, err := protojson.Marshal(msg)

	if err != nil {
		log.Fatalln("Can't marshal message", err)
	}

	if err := ioutil.WriteFile(fname, b, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
	}
}

func ReadProtoFromJSON(fname string, dest proto.Message) {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can't read file", fname, err)
	}

	if err := protojson.Unmarshal(in, dest); err != nil {
		log.Fatalln("Can't unmarshal", err)
	}
}

func WriteToJSONSample() {
	u := dummyUser()

	WriteProtoToJSON(&u, "superman.json")
}

func ReadFromJSONSample() {
	dest := new(basic.User)

	ReadProtoFromJSON("superman.json", dest)

	log.Println(dest)
}

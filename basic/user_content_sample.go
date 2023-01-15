package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicWriteUserContentV1() {
	uc := basic.UserContent{
		UserContentId: 1001,
		Slug:          "/this-is-v1",
		Title:         "10 Strongest Heroes at Multiverse",
		HtmlContent:   "<p>Just a dummy content for 10 Strongest Heroes</p>",
		AuthorId:      99,
	}

	WriteProtoToFile(&uc, "user_content_v1.bin")
}

func BasicReadUserContentV1() {
	log.Println("Read V1")

	uc := new(basic.UserContent)
	ReadProtoFromFile("user_content_v1.bin", uc)
	log.Println(uc)

	ucJson, _ := protojson.Marshal(uc)
	log.Println(string(ucJson))
	log.Println()
}

func BasicWriteUserContentV2() {
	uc := basic.UserContent{
		UserContentId: 1002,
		Slug:          "/this-is-v2",
		Title:         "Villains with God-like Power",
		HtmlContent:   "<p>Just a dummy content for God-Like Villains</p>",
		AuthorId:      98,
		// Category: "villain",
	}

	WriteProtoToFile(&uc, "user_content_v2.bin")
}

func BasicReadUserContentV2() {
	log.Println("Read V2")

	uc := new(basic.UserContent)
	ReadProtoFromFile("user_content_v2.bin", uc)
	log.Println(uc)

	ucJson, _ := protojson.Marshal(uc)
	log.Println(string(ucJson))
	log.Println()
}

func BasicWriteUserContentV3() {
	uc := basic.UserContent{
		UserContentId: 1003,
		Slug:          "/this-is-v3",
		HtmlContent:   "<p>Just a dummy content for Heroes of Earth</p>",
		// Category:      "hero",
		// SubCategory:   "earth",
	}

	WriteProtoToFile(&uc, "user_content_v3.bin")
}

func BasicReadUserContentV3() {
	log.Println("Read V3")

	uc := new(basic.UserContent)
	ReadProtoFromFile("user_content_v3.bin", uc)
	log.Println(uc)

	ucJson, _ := protojson.Marshal(uc)
	log.Println(string(ucJson))
	log.Println()
}

func BasicWriteUserContentV4() {
	uc := basic.UserContent{
		UserContentId: 1004,
		Slug:          "/this-is-v4",
		Title:         "Heroes with Insect Power",
		AuthorId:      96,
		// Rating:        4,
	}

	WriteProtoToFile(&uc, "user_content_v4.bin")
}

func BasicReadUserContentV4() {
	log.Println("Read V4")

	uc := new(basic.UserContent)
	ReadProtoFromFile("user_content_v4.bin", uc)
	log.Println(uc)

	log.Println("HTML Content is ", uc.HtmlContent)

	ucJson, _ := protojson.Marshal(uc)
	log.Println(string(ucJson))
	log.Println()
}

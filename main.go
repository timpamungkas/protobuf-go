package main

import (
	"fmt"
	"log"
	"my-protobuf/basic"
	"time"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05") + " " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	// basic.BasicHello()
	basic.BasicUser()
	// basic.ProtoToJsonUser()
	// basic.JsonToProtoUser()
	// basic.BasicHello()
	// basic.BasicUserGroup()
	// jobsearch.JobsearchCandidate()
	// jobsearch.JobsearchSoftware()
	// basic.BasicUnmarshalAnyKnown()
	// basic.BasicUnmarshalAnyNotKnown()
	// basic.BasicUnmarshalAnyIs()
	// basic.BasicOneOf()
	// basic.WriteToFileSample()
	// basic.ReadFromFileSample()
	// basic.WriteToJSONSample()
	// basic.ReadFromJSONSample()
	// basic.BasicWriteUserContentV1()
	// basic.BasicReadUserContentV1()
	// basic.BasicWriteUserContentV2()
	// basic.BasicReadUserContentV2()
	// basic.BasicWriteUserContentV3()
	// basic.BasicReadUserContentV3()
	// basic.BasicWriteUserContentV4()
	// basic.BasicReadUserContentV4()
	// basic.BasicReadUserPayment()
}

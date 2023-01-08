package jobsearch

import (
	"log"
	"my-protobuf/protogen/basic"
	"my-protobuf/protogen/dummy"
	"my-protobuf/protogen/jobsearch"

	"google.golang.org/protobuf/encoding/protojson"
)

func JobsearchSoftware() {
	js := jobsearch.JobSoftware{
		JobSoftwareId: 777,
		Application: &basic.Application{
			Version:   "1.0.0",
			Name:      "The amazing proto",
			Platforms: []string{"Mac", "Linux", "Windows"},
		},
	}

	jsonBytes, _ := protojson.Marshal(&js)
	log.Printf("Software : %v\n", string(jsonBytes))
}

func JobsearchCandidate() {
	jc := jobsearch.JobCandidate{
		JobCandidateId: 888,
		Application: &dummy.Application{
			ApplicationId:     887,
			ApplicantFullName: "Shazam",
			Phone:             "999-999-999",
			Email:             "shazam@movie.com",
		},
	}

	jsonBytes, _ := protojson.Marshal(&jc)
	log.Printf("Candidate : %v\n", string(jsonBytes))
}

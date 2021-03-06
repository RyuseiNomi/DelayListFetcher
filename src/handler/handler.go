package handler

import (
	"log"
	"os"

	urlFetcher "github.com/RyuseiNomi/delay_reporter_lm/src/fetcher"
	jsonWorker "github.com/RyuseiNomi/delay_reporter_lm/src/json"
	jsonConverter "github.com/RyuseiNomi/delay_reporter_lm/src/json/converter"
	s3Worker "github.com/RyuseiNomi/delay_reporter_lm/src/s3"
)

func Handler() {
	delayList := urlFetcher.GetDelayList()

	sess, err := s3Worker.GetCredential()
	if err != nil {
		log.Fatal("Load Credential File Error:  %+v\n", err)
	}

	s3Worker.DownloadTrainMaster(sess)

	convertedDelayList, err := jsonConverter.ConvertDelayList(delayList)
	if err != nil {
		log.Fatal(err)
	}

	if err := jsonWorker.CreateJSON(convertedDelayList); err != nil {
		log.Fatal("Create JSON Error:  %+v\n", err)
	}

	jsonFile, err := os.OpenFile("/tmp/json/delay-list.json", os.O_CREATE, 0777)
	if err != nil {
		log.Fatal("Open JSON File Error:  %+v\n", err)
	}
	defer jsonFile.Close()

	if err = s3Worker.Upload(jsonFile, sess); err != nil {
		log.Fatal("Upload JSON Error:  %+v\n", err)
	}

	log.Printf("Finished all process!")
}

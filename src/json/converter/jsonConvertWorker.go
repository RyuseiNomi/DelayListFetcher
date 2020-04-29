package JsonConverter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// URLより取得した遅延リスト
type FetchedDelayList struct {
	Name       string
	Company    string
	LastUpdate string
	Source     string
}

// 全路線のマスターデータTSV
type TrainRoutes struct {
	Company   []string
	RouteName []string
	Region    []string
	isValid   []string
}

type ConvertedDelayList struct {
	Name    string
	Company string
	Region  string
	Status  string
	Source  string
}

//type JsonConvertWorker struct {
//	DelayList FetchedDelayList
//}
//
//func newJsonConvertWorker(delayList []byte) FetchedDelayList {
//
//}

func ConvertDelayList(delayList []byte) {

	// Unmarshal JSON bytes to Go Object
	var fetchedDelayList []FetchedDelayList
	json.Unmarshal(delayList, &fetchedDelayList)

	// Read train route master tsv file
	tsv, err := os.Open("/tmp/tsv/train.tsv")
	if err != nil {
		log.Fatal("TSV Read Error:  %+v\n", err)
	}
	defer tsv.Close()

	render := csv.NewReader(tsv)

	for {
		record, err := render.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("TSV Render Error:  %+v\n", err)
		}
		fmt.Println(record)
	}
}

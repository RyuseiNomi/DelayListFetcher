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
type TrainRoutes []TrainRoute
type TrainRoute struct {
	Company   string
	RouteName string
	Region    string
	isValid   string
}

// ステータスなどを追加した変換後の遅延リスト
type ConvertedDelayLists []ConvertedDelayList
type ConvertedDelayList struct {
	Name    string
	Company string
	Region  string
	Status  string
	Source  string
}

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
	render.Comma = '\t' // change delimiter

	// Append all of train route
	var trainRoutes TrainRoutes
	for {
		record, err := render.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("TSV Render Error:  %+v\n", err)
		}
		trainRoute := TrainRoute{
			Company:   record[0],
			RouteName: record[1],
			Region:    record[2],
			isValid:   record[3],
		}
		trainRoutes = append(trainRoutes, trainRoute)
	}

	// determine delay status each train route
	var convertedDelayLists ConvertedDelayLists
	for _, trainRoute := range trainRoutes {
		var status = "○"
		for _, delayRoute := range fetchedDelayList {
			if delayRoute.Name == trainRoute.RouteName {
				status = "△"
			}
		}
		convertedDelayList := ConvertedDelayList{
			Name:    trainRoute.RouteName,
			Company: trainRoute.Company,
			Region:  trainRoute.Region,
			Status:  status,
			Source:  "鉄道com RSS",
		}
		convertedDelayLists = append(convertedDelayLists, convertedDelayList)
	}
	fmt.Println(convertedDelayLists)
}

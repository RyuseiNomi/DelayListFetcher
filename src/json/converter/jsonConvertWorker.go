package JsonConverter

import (
	"encoding/json"

	tasks "github.com/RyuseiNomi/delay_reporter_lm/src/json/converter/tasks"
	lists "github.com/RyuseiNomi/delay_reporter_lm/src/routeLists"
)

func ConvertDelayList(delayList []byte) ([]byte, error) {

	// Unmarshal JSON bytes to Go Object
	var fetchedDelayList []lists.FetchedDelayList
	json.Unmarshal(delayList, &fetchedDelayList)

	trainRoutes, err := tasks.ReadAllTrainRouteFromTsv()
	if err != nil {
		return nil, err
	}

	// determine delay status each train route
	var convertedDelayLists lists.ConvertedDelayLists
	for _, trainRoute := range trainRoutes {
		var status = "○"
		for _, delayRoute := range fetchedDelayList {
			if delayRoute.Name == trainRoute.RouteName {
				status = "△"
			}
		}
		convertedDelayList := lists.ConvertedDelayList{
			Name:    trainRoute.RouteName,
			Company: trainRoute.Company,
			Region:  trainRoute.Region,
			Status:  status,
			Source:  "鉄道com RSS",
		}
		convertedDelayLists = append(convertedDelayLists, convertedDelayList)
	}

	// Marshal ConvertedDelayLists into bytes
	bytes, _ := json.Marshal(convertedDelayLists)

	return bytes, nil
}

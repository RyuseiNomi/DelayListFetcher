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

	convertedDelayLists := tasks.ConvertDelayLists(fetchedDelayList, trainRoutes)

	// Marshal ConvertedDelayLists into bytes
	bytes, _ := json.Marshal(convertedDelayLists)

	return bytes, nil
}

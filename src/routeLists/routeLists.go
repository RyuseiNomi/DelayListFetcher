package routeLists

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
	IsValid   string
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

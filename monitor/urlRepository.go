package monitor

var urlCollection = make(map[string]*URLData)

func insertIntoURLCollection(id string, data *URLData) {

	urlCollection[id] = data
}

func getURLDataFromCollection(id string) (*URLData, bool) {

	downloadItem, ok := urlCollection[id]

	if ok {
		return downloadItem, ok
	}
	return nil, ok
}

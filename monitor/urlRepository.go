package monitor

var urlCollection = make(map[string]*URLData)

func insertIntoURLCollection(id string, data *URLData) {

	urlCollection[id] = data
}

func getURLDataFromCollection(id string) (*URLData, bool) {

	downloadItem, isPresent := urlCollection[id]

	if isPresent {
		return downloadItem, isPresent
	}
	return nil, isPresent
}

func removeURLFromCollection(id string) bool {

	if isURLPresentInCollection(id) {
		delete(urlCollection, id)
		return true
	}
	return false
}

func isURLPresentInCollection(id string) bool {
	_, check := urlCollection[id]
	return check
}

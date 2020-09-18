package monitor

func GetURLDataByID(id string) (*URLData, bool) {
	return getURLDataFromCollection(id)
}

func UpdateURL(id string, request URLPatchRequest) (*URLData, bool) {

	urlInfo, isPresent := getURLDataFromCollection(id)
	if !isPresent {
		return nil, false
	}
	urlInfo.Frequency = request.Frequency
	urlInfo.Status = request.Status
	urlInfo.FailureCount = 0
	return urlInfo, true
}

func ActivateURL(id string) (string, bool, bool) {
	urlInfo, isPresent := getURLDataFromCollection(id)
	if !isPresent {
		return "", false, false
	} else if urlInfo.Status == ACTIVE {
		return "", true, true
	}
	urlInfo.Status = ACTIVE
	urlInfo.FailureCount = 0
	go monitor(urlInfo)
	return urlInfo.URL, true, false
}

func DeactivateURL(id string) (string, bool, bool) {
	urlInfo, isPresent := getURLDataFromCollection(id)
	if !isPresent {
		return "", false, false
	} else if urlInfo.Status == INACTIVE {
		return "", true, true
	}
	urlInfo.Status = INACTIVE
	return urlInfo.URL, true, false
}

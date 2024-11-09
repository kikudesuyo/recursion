package list

func WebsitePagination(urls []string, pageSize int32, page int32) []string {
	startIndex := int(pageSize * (page - 1))
	currentPageUrls := []string{}
	for i := startIndex; i < int(pageSize)+startIndex && i < len(urls); i++ {
		currentPageUrls = append(currentPageUrls, urls[i])
	}
	return currentPageUrls
}

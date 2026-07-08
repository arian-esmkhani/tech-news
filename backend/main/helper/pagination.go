package helper

func HasNext(page, total, pageSize int64) bool {
	totalPages := (total + pageSize -1) / pageSize
	return page + 1 < totalPages
}

func HasPerv(page int) bool {
	return page > 0
}
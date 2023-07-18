package utils

import "beego_learning/models"

func PageUtil(count int64, pageCu int, pageSize int) models.Page {
	tp := count / int64(pageSize)
	if count%int64(pageSize) > 0 {
		tp = count/int64(pageSize) + 1
	}
	return models.Page{PageCurrent: pageCu, PageSize: pageSize, TotalPage: int(tp), TotalCount: count, FirstPage: pageCu == 1, LastPage: pageCu == int(tp)}
}

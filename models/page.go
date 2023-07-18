package models

type Page struct {
	PageCurrent int
	PageSize    int
	TotalPage   int
	TotalCount  int64
	FirstPage   bool
	LastPage    bool
}

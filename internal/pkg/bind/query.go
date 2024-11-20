package bind

type QueryPage struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit,default=500"`
}

type QueryPage2 struct {
	PageNo   int64 `form:"page_no"`
	PageSize int64 `form:"page_size,default=20"`
}

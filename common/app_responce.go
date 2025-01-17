package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessRes(data, paging, filter interface{}) *successRes {
	res := &successRes{data, paging, filter}
	return res
}

func SimpleSuccessResponce(data interface{}) *successRes {
	return NewSuccessRes(data, nil, nil)
}

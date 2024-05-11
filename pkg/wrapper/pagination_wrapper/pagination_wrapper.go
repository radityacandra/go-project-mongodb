package pagination_wrapper

type Meta struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
	MaxPage int `json:"maxPage"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func PaginationWrapper(page, perPage, maxPage int, data interface{}) *Response {
	return &Response{
		Data: data,
		Meta: Meta{
			Page:    page,
			PerPage: perPage,
			MaxPage: maxPage,
		},
	}
}

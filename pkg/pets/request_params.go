package pets

type ListPetsRequestParams struct {
	Limit *int64 `max:"100" queryParam:"limit"`
}

func (params *ListPetsRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}

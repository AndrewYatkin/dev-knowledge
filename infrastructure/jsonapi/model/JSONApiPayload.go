package jsonApiModel

type JSONApiPayload struct {
	Meta     JSONApiMeta      `json:"meta"`
	Data     []*JSONApiObject `json:"data"`
	Included []*JSONApiObject `json:"included"`
}

func NewEmptyJSONApiPayload() *JSONApiPayload {
	return &JSONApiPayload{
		Meta:     make(JSONApiMeta, 0),
		Data:     make([]*JSONApiObject, 0),
		Included: make([]*JSONApiObject, 0),
	}
}

func (r *JSONApiPayload) AddMeta(key string, data interface{}) {
	r.Meta[key] = data
}

func (r *JSONApiPayload) AddData(data ...*JSONApiObject) {
	r.Data = append(r.Data, data...)
}

func (r *JSONApiPayload) AddInclude(includes ...*JSONApiObject) {
	r.Included = append(r.Included, includes...)
}

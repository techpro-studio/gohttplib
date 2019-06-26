package gohttplib



type Parameters = map[string]interface{}

type IDContainer struct {
	ID string `json:"id"`
}

func NewIDContainer(id string)*IDContainer{
	return &IDContainer{id}
}

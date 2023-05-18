package client

type UnicornItem struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Colour string `json:"colour"`
}

type Unicorn struct {
	ID     string `json:"_id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Colour string `json:"colour"`
}

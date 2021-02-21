package todo

type List struct {
	ID   int    `json:"id"`
	NAME string `json:"name"`
}

type ListWithItems struct {
	ID    int    `json:"id"`
	Items []Item `json:"items"`
}

type Item struct {
	ID   int    `json:"id"`
	TEXT string `json:"text"`
	DONE bool   `json:"done"`
}

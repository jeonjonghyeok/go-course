package todo

type List struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ListwithItem struct {
	List
	Items []Item `json:"items"`
}

type Item struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

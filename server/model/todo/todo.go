package todo

type Todo struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	IsChecked bool   `json:"is_checked"`
}

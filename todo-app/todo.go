package todo

type TodoList struct {
	Id int `json:"id"'`
	Title string `json:"title"'`
	Description string `json:"description"`
}

type UserList struct {
	id int
	UserId int
	ListId int
}

type TodoItem struct {
	Id int `json:"id"'`
	Title string `json:"title"`
	Descriptional string `json:"descriptional"`
	Done bool `json:"done"`
}

type ListsItem struct {
	Id int
	ListId int
	ItemId int
}
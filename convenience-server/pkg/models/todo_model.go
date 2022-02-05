package models

type Todo struct {
	Title    string `json:"title,omitempty" binding:"required"`
	Content  string `json:"content,omitempty" binding:"required"`
	Status   bool   `json:"status,omitempty"`
	UserName string `json:"username,omitempty" binding:"required"`
	CreateAt string `json:"create_at,omitempty"`
}

type UpdateTodo struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	UserName string `json:"username"`
}

type TodoStatus struct {
	Status bool `json:"status,omitempty"`
}

package models

type Todo struct {
	ID              int    `json:"id"`
	ActivityGroupID int    `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        int    `json:"is_active"`
	Priority        string `json:"priority"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type ListTodo struct {
	ID              int         `json:"id"`
	ActivityGroupID int         `json:"activity_group_id"`
	Title           string      `json:"title"`
	IsActive        int         `json:"is_active"`
	Priority        string      `json:"priority"`
	CreatedAt       string      `json:"created_at"`
	UpdatedAt       string      `json:"updated_at"`
	DeletedAt       interface{} `json:"deleted_at"`
}

type RequestCreateTodo struct {
	ActivityGroupID int    `json:"activity_group_id" binding:"required"`
	Title           string `json:"title" binding:"required"`
}

type RequestUpdateTodo struct {
	Title    string `json:"title" binding:"required"`
	IsActive int    `json:"is_active" binding:"required"`
	Priority string `json:"priority" binding:"required"`
}

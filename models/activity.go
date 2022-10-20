package models

type Activity struct{}

type RequestCreateActivity struct {
	Title string `json:"title" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type ActivityCreate struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Email     string `json:"email"`
}

type RequestUpdateActivity struct {
	Title string `json:"title" binding:"required"`
}

type ActivityUpdate struct {
	ID        int         `json:"id"`
	Email     string      `json:"email"`
	Title     string      `json:"title"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	DeletedAt interface{} `json:"deleted_at"`
}

type ResponseActivity struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

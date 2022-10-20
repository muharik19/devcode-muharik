package repositories

import (
	"fmt"
	"log"
	"strconv"

	"github.com/devcode-muharik/models"

	dbconnect "github.com/devcode-muharik/databases"
)

func lastTodoCreateID(id int) (response *models.Todo) {
	var title, priority, createdAt, updatedAt string
	var activityGroupId, isActive int
	query := fmt.Sprintf(`SELECT id, activity_group_id, title, is_active, priority, created_at, updated_at FROM todo WHERE deleted_at is null AND id = %d;`, id)
	dbconnect.QueryRow(query).Scan(
		&id,
		&activityGroupId,
		&title,
		&isActive,
		&priority,
		&createdAt,
		&updatedAt,
	)
	d := &models.Todo{
		ID:              id,
		ActivityGroupID: activityGroupId,
		Title:           title,
		IsActive:        isActive,
		Priority:        priority,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
	response = d
	return
}

func CreatedTodo(request *models.RequestCreateTodo) (response *models.Todo) {
	var id int
	// mengkoneksikan ke db mysql
	db, _ := dbconnect.ConnectDb()
	_, errQuery := db.Exec("insert into todo (activity_group_id, title, priority) values (?, ?, ?)", request.ActivityGroupID, request.Title, "very-high")
	if errQuery != nil {
		fmt.Println(errQuery.Error())
		return
	}
	defer db.Close()
	last := fmt.Sprintf(`SELECT max(id) as id FROM todo WHERE deleted_at is null;`)
	dbconnect.QueryRow(last).Scan(&id)
	response = lastTodoCreateID(id)
	return
}

func UpdateTodo(request *models.RequestUpdateTodo, id int) (response *models.Todo) {
	// mengkoneksikan ke db mysql
	db, _ := dbconnect.ConnectDb()
	_, errQuery := db.Exec("update todo set title = ?, is_active = ?, priority = ? where id = ?", request.Title, request.IsActive, request.Priority, id)
	if errQuery != nil {
		fmt.Println(errQuery.Error())
		return
	}
	defer db.Close()
	response = lastTodoCreateID(id)
	return
}

func DeleteTodo(id int) (response *models.Todo) {
	response = lastTodoCreateID(id)
	// mengkoneksikan ke db mysql
	db, _ := dbconnect.ConnectDb()
	_, errQuery := db.Exec("update todo set deleted_at = now() where id = ?", id)
	if errQuery != nil {
		fmt.Println(errQuery.Error())
		return
	}
	db.Close()
	return
}

func ListTodoAll(activity_group_id int) (response *models.ResponseActivity) {
	arrData := []models.ListTodo{}
	fliter := ""
	if activity_group_id > 0 {
		fliter = fmt.Sprintf(`and activity_group_id = %d`, activity_group_id)
	}
	query := fmt.Sprintf(`select id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at from todo where deleted_at is null %s;`, fliter)

	rowsQ, err := dbconnect.Query(query)
	rowData := models.ListTodo{}

	if err != nil {
		log.Printf(err.Error())
		return
	}

	for rowsQ.Next() {
		err = rowsQ.Scan(
			&rowData.ID,
			&rowData.ActivityGroupID,
			&rowData.Title,
			&rowData.IsActive,
			&rowData.Priority,
			&rowData.CreatedAt,
			&rowData.UpdatedAt,
			&rowData.DeletedAt,
		)

		if err != nil {
			log.Printf(err.Error())
			return
		}
		arrData = append(arrData, rowData)
	}

	d := &models.ResponseActivity{
		Status:  "Success",
		Message: "Success",
		Data:    arrData,
	}

	response = d
	return
}

func ListTodoDetail(id int) (response *models.ResponseActivity) {
	var title, priority, createdAt, updatedAt string
	var activityGroupId, isActive int
	var deletedAt interface{}
	query := fmt.Sprintf(`SELECT id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at from todo WHERE deleted_at is null AND id = %d;`, id)
	dbconnect.QueryRow(query).Scan(
		&id,
		&activityGroupId,
		&title,
		&isActive,
		&priority,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)
	if activityGroupId > 0 {
		d := &models.ResponseActivity{
			Status:  "Success",
			Message: "Success",
			Data: &models.ListTodo{
				ID:              id,
				ActivityGroupID: activityGroupId,
				Title:           title,
				IsActive:        isActive,
				Priority:        priority,
				CreatedAt:       createdAt,
				UpdatedAt:       updatedAt,
				DeletedAt:       deletedAt,
			},
		}
		response = d
		return
	}
	d := &models.ResponseActivity{
		Status:  "Not Found",
		Message: "Activity with ID " + strconv.Itoa(id) + " Not Found",
		Data:    &models.Activity{},
	}
	response = d
	return
}

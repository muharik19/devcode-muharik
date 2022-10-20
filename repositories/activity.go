package repositories

import (
	"fmt"
	"log"
	"strconv"

	"github.com/devcode-muharik/models"

	dbconnect "github.com/devcode-muharik/databases"
)

func lastCreateID(id int) (response *models.ResponseActivity) {
	var email, title, createdAt, updatedAt string
	query := fmt.Sprintf(`SELECT id, title, email, created_at, updated_at FROM activity WHERE deleted_at is null AND id = %d;`, id)
	dbconnect.QueryRow(query).Scan(
		&id,
		&title,
		&email,
		&createdAt,
		&updatedAt,
	)
	d := &models.ResponseActivity{
		Status:  "Success",
		Message: "Success",
		Data: &models.ActivityCreate{
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			ID:        id,
			Title:     title,
			Email:     email,
		},
	}
	response = d
	return
}

func CreatedActivity(request *models.RequestCreateActivity) (response *models.ResponseActivity) {
	var email, title, createdAt, updatedAt string
	var id int
	check := fmt.Sprintf(`SELECT id, title, email, created_at, updated_at FROM activity WHERE email = '%s' AND deleted_at is null;`, request.Email)
	dbconnect.QueryRow(check).Scan(
		&id,
		&title,
		&email,
		&createdAt,
		&updatedAt,
	)
	if id > 0 {
		d := &models.ResponseActivity{
			Status:  "Email Already Exist",
			Message: "Email Already Exist",
			Data: &models.ActivityCreate{
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
				ID:        id,
				Title:     title,
				Email:     email,
			},
		}
		response = d
		return
	}
	// mengkoneksikan ke db mysql
	db, _ := dbconnect.ConnectDb()
	_, errQuery := db.Exec("insert into activity (title, email) values (?, ?)", request.Title, request.Email)
	if errQuery != nil {
		fmt.Println(errQuery.Error())
		return
	}
	defer db.Close()
	last := fmt.Sprintf(`SELECT max(id) as id FROM activity WHERE deleted_at is null;`)
	dbconnect.QueryRow(last).Scan(&id)
	response = lastCreateID(id)
	return
}

func lastUpdateID(id int) (response *models.ResponseActivity) {
	var email, title, createdAt, updatedAt string
	var deletedAt interface{}
	query := fmt.Sprintf(`SELECT id, title, email, created_at, updated_at, deleted_at FROM activity WHERE deleted_at is null AND id = %d;`, id)
	dbconnect.QueryRow(query).Scan(
		&id,
		&title,
		&email,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)
	if email != "" {
		d := &models.ResponseActivity{
			Status:  "Success",
			Message: "Success",
			Data: &models.ActivityUpdate{
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
				DeletedAt: deletedAt,
				ID:        id,
				Title:     title,
				Email:     email,
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

func UpdateActivity(request *models.RequestUpdateActivity, id int) (response *models.ResponseActivity) {
	// mengkoneksikan ke db mysql
	db, _ := dbconnect.ConnectDb()
	_, errQuery := db.Exec("update activity set title = ? where id = ?", request.Title, id)
	if errQuery != nil {
		fmt.Println(errQuery.Error())
		return
	}
	defer db.Close()
	response = lastUpdateID(id)
	return
}

func DeleteActivity(id int) (response *models.ResponseActivity) {
	var email string
	check := fmt.Sprintf(`SELECT email FROM activity WHERE id = %d AND deleted_at is null;`, id)
	dbconnect.QueryRow(check).Scan(&email)
	if email == "" {
		d := &models.ResponseActivity{
			Status:  "Not Found",
			Message: "Activity with ID " + strconv.Itoa(id) + " Not Found",
			Data:    &models.Activity{},
		}
		response = d
		return
	}
	// mengkoneksikan ke db mysql
	db, _ := dbconnect.ConnectDb()
	_, errQuery := db.Exec("update activity set deleted_at = now() where id = ?", id)
	if errQuery != nil {
		fmt.Println(errQuery.Error())
		return
	}
	defer db.Close()
	d := &models.ResponseActivity{
		Status:  "Success",
		Message: "Success",
		Data:    &models.Activity{},
	}
	response = d
	return
}

func ListActivityAll() (response *models.ResponseActivity) {
	arrData := []models.ActivityUpdate{}
	query := fmt.Sprintf(`select id, email, title, created_at, updated_at, deleted_at from activity where deleted_at is null;`)

	rowsQ, err := dbconnect.Query(query)
	rowData := models.ActivityUpdate{}

	if err != nil {
		log.Printf(err.Error())
		return
	}

	for rowsQ.Next() {
		err = rowsQ.Scan(
			&rowData.ID,
			&rowData.Email,
			&rowData.Title,
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

func ListActivityDetail(id int) (response *models.ResponseActivity) {
	var email, title, createdAt, updatedAt string
	var deletedAt interface{}
	query := fmt.Sprintf(`SELECT id, email, title, created_at, updated_at, deleted_at FROM activity WHERE deleted_at is null AND id = %d;`, id)
	dbconnect.QueryRow(query).Scan(
		&id,
		&email,
		&title,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)
	if email != "" {
		d := &models.ResponseActivity{
			Status:  "Success",
			Message: "Success",
			Data: &models.ActivityUpdate{
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
				DeletedAt: deletedAt,
				ID:        id,
				Title:     title,
				Email:     email,
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

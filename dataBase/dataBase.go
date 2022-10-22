package dataBase

import (
	"database/sql"
	"errors"
	"v4/models"
)

func CreateProd(prod *models.Product, db *sql.DB) (int64, error) {

	sql := "INSERT INTO product(name, description, value, status) VALUES( ?, ?, ?, ?)"

	stmt, err := db.Prepare(sql)

	if err != nil {
		return 0, errors.New("Erro ao inserir produto.")
	}
	defer stmt.Close()
	result, err := stmt.Exec(prod.Name, prod.Description, prod.Value, prod.Status)

	if err != nil {
		return 0, errors.New("Erro ao inserir produto.")
	}
	id, _ := result.LastInsertId()
	return id, nil
}
func GetProd(id int64, db *sql.DB) (prod models.Product, err error) {
	query := "SELECT id, name, description, value, status FROM product WHERE id = ?"

	err = db.QueryRow(query, id).Scan(&prod.Id, &prod.Name, &prod.Description, &prod.Value, &prod.Status)

	if err != nil {
		return
	}

	return
}

func DeleteProd(id int64, db *sql.DB) error {
	sql := "DELETE FROM product WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return errors.New("Erro ao excluir produto.")
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New("Erro ao excluir produto.")
	}

	return nil

}

func UpdateProd(prod *models.Product, db *sql.DB) error {
	sqlStatement := "UPDATE product SET name = ?, description = ?, status = ?, value = ? WHERE id = ?"
	_, err := db.Exec(sqlStatement, prod.Name, prod.Description, prod.Status, prod.Value, prod.Id)

	return err

}

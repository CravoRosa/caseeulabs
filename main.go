package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"v4/dataBase"
	"v4/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	db, err := sql.Open("mysql", "root:root@(localhost:3306)/caseeulabs")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	e.POST("/product", func(c echo.Context) error {
		prod := new(models.Product)
		if err := c.Bind(prod); err != nil {
			return err
		}
		id, err := dataBase.CreateProd(prod, db)
		if err != nil {
			return c.JSON(400, err)
		}
		if id == 0 {
			return c.JSON(400, "Não foi possível criar o produto.")
		}

		return c.JSON(http.StatusCreated, id)
	})

	e.DELETE("/product/:id", func(c echo.Context) error {
		id := c.Param("id")
		ids, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(400, "Deve ser um Inteiro.")
		}
		err = dataBase.DeleteProd(int64(ids), db)
		if err != nil {
			return c.JSON(400, "Erro ao deletar produto.")
		}
		return c.JSON(http.StatusOK, "Deletado.")
	})

	e.GET("/product/:id", func(c echo.Context) error {
		id := c.Param("id")
		ids, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(400, "Deve ser um Inteiro.")
		}

		prod, err := dataBase.GetProd(int64(ids), db)
		if err != nil {
			return c.JSON(400, "Produto não encontrado.")
		}

		return c.JSON(http.StatusOK, prod)
	})

	e.PUT("/product", func(c echo.Context) error {
		prod := new(models.Product)
		if err := c.Bind(prod); err != nil {
			return err
		}
		err = dataBase.UpdateProd(prod, db)
		if err != nil {
			return c.JSON(400, "Não foi possivel atualizar o produto")
		}
		return c.String(http.StatusOK, "Atualizado.")
	})

	e.Start(":9000")
}

package models

import (
	"fmt"
	"net/http"
	"new-project2/db"
)

type pegawai struct {
	Id		  int		`json:"id"`
	Nama      string	`json:"nama"`
	Pekerjaan string	`json:"pekerjaan"`
	Umur      string	`json:"umur"`
	Telepon   string	`json:"telepon"`
	Email     string	`json:"email"`
	Password  string	`json:"password"`
}

func FetchPegawai() (Response, error) {
	var obj pegawai
	var arrobj []pegawai
	var res Response

	con := db.CreateConn()

	sqlFetch := "SELECT * FROM restful"

	rows, err := con.Query(sqlFetch)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&obj.Id,&obj.Nama, &obj.Pekerjaan, &obj.Umur, &obj.Telepon, &obj.Email, &obj.Password)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status 	= http.StatusOK
	res.Message = "SUCCESS"
	res.Data 	= arrobj

	return res, nil


}

func StorePegawai(nama string, pekerjaan string, umur string, telepon string, email string, password string) (Response, error) {
	var res Response

	con := db.CreateConn()

	sqlStore := "INSERT INTO restful (nama, pekerjaan, umur, telepon, email, password)  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

	var lastInsertedId int64

	err := con.QueryRow(sqlStore, nama, pekerjaan, umur, telepon, email, password).Scan(&lastInsertedId)
	if err != nil {
		fmt.Printf("error query your database: %v", err)
	}

	res.Status  = http.StatusOK
	res.Message = "SUCCESS"
	res.Data 	= map[string]interface{}{
		"nama" 		:	nama,
		"pekerjaan" :	pekerjaan,
		"umur"		: 	umur,
		"telepon"	: 	telepon,
		"email"		:	email,
		"password"	:	password,
		"last_inserted_id" : lastInsertedId,
	}

	return res, nil
}

func UpdatePegawai(id int, nama string, pekerjaan string, umur string, telepon string, email string, password string ) (Response, error) {
	var res Response

	con := db.CreateConn()

	sqlUpdate := "UPDATE restful SET nama = $1, pekerjaan = $2, umur = $3, telepon = $4, email = $5, password = $6 WHERE id = $7"

	updt, err := con.Prepare(sqlUpdate)
	if err != nil {
		return res, err
	}

	result, err := updt.Exec(nama, pekerjaan, umur, telepon, email, password, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status 	= http.StatusOK
	res.Message = "success"
	res.Data   	= map[string]int64{
		"rows_Affected" : rowsAffected,
	}

	return res, nil
}

func DeletePegawai(id int) (Response, error) {
	var res Response

	con := db.CreateConn()

	sqlDelete := "DELETE FROM restful WHERE id= $1"

	dlte, err  := con.Prepare(sqlDelete)
	if err != nil {
		return res, err
	}

	result, err := dlte.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUCCESS"
	res.Data = map[string]int64{
		"rows_Affected" : rowsAffected,
	}

	return res, nil
}
package models

import (
	"API/structs"
	"fmt"
)

func Logout(id int64) structs.Logout {

	var data structs.Logout

	stmt, err := db.Prepare("UPDATE users SET login_status = ? WHERE id = ?")
	checkErr(err)

	res, err := stmt.Exec(0, id)
	checkErr(err)

	fmt.Println(res)

	data = structs.Logout{"S", id}

	return data

}

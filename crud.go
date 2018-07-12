package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Watchlist struct {
	container   string    `json:"container"`
	using_s string `json:"using_s"`
	update_date string `json:"update_date"`
	id_user string `json:"id_user"`
}

type Credit struct {
	id_user   string    `json:"id_user"`
	credit string `json:"credit"`
}

type Configuration struct {
	container   string    `json:"container"`
	ram string `json:"ram"`
	cpu string `json:"cpu"`
	storage string `json:"storage"`
	id_user string `json:"id_user"`
}


func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "root"
    dbName := "payg"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil { panic(err.Error())}
    return db
}

func insert_credit(id_user string,credit string){

	db := dbConn()
	 insert, err := db.Query("INSERT INTO credit VALUES ('"+id_user+"', '"+credit+"')")
	 if err != nil {panic(err.Error())}

	 defer insert.Close()
}

func update_credit(id_user string,credit string){

	db := dbConn()
	insForm, err := db.Prepare("UPDATE credit SET credit=? WHERE id_user=?")
    if err != nil {panic(err.Error())}

    insForm.Exec(credit,id_user)
    defer db.Close()
}

func select_credit() map[string][]string{

	db := dbConn()
	c := make(map[string][]string)

	results, err := db.Query("SELECT id_user,credit FROM credit")
	if err != nil {panic(err.Error())}

	for results.Next() {
		var credit Credit
		err = results.Scan(&credit.id_user, &credit.credit)
		if err != nil {panic(err.Error())}
		c[credit.id_user] = append(c[credit.id_user], credit.credit)
	}
	defer db.Close()
	return c
}

func main() {

	insert_credit("12","12535644.12")

	fmt.Println(select_credit())

	update_credit("11","0")

	fmt.Println(select_credit())

}

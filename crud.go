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

//file_to_c
func select_watchlist() [][]string{

	db := dbConn()
    var c_file [][]string
    var c []string

	results, err := db.Query("SELECT container,using_s,update_date FROM watchlist")
	if err != nil {panic(err.Error())}

	for results.Next() {
		var watchlist Watchlist
		err = results.Scan(&watchlist.container,&watchlist.using_s,&watchlist.update_date)
		if err != nil {panic(err.Error())}
		
		c = append(c, watchlist.container)
	    c = append(c, watchlist.using_s)
	    c = append(c, watchlist.update_date)

	    c_file = append(c_file, c)
	    c = nil
	}
	defer db.Close()
	return c_file

}

//file_to_c_name
func select_watchlist_name() []string{

	db := dbConn()
    var c_names []string

	results, err := db.Query("SELECT container FROM watchlist")
	if err != nil {panic(err.Error())}

	for results.Next() {
		var watchlist Watchlist
		err = results.Scan(&watchlist.container)
		if err != nil {panic(err.Error())}
		
		c_names = append(c_names, watchlist.container)
	}
	defer db.Close()
	return c_names

}

//conf_to_map
func select_configuration() map[string][]string{

	db := dbConn()
    conf := make(map[string][]string)

	results, err := db.Query("SELECT container,ram,cpu,storage,id_user FROM configuration")
	if err != nil {panic(err.Error())}

	for results.Next() {
		var configuration Configuration
		err = results.Scan(&configuration.container,&configuration.ram,&configuration.cpu,&configuration.storage,&configuration.id_user)
		if err != nil {panic(err.Error())}
		
		conf[configuration.container] = append(conf[configuration.container], configuration.container)
        conf[configuration.container] = append(conf[configuration.container], configuration.ram)
        conf[configuration.container] = append(conf[configuration.container], configuration.cpu)
        conf[configuration.container] = append(conf[configuration.container], configuration.storage)
        conf[configuration.container] = append(conf[configuration.container], configuration.id_user)
	}
	defer db.Close()
	return conf

}


//credit_to_map()
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

//credit_to_file
func update_watchlist(container string,using_s string,update_date string){

	db := dbConn()
	insForm, err := db.Prepare("UPDATE watchlist SET using_s=?,update_date=? WHERE container=?")
    if err != nil {panic(err.Error())}

    insForm.Exec(using_s,update_date,container)
    defer db.Close()
}


//credit_to_file
func update_credit(id_user string,credit string){

	db := dbConn()
	insForm, err := db.Prepare("UPDATE credit SET credit=? WHERE id_user=?")
    if err != nil {panic(err.Error())}

    insForm.Exec(credit,id_user)
    defer db.Close()
}


//add_container_file
func add_new_container(containers [][]string){

        c_names := select_watchlist_name

        for i := range containers {

                if !stringInSlice(containers[i][0], c_names) {

						db := dbConn()
						 insert, err := db.Query("INSERT INTO watchlist VALUES ('"+containers[i][0]+"','0','"+containers[i][3]+"','')")
						 if err != nil {panic(err.Error())}

						 defer insert.Close()

                }
        }
}


func insert_credit(id_user string,credit string){

	db := dbConn()
	 insert, err := db.Query("INSERT INTO credit VALUES ('"+id_user+"', '"+credit+"')")
	 if err != nil {panic(err.Error())}

	 defer insert.Close()
}





func main() {

	// insert_credit("12","12535644.12")
	// fmt.Println(select_credit())
	// update_credit("11","0")
	// fmt.Println(select_credit())

	fmt.Println(select_watchlist_name())
	fmt.Println(select_watchlist())
	fmt.Println(select_configuration())

}

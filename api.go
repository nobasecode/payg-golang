package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
    "strconv"
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

type Log struct {
    container   string    `json:"container"`
    id_user string `json:"id_user"`
    using_s string `json:"using_s"`
    using_moment string `json:"using_moment"`
    update_date string `json:"update_date"`
    ram string `json:"ram"`
    cpu string `json:"cpu"`
    storage string `json:"storage"`
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

func select_use(container string,id_user string,update_date1 string,update_time1 string,update_date2 string,update_time2 string){

    db := dbConn()
    all_use := 0

    update_date1 = update_date1+" "+update_time1
    update_date2 = update_date2+" "+update_time2

    results, err := db.Query("SELECT * FROM log WHERE container=? and id_user=? and update_date between ? and ?", container,id_user,update_date1,update_date2)
    if err != nil {panic(err.Error())}

    fmt.Println("\n\033[32mRecord between ["+update_date1+"] and ["+update_date2+"]\033[39m\n")
    fmt.Println("\033[31m| Using  |     Update time     | RAM  |CPU| Storage |\033[39m")

    for results.Next() {
        var logi Log
        err = results.Scan(&logi.container,&logi.id_user,&logi.using_s,&logi.using_moment,&logi.update_date,&logi.ram,&logi.cpu,&logi.storage)
        if err != nil {panic(err.Error())}
        
        fmt.Println("|"+logi.using_s+" | "+logi.update_date+" | "+logi.ram+" | "+logi.cpu+" | "+logi.storage+"    |")

        using_moment , err:= strconv.Atoi(logi.using_moment)
        if err != nil {fmt.Println(err)}

        all_use = all_use+using_moment
    }

    fmt.Println("\n\033[32mTotale use :"+strconv.Itoa(all_use)+"\033[39m\n")
    defer db.Close()

}


func main() {

    fmt.Print("Container name: ")
    var c string
    fmt.Scanln(&c)  

    fmt.Print("User ID : ")
    var id string
    fmt.Scanln(&id)

    fmt.Print("Start date [example : 2018-07-16] : ")
    var sd string
    fmt.Scanln(&sd)

    fmt.Print("Start time [example : 11:37:52] : ")
    var st string
    fmt.Scanln(&st)    

    fmt.Print("End date [example : 2018-07-16] : ")
    var fd string
    fmt.Scanln(&fd)             

    fmt.Print("End time [example :  11:37:56] : ")
    var ft string
    fmt.Scanln(&ft)   

    select_use(c,id,sd,st,fd,ft)


}

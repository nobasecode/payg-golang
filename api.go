package main

import (
	"fmt"
    "strconv"
    "log"
	"database/sql"
    "net/http"
    "encoding/json"

	_ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)

type Watchlist struct {
	container   string    `json:"container"`
	using_s string `json:"using_s"`
	update_date string `json:"update_date"`
	id_user string `json:"id_user"`
}

type Credit struct {
	Id   string    `json:"id_user"`
	Crd string `json:"credit"`
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

func call_select_use() {
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


func GetCredit(w http.ResponseWriter, r *http.Request) {

    db := dbConn()
    var credit []string

    results, err := db.Query("SELECT id_user,credit FROM credit")
    if err != nil {panic(err.Error())}
    for results.Next() {
        var c Credit
        err = results.Scan(&c.Id, &c.Crd)
        if err != nil {panic(err.Error())}
        credit = append(credit, "{id_user: "+c.Id+", credit: "+c.Crd+"}")
    }

    defer db.Close()
    json.NewEncoder(w).Encode(credit)
    fmt.Println(credit)


}


func GetCredit_json(w http.ResponseWriter, r *http.Request){

    db := dbConn()
    var credit []Credit

    results, err := db.Query("SELECT id_user,credit FROM credit")
    if err != nil {panic(err.Error())}
    for results.Next() {
        var c Credit
        err = results.Scan(&c.Id, &c.Crd)
        if err != nil {panic(err.Error())}
        credit = append(credit, Credit{Id:c.Id, Crd:c.Crd})
    }
    json.NewEncoder(w).Encode(credit)
    defer db.Close()

}

func GetCreditById_json(w http.ResponseWriter, r *http.Request){

    db := dbConn()
    id := r.URL.Query().Get("id")
    var credit Credit

    results, err := db.Query("SELECT id_user,credit FROM credit WHERE id_user=?",id)
    if err != nil {panic(err.Error())}
    for results.Next() {
        var c Credit
        err = results.Scan(&c.Id, &c.Crd)
        if err != nil {panic(err.Error())}
        credit = Credit{Id:c.Id, Crd:c.Crd}
    }
    json.NewEncoder(w).Encode(credit)
    defer db.Close()

}



func GetUse_json(w http.ResponseWriter, r *http.Request){

    db := dbConn()
    container := r.URL.Query().Get("container")
    id_user := r.URL.Query().Get("id_user")
    update_date1 := r.URL.Query().Get("update_date1")
    update_time1 := r.URL.Query().Get("update_time1")
    update_date2 := r.URL.Query().Get("update_date2")
    update_time2 := r.URL.Query().Get("update_time2")

    all_use := 0
    var use []string

    update_date1 = update_date1+" "+update_time1
    update_date2 = update_date2+" "+update_time2

    results, err := db.Query("SELECT * FROM log WHERE container=? and id_user=? and update_date between ? and ?", container,id_user,update_date1,update_date2)
    if err != nil {panic(err.Error())}

    for results.Next() {
        var logi Log
        err = results.Scan(&logi.container,&logi.id_user,&logi.using_s,&logi.using_moment,&logi.update_date,&logi.ram,&logi.cpu,&logi.storage)
        if err != nil {panic(err.Error())}
        
        using_moment , err:= strconv.Atoi(logi.using_moment)
        if err != nil {fmt.Println(err)}

        all_use = all_use+using_moment

        use = append(use, "{container:"+logi.container+", id_user:"+logi.id_user+", using_s:"+logi.using_s+", using_moment:"+logi.using_moment+", update_date:"+logi.update_date+", ram:"+logi.ram+", cpu:"+logi.cpu+", storage:"+logi.storage+"}")
    
    }

    use = append(use, "{total use: "+strconv.Itoa(all_use)+"}")        
    
    fmt.Println(use)
    json.NewEncoder(w).Encode(use)
    
    defer db.Close()

}



func main() {


    //call_select_use()
   //fmt.Println(GetCredit_json())


    router := mux.NewRouter()

    //http://ip:8000/credit
    router.HandleFunc("/credit", GetCredit_json).Methods("GET")

    //http://ip:8000/credit_user?id=12
    router.HandleFunc("/credit_user", GetCreditById_json).Methods("GET")
    	
    //http://172.23.236.111:8000/container_use?container=c1&id_user=11&update_date1=2018-07-16&update_time1=11:37:33&update_date2=2018-07-16&update_time2=11:37:56
    router.HandleFunc("/container_use", GetUse_json).Methods("GET")        
    
    log.Fatal(http.ListenAndServe(":8000", router))


}

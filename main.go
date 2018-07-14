package main

import (
    "fmt"
    "log"
    "os/exec"
    "strings"
    "time"
    "os"
    "strconv"
    "./cron"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

)


//mysqldump -u root -p payg > /home/blackbase/project/payg.sql


const (
    timeFormat = "2006-01-02 15:04:05"
    ram_price = 0.001
    cpu_price = 0.002
    disk_price = 0.001

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

//verify if a string exist in a slice
func stringInSlice(str string, list []string) bool {
    for _, v := range list {
        if v == str {
            return true
        }
    }
    return false
}

//String to date
func s_date(d string) time.Time{
    
    t , err := time.Parse(timeFormat, d)
        if err != nil {
        fmt.Println(err)
    }
    
    return t
}

//Difference bettwen Last-updated & Container-Finished Dates
func diff_date(finished time.Time,lastup time.Time) int{

    diff := finished.Sub(lastup)
    if diff.Seconds() > 0 { 
        return int(diff.Seconds())
    } else { return 0 }

    // diff := lastup.Sub(finished)
    // return int(diff.Seconds())
}

//Difference bettwen Last-updated & Now Dates
func diff_date_now(lastup time.Time) int{


        current_time := time.Now().Local()
        now_d := s_date(current_time.Format(timeFormat))
        diff := now_d.Sub(lastup)
        //diff := time.Since(lastup)
        return int(diff.Seconds())
        
}

//show status of containers
func show_status(a string){
    
    a = a+"\033[39m"
    output, err := exec.Command("echo", "-e",a).CombinedOutput()
    if err != nil {
      os.Stderr.WriteString(err.Error())
    }
    fmt.Println(string(output))
     
}


//get a list of running docker containers
func inspect() [][]string {
    
    var containers [][]string
    var container []string
        
    out, err := exec.Command("sh","-c","docker inspect --format='{{.State.Running}}|{{.Name}}|{{.State.StartedAt}}|{{.State.FinishedAt}}|{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(docker ps -qa)").Output()  
          if err != nil {log.Fatal(err)}
    
    s := strings.Split(string(out), "\n")
    

    for i := range s[:len(s)-1] {

        infos := strings.Split(s[i], "|")

        container = append(container, infos[1][1:3])
        container = append(container, infos[0])         
        container = append(container, infos[4])         
        container = append(container, infos[2][0:10]+" "+infos[2][11:19])
        container = append(container, infos[3][0:10]+" "+infos[3][11:19])

        containers = append(containers, container)
        
        container = nil
    }
    
    return containers

}


//read into watchlist & get container infos
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

//read into watchlist & get container names
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

//Configuration liste
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

//Credit liste
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


func credit_to_file(credit map[string][]string){

    all_c := ""
    show := "|ID|Credit\n"

    for k, v := range credit {
        all_c = all_c+k+"|"+v[0]+"|\n"

        show = show+"|"+k+"|"+v[0]+"\n"
    }    
    fmt.Println(show)
  

}
//credit_to_file
func update_credit(id_user string,credit string){

    db := dbConn()
    insForm, err := db.Prepare("UPDATE credit SET credit=? WHERE id_user=?")
    if err != nil {panic(err.Error())}

    insForm.Exec(credit,id_user)
    defer db.Close()
}


//calculate using /DH
func calculate_use(ram int,cpu int,disk int,s int) float64{

    return float64(float64(ram)*float64(s)*ram_price)+(float64(cpu)*float64(s)*cpu_price)+(float64(disk)*float64(s)*disk_price)
}


//add_container_file
func add_new_container(containers [][]string){

        c_names := select_watchlist_name()

        for i := range containers {

                if !stringInSlice(containers[i][0], c_names) {

                        db := dbConn()
                         insert, err := db.Query("INSERT INTO watchlist VALUES ('"+containers[i][0]+"','0','"+containers[i][3]+"','')")
                         if err != nil {panic(err.Error())}

                         defer insert.Close()

                }
        }
}

func update_watchlist(container string,using_s string,update_date string){

    db := dbConn()
    insForm, err := db.Prepare("UPDATE watchlist SET using_s=?,update_date=? WHERE container=?")
    if err != nil {panic(err.Error())}

    insForm.Exec(using_s,update_date,container)
    defer db.Close()
}


func payg(containers [][]string) {

    add_new_container(containers)
    conf := select_configuration()

    credit := select_credit()

    c_file := select_watchlist()

    // all := ""
    a := "\n"

    for i := range containers {
        
        running , err := strconv.ParseBool(containers[i][1])
        if err != nil {panic(err)}

        if running == true {
            for j := range c_file {
                if containers[i][0] == c_file[j][0] {

                    last_value , err:= strconv.Atoi(c_file[j][1])
                    if err != nil {log.Fatal(err)}
                    date_diff := diff_date_now(s_date(c_file[j][2]))
                    new_time := last_value+date_diff


                    id := conf[containers[i][0]][4]
                    ram , err:= strconv.Atoi(conf[containers[i][0]][1])
                    if err != nil {log.Fatal(err)}
                    cpu , err:= strconv.Atoi(conf[containers[i][0]][2])
                    if err != nil {log.Fatal(err)}
                    disk , err:= strconv.Atoi(conf[containers[i][0]][3])
                    if err != nil {log.Fatal(err)}
                    now_credit, err := strconv.ParseFloat(credit[id][0],64)
                    if err != nil {log.Fatal(err)}


                    use := calculate_use(ram,cpu,disk,date_diff)                    
                    new_credit := now_credit-use
                    credit[id][0] = strconv.FormatFloat(new_credit, 'f', 3, 64)
                    
                    update_credit(id,credit[id][0])

                    t := time.Now()
                    a = a+"\033[32m"+containers[i][0]+" - "+strconv.Itoa(new_time)+" - "+t.Format("2006-01-02 15:04:05")+"\n"                

                    //all = all+containers[i][0]+"|"+strconv.Itoa(new_time)+"|"+t.Format("2006-01-02 15:04:05")+"\n"

                    update_watchlist(containers[i][0],strconv.Itoa(new_time),t.Format("2006-01-02 15:04:05"))

                }
            }
        } else if running == false {
            for j := range c_file {
                if containers[i][0] == c_file[j][0] {

                    last_value , err:= strconv.Atoi(c_file[j][1])
                    if err != nil {log.Fatal(err)}
                    date_diff := diff_date(s_date(containers[i][4]).Add(time.Hour * time.Duration(1) +time.Minute * time.Duration(0) +time.Second * time.Duration(0)),s_date(c_file[j][2]))
                    new_time := last_value+date_diff

                     id := conf[containers[i][0]][4]
                     ram , err:= strconv.Atoi(conf[containers[i][0]][1])
                     if err != nil {log.Fatal(err)}
                     cpu , err:= strconv.Atoi(conf[containers[i][0]][2])
                     if err != nil {log.Fatal(err)}
                     disk , err:= strconv.Atoi(conf[containers[i][0]][3])
                     if err != nil {log.Fatal(err)}
                     now_credit, err := strconv.ParseFloat(credit[id][0],64)
                     if err != nil {log.Fatal(err)}                    


                    use := calculate_use(ram,cpu,disk,date_diff)                    
                    new_credit := now_credit-use
                    credit[id][0] = strconv.FormatFloat(new_credit, 'f', 3, 64)
                    
                    update_credit(id,credit[id][0])

                    t := time.Now()
                    a = a+"\033[31m"+containers[i][0]+" - "+strconv.Itoa(new_time)+" - "+t.Format("2006-01-02 15:04:05")+"\n"

                    update_watchlist(containers[i][0],strconv.Itoa(new_time),t.Format("2006-01-02 15:04:05"))

                }
            }
        }

    }
     
     credit_to_file(credit)
     show_status(a)

}

func main(){

    c := cron.New()
    c.AddFunc("@every 2s", func() { 
        payg(inspect())
     })
    c.Run()
    c.Start()

}

package main

import (
    "fmt"
    "log"
    "os/exec"
    "strings"
    "time"
    "os"
    "io/ioutil"
    "strconv"
    "./cron"

)


const (
    timeFormat = "2006-01-02 15:04:05"
    ram_price = 0.001
    cpu_price = 0.002
    disk_price = 0.001

)

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

    diff := lastup.Sub(finished)
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
func file_to_c(f_name string) [][]string{

    var c_file [][]string
    var c []string

    watchlist, err := os.Open(f_name)
        if err != nil {
                log.Fatal(err)
        }

        lines, err := ioutil.ReadAll(watchlist)
                if err != nil {
                log.Fatal(err)
        }

        line := strings.Split(string(lines), "\n")

        for i := range line[:len(line)-1] {
                line_part := strings.Split(string(line[i]), "|")
                c = append(c, line_part[0])
                c = append(c, line_part[1])
                c = append(c, line_part[2])

                c_file = append(c_file, c)
                c = nil
        }

    watchlist.Close()

    return c_file
}

//read into watchlist & get container names
func file_to_c_name(f_name string) []string{

    watchlist, err := os.Open(f_name)
        if err != nil {
                log.Fatal(err)
        }

        lines, err := ioutil.ReadAll(watchlist)
                if err != nil {
                log.Fatal(err)
        }

        line := strings.Split(string(lines), "\n")

        var c_names []string
        for i := range line[:len(line)-1] {
                line_part := strings.Split(string(line[i]), "|")
                c_names = append(c_names, line_part[0])

        }

    watchlist.Close()

    return c_names
}

//Configuration liste
func conf_to_map() map[string][]string{

    conf := make(map[string][]string)

    conf_file, err := os.Open("conf")
    if err != nil {log.Fatal(err)}

    lines, err := ioutil.ReadAll(conf_file)
    if err != nil {log.Fatal(err)}

    line := strings.Split(string(lines), "\n")

    for i := range line[:len(line)] {
            line_part := strings.Split(string(line[i]), "|")

            conf[line_part[0]] = append(conf[line_part[0]], line_part[0])
            conf[line_part[0]] = append(conf[line_part[0]], line_part[1])
            conf[line_part[0]] = append(conf[line_part[0]], line_part[2])
            conf[line_part[0]] = append(conf[line_part[0]], line_part[3])
            conf[line_part[0]] = append(conf[line_part[0]], line_part[4])
    }

    conf_file.Close()

    return conf
}

//Credit liste
func credit_to_map() map[string][]string{

    credit := make(map[string][]string)

    credit_file, err := os.Open("credit")
    if err != nil {log.Fatal(err)}

    lines, err := ioutil.ReadAll(credit_file)
    if err != nil {log.Fatal(err)}

    line := strings.Split(string(lines), "\n")

    for i := range line[:len(line)-1] {
            line_part := strings.Split(string(line[i]), "|")

            credit[line_part[0]] = append(credit[line_part[0]], line_part[1])
    }

    credit_file.Close()

    return credit
}

func credit_to_file(credit map[string][]string){

    all_c := ""
    show := "|ID|Credit\n"

    for k, v := range credit {
        all_c = all_c+k+"|"+v[0]+"|\n"

        show = show+"|"+k+"|"+v[0]+"\n"
    }    
    fmt.Println(show)

    err := ioutil.WriteFile("credit", []byte(all_c), 0666)
    if err != nil {log.Fatal(err)}    

}

//calculate using /DH
func calculate_use(ram int,cpu int,disk int,s int) float64{

    return float64(float64(ram)*float64(s)*ram_price)+(float64(cpu)*float64(s)*cpu_price)+(float64(disk)*float64(s)*disk_price)
}

//add new containers to file
func add_container_file(containers [][]string){

        c_names := (file_to_c_name("watchlist"))

        for i := range containers {

                if !stringInSlice(containers[i][0], c_names) {

                        f, err := os.OpenFile("watchlist", os.O_APPEND|os.O_WRONLY, 0644)
                        if err != nil {
                                panic(err)
                        }
                        f.WriteString(containers[i][0]+"|0|"+containers[i][3]+"\n")
                        f.Close()
                }
        }
}

func payg(containers [][]string) {

    add_container_file(containers)
    conf := conf_to_map()

    credit := credit_to_map()

    c_file := file_to_c("watchlist")

    // fmt.Println(containers)
    // fmt.Println(conf_to_map())
    // fmt.Println(credit_to_map())


    all := ""
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

                    //fmt.Println(c_file[j][2])
                    //fmt.Println(s_date(c_file[j][2]))
                    
                    //current_time := time.Now().Local()
                    //fmt.Println(s_date(current_time.Format(timeFormat)))

                    //fmt.Println(diff_date(s_date(c_file[j][2]),s_date(current_time.Format(timeFormat))))
                    //fmt.Println(diff_date_now(s_date(c_file[j][2])))


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
                     
                     //fmt.Println(credit)
                     //fmt.Println(now_credit)
                    

                    t := time.Now()
                    a = a+"\033[32m"+containers[i][0]+" - "+strconv.Itoa(new_time)+" - "+t.Format("2006-01-02 15:04:05")+"\n"                

                    all = all+containers[i][0]+"|"+strconv.Itoa(new_time)+"|"+t.Format("2006-01-02 15:04:05")+"\n"
                }
            }
        } else if running == false {
            for j := range c_file {
                if containers[i][0] == c_file[j][0] {

                    last_value , err:= strconv.Atoi(c_file[j][1])
                    if err != nil {log.Fatal(err)}
                    date_diff := diff_date(s_date(containers[i][3]).Add(time.Hour * time.Duration(1) +time.Minute * time.Duration(0) +time.Second * time.Duration(0)),s_date(c_file[j][2]))
                    new_time := last_value+date_diff

                    // fmt.Println(s_date(containers[i][3]).Add(time.Hour * time.Duration(0) +
                    //              time.Minute * time.Duration(0) +
                    //              time.Second * time.Duration(0)))
                    // fmt.Println(s_date(c_file[j][2]))

                    // fmt.Println(diff_date(s_date(containers[i][3]).Add(time.Hour * time.Duration(1) +time.Minute * time.Duration(0) +time.Second * time.Duration(0)),s_date(c_file[j][2])))


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
                    
                    // fmt.Println(credit)
                    // fmt.Println(now_credit)


                    t := time.Now()
                    a = a+"\033[31m"+containers[i][0]+" - "+strconv.Itoa(new_time)+" - "+t.Format("2006-01-02 15:04:05")+"\n"

                    all = all+containers[i][0]+"|"+strconv.Itoa(new_time)+"|"+t.Format("2006-01-02 15:04:05")+"\n"

                }
            }
        }

    }

    err := ioutil.WriteFile("watchlist", []byte(all), 0666)
    if err != nil {log.Fatal(err)}
    
    //fmt.Println(containers)
    credit_to_file(credit)
    
    show_status(a)

}

func main(){

    c := cron.New()
    c.AddFunc("@every 2s", func() { 
        payg(inspect())
        //fmt.Println(inspect())

     })
    c.Run()
    c.Start()

}

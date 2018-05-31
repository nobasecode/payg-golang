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

    diff := finished.Sub(lastup)
    if diff.Seconds() > 0 { 
        return int(diff.Seconds())
    } else { return 0 }
}

//Difference bettwen Last-updated & Now Dates
func diff_date_now(lastup time.Time) int{

        diff := time.Since(lastup)
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



//add new containers to file
func add_container_file(containers [][]string){

        c_names := (file_to_c_name("watchlist"))
        //fmt.Println(c_names)

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

func pay_as_go(containers [][]string) {

    add_container_file(containers)
    
    c_file := file_to_c("watchlist")
    //fmt.Println(c_file)
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
                    new_time := last_value+diff_date_now(s_date(c_file[j][2]))
                    
                    //fmt.Println(diff_date_now(s_date(c_file[j][2])))
                    
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
                    new_time := last_value+diff_date(s_date(containers[i][3]),s_date(c_file[j][2]))
                    
                    //fmt.Println(diff_date(s_date(containers[i][3]),s_date(c_file[j][2])))
                    
                    t := time.Now()
                    a = a+"\033[31m"+containers[i][0]+" - "+strconv.Itoa(new_time)+" - "+t.Format("2006-01-02 15:04:05")+"\n"
                    
                    all = all+containers[i][0]+"|"+strconv.Itoa(new_time)+"|"+t.Format("2006-01-02 15:04:05")+"\n"

                }
            }
        }

    }

    err := ioutil.WriteFile("watchlist", []byte(all), 0666)
    if err != nil {log.Fatal(err)}

    
    show_status(a)
     
}

func main(){

    c := cron.New()
    c.AddFunc("@every 2s", func() { 
        pay_as_go(inspect())
        //fmt.Println(inspect())
     })
    c.Run()
    c.Start()


}

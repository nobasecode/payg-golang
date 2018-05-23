package main

import (
        "fmt"
        "log"
        "os/exec"
	"strings"
	"time"
	"os"
	"io/ioutil"
	//"strconv"
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
        return int(diff.Seconds())
}

//Difference bettwen Last-updated & Now Dates
func diff_date_now(lastup time.Time) int{

        diff := time.Since(lastup)
        return int(diff.Seconds())
}

//get a list of running docker containers
func inspect() [][]string {
	
	var containers [][]string
	var container []string
        
	out, err := exec.Command("sh","-c","docker inspect --format='{{.State.Running}}|{{.Name}}|{{.State.StartedAt}}|{{.State.FinishedAt}}|{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(docker ps -qa)").Output()  
	      if err != nil {
                    log.Fatal(err)
             }
 	
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
        fmt.Println(c_names)

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
        
	watchlist, err := os.Open("watchlist")
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

}

func main(){

        containers:=inspect()

        fmt.Println(containers)
	
	pay_as_go(containers)

}

package main

import (
        "fmt"
        //"log"
        //g "gosnmp"
        //toml "toml"
        //"io/ioutil"
        "strings"
        //"os"
        //"time"
       "os/exec"
)


func snmpget(user string,password string,ipadrr string) string {

        out, err := exec.Command("snmpget", "-u", user, "-l", "authPriv", "-a", "MD5", "-x", "DES", "-A", password, "-X", password, ipadrr, "1.3.6.1.2.1.1.3.0").Output()
                if err != nil {
                       // log.Fatal(err)
			return "dead"
              	} else {
		s := strings.Split(string(out), " ")
		return s[4]
		}
}


func main(){
	
	user:= "bootstrap"
	password:= "azertyui"
	ipadrr:= "172.23.237.78"
	result:=snmpget(user,password,ipadrr)

	fmt.Println(result)


}

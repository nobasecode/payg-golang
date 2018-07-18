
# PAY AS YOU GOlang

- Golang script for counting Docker containers running time & users credit with instant update.
- If you want to build a Containers as a Service (CaaS) platforme with Pay As You Go (PAYG) system, you will need this script.

**_Instructions_**

---------------------------------------
### **SQL Storage (MariaDB)**

To install this project:

```
$ git clone https://github.com/nobasecode/payg-golang.git
```

Import database from payg.sql

```
$ CREATE DATABASE newdatabase;
```
```
$ mysql -u [username] -p newdatabase < [database name].sql
```

You need to clean the follwing tables (credit & configuration) and add your own containers informations.

lunch script

```
$ go run main.go
```

The result will be like below (for every 2s):


![alt text](https://i.imgur.com/TnLaW31.png)

```diff
- You can change your execution time in cron function [c.AddFunc("@every 2s", func()]
```

---------------------------------------


### **Flat Files Storage**

To install this project:

```
$ git clone https://github.com/nobasecode/payg-golang.git
```

Clean information store files:

```
$ echo "" > watchlist
```
```
$ echo "" > credit
```
```
$ echo "" > conf
```
One you clean all files you will need to add your own informations:

- Watchlist will be auto-updated with your containers details once you start the script.

- Add every your containers configurations manually in "conf" file , following this schema:
`container name|RAM(mb)|CPU Number|Disk Size(gb)|User ID|`

- Add users credit in "credit" file, following this schema:
`User ID|Credit($)|`

lunch script

```
$ go run inspect.go
```

---------------------------------------


### **Using in Log**

you can fetch the using informations (`using time|using credit|CPU|Disk|Total use|`) in a given period of time by executing `use_log.go` script.

The result will be like below :

![alt text](https://i.imgur.com/pj5eJXw.png)

**Golang script to create Docker containers : https://github.com/nobasecode/ContainerGoCreator.git**

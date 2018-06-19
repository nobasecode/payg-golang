
### **PAY AS YOU GOlang**
- Golang script for counting Docker containers running time & users credit with instant update.
- If you want to build a Containers as a Service (CaaS) platforme with Pay As You Go (PAYG) system, you will need this script.


**_Instructions_**

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

**Soon : Golang script to create Docker containers**

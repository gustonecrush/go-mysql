# Go Mysql

## Agenda
We are going to learn about : 
1. Introduction Golang Database
2. Package Database
3. How to Create Connection Database
4. Execution SQL Command
5. SQL Injection
6. Prepare Statement
7. Database Transaction

## Introduction Golang Database

### Introduction Package Database
- Go-Lang has a package named database defaultly
- Package database is a package which contains a collection of standard interface which be the standard tool to communicate with database
- This makes our code program we created to access any type of database can use the code base
- The different is only in SQL Code which need to use same like the database we used

### How It Works
- `Aplikasi >>(call) Database Interface >>(call) Database Driver >>(call) DBMS`

### MySQL
- In this subject will be focus on MySQL as DBMS
 
## Golang Database Driver

### Database Driver
- Before we create code program which use database in Go-Lang, firstly, we need to add database driver
- Without database driver, database package in Go-Lang will not understand anything, because only contain contract interface
- You can visit <a href="https://golang.org/s/sqld rivers">for list of driver that you can use</a>
- There is so much module driver that you can use, the best practice to choose which one is the most used one

### Add Module Database MySQL
- `go get -u github.com/go-sql-driver/mysql`

### Import Package MySQL
- We can create our program, but firstly you have to import package MySQL
```
import (
  "database/mysql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "testing"
)
 ```
- Why we use _ because we only need to call init function, why init function? we only want to import package, without to use package manual, so automatically 

## Create Connection to Database

### Create Connection to Database
- The first things that we have to do before we create our application is open connection to database
- To open connection to database in Golang, we can create object sql.DB using function sql.Open(driver, dataSourceName)
- To use MySQL database, we can use driver "mysql"
- For dataSourceName, every database has their own writing to connect to database.
  - For MySQL, we can use dataSourceName like this
    `username:password@tcp(host:port)/database_name`
- If object sql.DB no more used, recommend to close the connection using Close() function to avoid connection league (condition where our application has been used but the connection is still open and running, which will make the database connection will always add that will make our database crashed because of there is so much connection used, because database has maximum connection)

### Code : Open Connection to Database
```
db, err := sql.Open("mysql", "user:password@tcp(host:3306)/dbname")
if err != nil {
  panic(err)
}
defer db.Close()
```

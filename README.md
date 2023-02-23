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
```go
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
```go
db, err := sql.Open("mysql", "user:password@tcp(host:3306)/dbname")
if err != nil {
  panic(err)
}
defer db.Close()
```

## Database Pooling

### Database Pooling
- sql.DB in Golang actually is not a connection to database
- But it is a pool to database, or known as Database Pooling concept
- In sql.DB, Golang do management connection to database automatically. This make us not have to manage our connection to database manually
- With this database pooling ability, we can determine minimal and maximal connection created by Golang. So, not make flood to our connection to database, because usually there is maximum connection which handled by database we used
- Why we need to declare minimum connection, because to avoid when suddenly up traffic that we don't know, so will make the connection faster if we determine the minimum as much as we expected
- Why we need to declare maximum connection, because to border the connection when there is bomb traffic, so they have to queue to use the connection

### Database Pooling Settings
- Method
  - (DB) SetMaxIdleConns(number) >> setting how minimal connection created
  - (DB) SetMaxOpenConns(number) >> setting how maximal connection created
  - (DB) SetConnMaxIdleTime(duration) >> setting how long unused connection will be removed
  - (DB) SetConnMaxLifetime(duration) >> setting how long connection might be used

### Database Pooling in Go-Lang Database
```go
func GetConnection() *sql.DB {
	db, err := sql.Open("myql", "root:root@tcp(localhost:8889)/db_golang")
	if err != nil {
	    panic(err)	
    }
	
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
}
```

## Execution SQL Command

### Execution SQL Command
- When we use database in our application, we must to communicate with database using SQL Command
- In Golang, there is functions that we can use to send SQL command to database using function `(DB) ExecContext(context, sql, params)`
- When sending SQL Command, we need to send context like we have learn before in course Golang Context, with context, we can send signal cancel if we cancel our SQL command sending

### Code: Create Table Customer
```sql
CREATE TABLE customer
(
    id VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;
```

### Code: Sending SQL Command Insert
```go
db := GetConnection()
defer db.Close()

ctx := context.Background()

_, err := db.ExecContext(ctx, "INSERT INTO customer(id, name) VALUES ('farhan', 'Farhan);")
if err != nil {
	panic(err)
}
fmt.Println("Success Insert Data to Database")
```

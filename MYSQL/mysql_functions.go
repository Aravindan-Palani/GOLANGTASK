package main

import (
        "fmt"
        "log"
        "database/sql"
        _"github.com/go-sql-driver/mysql"
)

func main() {
        db, err := sql.Open("mysql", "root:roadrunner@tcp(172.31.16.82:3306)/testDB")
        if err != nil {
                fmt.Println("Database Connected Fail")
        }
        fmt.Println("Database Connected Sucessfully")
        fmt.Printf("Enter a Option : ")
        var input int
        fmt.Scanln(&input)
        switch input {
        case 1:
                createdatabase(db)
        case 2:
                createtable(db)
        case 3:
                inserData(db)
        case 4:
                selectData(db)
        }
        defer db.Close()
}

func createdatabase(db *sql.DB) {
        res, err := db.Exec("CREATE DATABASE testDB;")
        //res, err := db.Exec("DROP DATABASE testDB;")
        if err != nil {
                panic(err)
        }

        numadded, err := res.RowsAffected()
        if err != nil {
                panic(err)
        }
        fmt.Print(numadded)

}

func createtable(db *sql.DB) {
        res, err := db.Exec("CREATE TABLE employee(id int NOT NULL AUTO_INCREMENT, first_name varchar(50), last_name varchar(30), PRIMARY KEY (id));")
        if err != nil {
                panic(err)
        }

        rowadded, err := res.RowsAffected()
        if err != nil {
                panic(err)
        }
        fmt.Print(rowadded)
}

func inserData(db *sql.DB) {
        sql := "INSERT INTO employee(first_name, last_name) VALUES ('jhon','wick')"

        res, err := db.Exec(sql)

        if err != nil {
                panic(err.Error())
        }

        lastId, err := res.LastInsertId()

        if err != nil {
                log.Fatal(err)
        }

        fmt.Printf("The last inserted row id: %d\n", lastId)
}


func selectData(db *sql.DB) {
        results, err := db.Query("Select * from employee;")
        if err != nil {
                panic(err.Error())
        }
        fmt.Println("Showing Tables", results)
        for results.Next() {
                var id int
                var first_name string
                var last_name string
                err = results.Scan(&id, &first_name, &last_name)
                if err != nil {
                        panic(err.Error())
                }
                fmt.Println(id)
                fmt.Println(first_name)
                fmt.Println(last_name)
        }
}


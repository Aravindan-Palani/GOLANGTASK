package main

import(
        "fmt"
        "github.com/patrickmn/go-cache"
        "time"
)

type Example struct {
        Name string
        Age int
}

func main() {
        newExample := &Example {
                Name: "jhon",
                Age: 12,
        }
        c := cache.New(5*time.Minute, 10*time.Minute)

        foo, found := c.Get("foo")
        if found {
                a := foo.([]Example)
                a = append(a, *newExample)
                c.Set("foo", a, cache.NoExpiration)
                fmt.Println("Inside If")
        } else {
                c.Set("foo", newExample, cache.NoExpiration)
                fmt.Println("Inside else")

        }

        foo1, found := c.Get("foo")
        if found {
                fmt.Println(foo1)
        }
}

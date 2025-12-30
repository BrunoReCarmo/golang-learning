package main

import "fmt"

func main() {
    var firstName string = "John"
    var lastName string = "Doe"

    var dayOfBirth int = 31
    var monthOfBirth int = 12
    var yearOfBirth int = 2001

    const userInfoFormat = "The user name is %s %s (Born date: %02d/%02d/%d)\n"
    //Console output will be ("The user name is John Dow (Born date: 31/12/2001)")
    fmt.Printf(userInfoFormat, firstName, lastName, dayOfBirth, monthOfBirth, yearOfBirth)
}
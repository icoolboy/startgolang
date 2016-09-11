package main

import (
    "fmt"
)

func main() {
    //Example 1 :
    /*
    x := 9
    switch x {
        case 10 :
            fmt.Println("x barabar ba 10 ast")
        case 9 :
            fmt.Println("x barabar ba 10 nist [Lazem nist] -> x barabar ba 9 ast [Tahodoodi lazeme]")
        default :
            fmt.Println(" x barabar ba 10 nist")
    }*/
    
    //Example 2 :
    /*
    language := "NodeJs"
    prefix := "my favorite language is "
    switch language {
        case "PHP" :
            fmt.Println(prefix + "PHP")
        case "C" :
            fmt.Println(prefix + "C")
        case "Go" :
            fmt.Println(prefix + "Go")
        default :
            fmt.Println(prefix + language)
    }*/
    
    //Example 3 :
    /*
    age, result := 7, ""
    switch {
        case age >= 20 && age <= 35 :
            result = "javan"
        case age >= 10 && age <= 20 :
            result = "nojavan"
        case age >= 5 && age < 10 :
            result = "koodak"
        case age >= 1 && age < 5 :
            result = "khordsal"
        default :
            result = "Az rade seni mored nazar ma kharej ast"
    }
    fmt.Println(result)
    */
    
    //Example 4 :
    /*
    city := "Gorgan"
    switch  {
        case city == "Yazd" || city == "Esfahan" :
            fmt.Println("garm va khoshk")
        case city == "Gorgan" || city == "Lavij" :
            fmt.Println("Martoob")
        default :
            fmt.Println("gheyre")
    }*/
    
    //Example 5 :
    /*
    name := "ali"
    switch {
        case name != "erfan" :
            fmt.Println("name shoma erfan nist")
        default :
            fmt.Println("name shoma erfan ast")
    }*/
    var age = 12
    switch {
        case age >= 10 && age <= 20 :
            switch {
                case age == 11 :
                    fmt.Println("11")
                case age != 11 :
                    fmt.Println("no")
            }
        default :
            fmt.Println("No")
    }
}

//Switch
/*
syntax of switch's base : 
switch [expression] -> no need {
    case [n1] :
        //op when n1 is true
    case [n2] :
        //op when n2 is true
    default :
        //op when anyone of top case aren't true
}

Logical Examples :
------------------
1. x = 9
    A : case 10 -> result : x barabar ba 10 ast
    B : case 9 -> result : x barabar ba 10 nist [Lazem nist] -> x barabar ba 9 ast [Tahodoodi lazeme]
    C : default -> result : x barabar ba 10 nist

2. language = "Go"
    A : case "PHP" -> result : my favorite language is php
    B : case "C" -> result : my favorite language is c
    C : case "Go" -> result : my favorite language is Go
    D : default -> result : my favorite language is [language]
    
3. age = 20
    A : case age >= 20 && age <= 35 -> result : javan
    B : case age >= 10 && age <= 20 -> result : nojavan
    C : case age >= 5 && age < 10 -> result : koodak
    D : case age >= 1 && age < 5 -> result : khordsal
    E : default : Az rade seni mored nazar ma kharej ast
    
4. city = "{Tehran | Mashhad | Lavij | Gorgan | Yazd | Esfahan | ...}"
    A : case "Yazd" || city == "Esfahan" -> result : garm va khoshk
    B : case "Gorgan" || city == "Lavij" -> result : martoob
    C : default -> result : gheyre
    
5. name = "Erfan"
    A : case name != "Erfan" -> result : name shoma erfan nist
    B : default -> result : name shoma erfan ast
    
6. nested, break & fallthrough and ...

Sample (nested switch):
    var age = 11
    switch {
        case age >= 10 && age <= 20 :
            switch {
                case age == 11 :
                    fmt.Println("11")
                case age != 11 :
                    fmt.Println("no")
            }
        default :
            fmt.Println("No")
    }

*/

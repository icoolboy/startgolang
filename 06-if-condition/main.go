package main

import (
    "fmt"    
)

func main() {
    /*
    Example 1 :
    var x = 5
    if x == 10 {
        fmt.Println("x barabar ba 10 ast")
    } else if x == 9 {
        fmt.Println("x barabar ba 9 ast")
    } else {
        fmt.Println("x barabar ba 10 nist")
    }
    
    Example 2 :
    var language = "php"
    const TEXT = "my favorite language is "
    if language == "PHP" || language == "php" {
        fmt.Println(TEXT + "Php")
    } else if language == "C" {
        fmt.Println(TEXT + "c")
    } else if language == "GO" {
        fmt.Println(TEXT + "Go")
    } else {
        fmt.Println("Else -> " + TEXT + language)
    }
    
    Example 3 :
    var age, result = 55, ""
    
    if age >= 20 && age <=35 {
        result = "javan"
    } else if age >= 10 && age < 20 {
        result = "Nojavan"
    } else if age >= 5 && age < 10 {
        result = "Koodak"
    } else if age >= 1 && age < 5 {
        result = "Khordsal"
    } else {
        result = "Az rade seni mored nazar ma kharej ast"
    }
    
    fmt.Println(result)
    
    Example 4 :
    var city = "Gorgan"
    if city == "Yazd" || city == "Esfahan" {
        fmt.Println("Garm va khoshk")
    } else if city == "Gorgan" || city == "Lavij" {
        fmt.Println("Martoob")
    } else {
        fmt.Println("Gheyre")
    }
    
    Example 5 :
    */
    var name = "Ali"
    if name != "Erfan" {
        fmt.Println("Name shoma erfan nist")
    } else {
        fmt.Println("Name shoma erfan ast")
    }
    
}
/*
if command syntax :
-------------------
if condition {
    //if condition == true this code will run
} else if condition {
    //if second condition == true this code will run
} else {
    //And if any conditions != true this code will run
}

Live Examples :
----------
1. Agar nakham aval sobh be madrese beram !
        A : Ba ordangi miram
        B : Bazam ba ordangi miram
        C : Film bazi mikonam ke mariz shodam nemiram :D
        D : Hichkodam
        
2. Agar man bekham ezdevaj konam chand halat dare !
        A : Be khak siyah mishinam
        B : Khoshbakht misham
        C : Rooz bade ezdevaj talagh migiram :D
        D : Hichkodam
        
3. Agar dars hay daneshgah ro nakhoonam chi mishe ?
        A : Hichi mashroot misham be hamin rahati
        B : Daneshgah kiloo chande?
        C : Enseraf midam madrak che mikham , vallaaa !
        D : Akhar term ke shod rajebesh fekr mikonam
        E : Hichkodam

Logical Examples :
------------------
1. x = 9
    A : if x == 10 -> result : x barabar ba 10 ast
    B : else if x == 9 -> result : x barabar ba 10 nist [Lazem nist] -> x barabar ba 9 ast [Tahodoodi lazeme]
    C : else -> result : x barabar ba 10 nist

2. language = "Go"
    A : if language == "PHP" -> result : my favorite language is php
    B : else if language == "C" -> result : my favorite language is c
    C : else if language == "Go" -> result : my favorite language is Go
    D : else -> result : my favorite language is [language]
    
3. age = 20
    A : if age >= 20 && age <= 35 -> result : javan
    B : if age >= 10 && age <= 20 -> result : nojavan
    C : if age >= 5 && age < 10 -> result : koodak
    D : if age >= 1 && age < 5 -> result : khordsal
    E : Az rade seni mored nazar ma kharej ast
    
4. city = "{Tehran | Mashhad | Lavij | Gorgan | Yazd | Esfahan | ...}"
    A : if city == "Yazd" || city == "Esfahan" -> result : garm va khoshk
    B : if city == "Gorgan" || city == "Lavij" -> result : martoob
    C : else -> result : gheyre
    
5. name = "Erfan"
    A : if name != "Erfan" -> result : name shoma erfan nist
    B : else -> result : name shoma erfan ast
*/
package main

import (
    "fmt"
    _"time"
)

//Here for "var"
//Shortcut -> error

func main() {
    //Here for "var"
    //Shortcut -> Ok


    /*
    var x bool //Default values : string -> "" | number -> 0 | bool -> false | Static type
    var x = 4.5 //Dynamic type
    x := 10 //Dynamic type
    */
    /*
    Method 1 :
    var x string
    var y int32
    var z float32
    var b bool
    
    Method 2 :
    var x, y int16
    var z, b, c string
    
    Method 3 :
    var x string = "Erfan"
    var y bool = true
    var z int32 = 2000000
    var b float64 = 3.14000000999
    var c string = "Json"

    Method 4 :
    var x, y, z, b, c = 10, "Go", true, "Hi", 4.5

    Method 5 :
    var x, y, z, b, c string = "Go", "Hi", "React", "Js", "PHP"
    
    Method 6 :
    var (
        x = 10
        y string = "Hi"
        z string
        b bool = true
        c = 4.5
    )
    
    Method 7 :
    x := 10
    y := "hi"
    z := 4.3
    b := true
    c := false
    
    Method 8 :
    x, y, z, b, c := 10, "Go", "C", 8.9, false
    
    Method 9 :
    hello := func() {
        fmt.Println("Hello")
    }
    
    hello()
    
    Method 10 :
    _, x := 10, 30
    
    Method 11 :
    var x = 10
    var y = true
    var z = 7.5
    var b = "my data"
    var c = false
    */
    var (
        x = 10
        y string = "Hi"
        z string
        b bool = true
        c = 4.5
    )
    println(x, y, z, b, c) // memory address , find value
    fmt.Printf("%T %T %T %T %T", x, y, z, b, c)
    //time.Sleep(10 * time.Second)
}

//Here for "var"
//Shortcut -> error

/*

Memory {
    [0xf00 -> 10][][][][][][][][][][][][][][][][0xfff00 [x]-> false][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
    [][][][][][][][][][0xff09-> Erfan][][][][][][][][][][][][][][][][][][][][][][][][][0x0000146 -> 4.5569][][][][][][][][][][][][][][][][][][][][][][][]
    [0xffc10 -> daneshjooyar@gmail.com][][][][][][][][][][][][][][][][37D00F -> -20][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
}
 
A lot of models for define variables :
------------------------------------
1. var varName type
2. var varName1, varName2 type
3. var varName type = value
4. var varName1, varName2, varName3 = value1, value2, value3
5. var varName1, varName2, varName3 type = value1, value2, value3
6. var (
        varName type
        varName type = value
        varName1 = value1
   )
7. varName := value
8. varName1, varName2 := value1, value2
9. varName := func() {
        ... 
   }
10. _, varName := 10, 20
11. var varName = value
12. arrays, slices, etc ...
*/
package main

import (
    "fmt"
)

func main() {
    //fmt.Println("Hi Go 813 @# $+ ", 88, 4.5)
    fmt.Printf("%d %s %f\r\n", 10, "Erfan", 4.3)
    fmt.Printf("%T %T\r\n", 70, 3.5)
    fmt.Printf("Erfan type -> %T", "Erfan")
}

/*
 Sahih : -78 -46 -100 -7000 0 6 1000 100000 78000 78000000 -> A
 Hesabi : 0 67 90 80000  -> B 
 Ashari : 2.5 7.3  9.7  8.90 8.0
 Reshte ya matn : Erfan Go Hi Have 147
 
 Format :
 %d -> Sahih -> int
 %s -> Matn ya reshte -> string
 %T -> Moshakhas kardan noe dade voroodi
 %f -> Adad ashari -> float32 float64 
*/

//Space -> 88 -> *Hesabi & Sahih
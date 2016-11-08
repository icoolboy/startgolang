package main

import (
    "fmt"
)

func main() {
    var m, n, b [7]int
    fmt.Println(m, n, b)
    //var x, y, z, j, c, t, g, f int = 10, 20, 30, 40, 50, 60, 80, 800
    var x [8]int //method 1
    x[0] = 10
    x[1] = 20
    x[2] = 30
    x[3] = 40
    x[4] = 50
    x[5] = 60
    x[6] = 70
    x[7] = 80
    //x[8] = 90 ->  invalid array index 8 (out of bounds for 8-element array)
    fmt.Println(x, len(x))
    
    var y [4]string
    y = [4]string{"Hi", "Hello", "Bye", "GoodBye"}
    fmt.Println(y, len(y), y[3])
    
    z := [2]bool{false, true}
    fmt.Println(z, len(z), z[1])
    
    language := [5]string{
        "Go",
        "PHP",
        "C#",
        "C++",
        "NodeJs",
    }
    
    fmt.Println(language)
    
    var i, j int
    
    for i = 0; i < len(x); i++ {
        fmt.Println(x[i])
    }
    
    var sum = 0
    for i = 0; i < len(x); i++ {
        sum += x[i] //sum = sum + x[i]
    }
    
    fmt.Println("Total of data : ", sum)
    
    var t [2][4]int
    for i = 0; i < 2; i++ {
        for j = 0; j < 4; j++ {
            t[i][j] = i + 10 + j - 5
        }
    }
    
    fmt.Println(t)
    
    xyz := [5]float32{10.5, 3.6, 5.8, 9.2, 4.1}
    printArray(xyz)
    
}

func printArray(p [5]float32) {
    for i := 0; i < len(p); i++ {
        fmt.Println("Your array is : \r\n", p[i])
    }
}
/*
Arrays :
It's similar to variables !!! but we have one for all
Wow that's good !!!!

Now see :
    Variable => (x = 10)
        OR
    Array => (x = [arrayLen]arrayType{10, 20, 30, 8, 63, 92, 75, 2} || x[0] = 1 , ...)

Basic Syntax :
1. var arrayName [arrayLen]arrayType => var x [8]int
2. arrayName := [arrayLen]arrayType{} => x := [8]int{10, 20, 30, 8, 63, 92, 75, 2}
3. var arrayName [arrayLenOne][arrayLenTwo]arrayType

Method 1 :
10, 20, 3, 671, 9103, 7759, 135, 31, 72
print array's data

Method 2 :
Go, PHP, Ruby, C++, C#, NodeJs, Matlab
print array's index (optional)

Method 3 :
Musics, Videos, Documents, Projects, Images
Change data of index 4 to Pictures

Method 4 :
array & loop

Yesterday, Today, Tomarrow

func myFunc(arr [4]string) {
    //... ???
    //print ???
}
make function which have one input argument, one array pass to function
with for loop try to print array's data

*/
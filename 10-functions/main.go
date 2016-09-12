package main

import (
    "fmt"        
)

//Method 1 :
func hello() {
    println("Hello world")
}

func main() {
    //hi()
    hello()
    welcome("Ali")
    welcomeUser("Erfan", "Akbarimanesh", 21)
    s := welcomeMember("Ali", "Alavi", 22)
    println(s)
    
    nam, sen, msg := welcomeAdmin("Erfan", "Akbarimanesh", 21)
    println(nam)
    println(sen)
    println(msg)
    
    salam := func(name string) string {
        return "Welcome " + name
    }
    
    println(salam("Erfan"))
}

//Method 2 :
func welcome(name string) {
    println("welcome " + name)
}

//Method 3 :
func welcomeUser(name, family string, age int) {
    //println("welcome " + name + " " + family + "Age :", age)
    s := fmt.Sprintf("welcome %s %s Age : %d", name, family, age)
    println(s)
}

//Method 4 :
func welcomeMember(name, family string, age int) string {
    return fmt.Sprintf("welcome %s %s Age : %d", name, family, age)
}

//Method 5 :
func welcomeAdmin(name, family string, age int) (string, int, string) {
    return name + " " + family, age, "Myname is warning"
}

/*
function syntax :
func functionName(inputs type) [returns] {
    //code here
}

1 : func hi()
2 : func hi(with input)
3 : func hi(with multi inputs)
4 : func hi(with input) with return
5 : func hi(with input) with multi return
6 : x := func hi(...) ... ==> init func to variable -> Closure functions

7 : func add(args ...int) int {} -> Variadic functions ...
8 : callback function(Recursion) and ...
*/
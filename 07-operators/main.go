package main

import (
    "fmt"
)

func main() {
    var ( 
        x, y float32 = 2, 3 //2.00000 , 3.00000
        z, j int = 10, 3
        counter = 10
    )
    //Arithmetic operators : 
    //x -> 2.00000 + 1.00000 = 3.00000
    //x++
    counter++ //11
    fmt.Println(x, y)
    fmt.Println("x + y :", x + y) //x + y : 5
    fmt.Printf("x - y : %f\r\n", x - y) // x - y : -1
    fmt.Printf("x * y : %f\r\n", x * y)
    fmt.Printf("x / y : %f\r\n", x / y)
    fmt.Println("z % j : ", z % j)
    counter-- //10
    fmt.Println(counter)
    
    //Relational operators examples -> go to last lesson
    
    //Logical operators
    //op : && || !
    /*
    false -> 0
    true -> non-zero 
    AND : 
    true && true -> true
    true && false -> false
    false && false -> false
    
    OR :
    true || true -> true
    true || false -> true
    false || false -> false
    
    NOT :
    !true -> false
    !false -> true
    */
    i, u := true, false
    fmt.Println(!(i && u))
    fmt.Println(!(i || u))
    
    k, r := 10, 2
    k+=r //k = k + r
    k-=r // k = k - r
    k*=r // k = k * r
    k/=r //k = k / r
    k%=r //k = k % r
    fmt.Println(k, r)
}
/*
All operators models in Go :
----------------------------
1. Arithmetic Operators
    +   =>   x + y | 2 + 2
    -   =>   x - y | 2 - 2
    *   =>   x * y | 2 * 2
    /   =>   x / y | 2 / 2  : kharej ghesmat taghsim x bar y    
    %   =>   x % y | 2 % 2 : baghimande taghsim x bar y 
    ++  =>   x = 1  x++ 
    --  =>   x = 1  x-- 
    
2. Relational Operators
    ==  ==>   x == y
    !=  ==>   x != y
    >   ==>   x > y
    <   ==>   x < y
    >=  ==>   x >= y
    <=  ==>   x <= y

3. Logical Operators
    Init :
    [x = true, y = false]
    Notice : [0 is false, non-zero is true]
    
    &&  =>  x && y  -> agar x va y har 2 gheyr sefr bashand natije true ast
            if x && y {
                //true
            } else {
                //false
            }
    ||  =>  x || y  -> agar yeki az moteghayer ha gheyr sefr bashad natije true ast
            if x || y {//true} else {//false}
    !   =>  !(x||y) -> naghz va makoos kardan
            if !(x || y) {//tayin bar asas natije} else {//tayin bar asas natije}

4. Bitwise Operators
    bit-by-bit
    Now no need, after speak about this operators
    
5. Assignment Operators
    =    -> z = x + y
    +=   -> x+=y  => x = x + y
    -=   -> x-=y  => x = x - y
    *=   -> x*=y  => x = x * y
    /=   -> x/=y  => x = x / y 
    %=   -> x%=y  => x = x % y
    <<=  ->  future
    >>=  ->  future
    &=   ->  future
    |=   ->  future
    ^=   ->  future
    
6. Misc Operators
    &  ->  return the memory address of variable
    *  ->  pointer to a value or data
    In the future we speak about this topic
    
7. Etc for e.g{[], (), ., ->, ?:, ...}
*/
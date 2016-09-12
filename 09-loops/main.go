package main

func main() {
    /*
    println("salam")
    println("salam")
    println("salam")
    println("salam")
    println("salam")
    println("salam")
    println("salam")
    println("salam")
    println("salam")
    println("salam")
    */
    for i := 1; i <= 10; i++ {
        println("salam")
    }
    
    /*
    println(1)
    println(2)
    println(3)
    println(4)
    println(5)
    println(6)
    println(7)
    println(8)
    println(9)
    println(10)
    */
    for x := 1; x <= 10; x++ {
        println(x)
    }
    
    /*
    println(10)
    println(9)
    println(8)
    println(7)
    println(6)
    println(5)
    println(4)
    println(3)
    println(2)
    println(1)
    */
    for i := 10; i >= 1; i-- {
        println(i)
    }
    
    /*
    println(-10)
    println(-9)
    println(-8)
    println(-7)
    println(-6)
    println(-5)
    println(-4)
    println(-3)
    println(-2)
    println(-1)
    */
    for j := -10; j <= -1; j++ {
        println(j)
    }
    
    /*
    println(-1)
    println(-2)
    println(-3)
    println(-4)
    println(-5)
    println(-6)
    println(-7)
    println(-8)
    println(-9)
    println(-10)
    */
    for y := -1; y >= -10; y-- {
        println(y)
    }
    
    //infinite loop 
    /*
    for {
        println("Hi Go")
    }
    */
    
    p := 2
    for p <= 10 {
        //p++
        println(p)
        p++
    }
    
    //sum of 0 to 10 :
    var sum int = 0
    for i := 0; i <= 10; i++ {
        sum += i //sum = sum + i
    }
    
    println("sum of 0 to 10 : ", sum)
    
    //even and odd in 0 to 10 :
    //2 % 2 -> if result 0 -> even
    //3 % 2 -> result -> non 0 odd
    for i := 0; i <= 10; i++ {
        if i % 2 == 0 {
            println("Even number : ", i)
        } else {
            println("Odd number : ", i)
        }
    }
    
}

/*
For Syntax :
for [first] -> init; [second] -> condition; [third] -> step {
    //code
}

Examples :
1. salam salam salam salam salam :
for i := 1; i <= 5; i++ {
    //code
}

2. 1  2  3  4  5  6  7  8  9  10 :
for i := 1; i <= 10; i++ {
    //code
}

3. 10  9  8  7  6  5  4  3  2  1 :
for i := 10; i <= 1; i-- {
    //code
}

4. -10 -9 -8 -7 -6 -5 -4 -3 -2 -1 :
for i := -10; i <= -1; i++ {
    //code
}

5. -1 -2 -3 -4 -5 -6 -7 -8 -9 -10 :
for i := -1; i >= -10; i-- {
    //code
}

6.Infinite loop : 
for {
  //code 
}

for while :
i := 1
for i <= 10 {
    i++
}

Practice : 
Sample : sum of 0 to 10, even and odd in 0 to 10, ...

*/
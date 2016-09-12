package main

func main() {
    /*var PI = 3.14
    println(PI)
    PI = 3.15
    println(PI)*/
    /*
    Method 1 :
        const PI = 3.14
    Method 2 :
        const VERSION string = "1.0.0"
    Method 3 :
        const APPNAME, APPVERSION, APPDEVELOPER = "MyChat", "1.0.1", "Erfan"
        
    Method 4 :
        const APPNAME, APPVERSION, APPDEVELOPER string = "MyChat", "1.0.1", "Erfan"
    Method 5 :
        const (
            APPTITLE = "Telegram"
            APPCOMPANY string = "telegram.me - telegram co"
            APPVERSION int = 1
            APPENABLE bool = true
        )
        
    */
    const (
        APPTITLE = "Telegram"
        APPCOMPANY string = "telegram.me - telegram co"
        APPVERSION int = 1
        APPENABLE bool = true
    )
    
    println(APPTITLE, APPCOMPANY, APPVERSION, APPENABLE)
    //PI = 3.15 //Error
}

/*

Memory {
    [0xf00 -> 10][][][][][][][][][][][][][][][][0xfff00 [x]-> false][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
    [][][][][][][][][][0xff09-> Erfan][][][][][][][][][][][][][][][][][][][][][][][][][0x0000146 -> 4.5569][][][][][][][][][][][][][][][][][][][][][][][]
    [0xffc10 -> daneshjooyar@gmail.com][][][][][][][][][][][][][][][][37D00F -> -20][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
}
 
A lot of models for define constants :
------------------------------------
1. const constName = value
2. const constName type = value
3. const constName1, constName2, constName3 = value1, value2, value3
4. const constName1, constName2, constName3 type = value1, value2, value3
5. const (
        constName type = value
        constName1 = value1
   )

*/
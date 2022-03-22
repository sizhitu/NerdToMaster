package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var (
        n int64
        t int
        )
    fmt.Scan(&t)
    for i := 0; i < t; i++{
        fmt.Scan(&n)
        judgePrime(n)
    }
}

func judgePrime(num int64){
    var i int64
    if num < 2 {
        fmt.Println("Not prime")
        return
    }
    for i = 2; i * i <= num; i++{
        if num % i == 0{
            fmt.Println("Not prime")
            return
        }
    }
    fmt.Println("Prime")
}

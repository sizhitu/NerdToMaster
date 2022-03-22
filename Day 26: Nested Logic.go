package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var actualDay,actualMonth,actualYear,expectDay,expectMonth,expectYear,fine int
    fmt.Scan(&actualDay,&actualMonth,&actualYear)
    fmt.Scan(&expectDay,&expectMonth,&expectYear)
    if actualYear > expectYear{
        fine = 10000
    }else if actualYear == expectYear{
        if actualMonth <= expectMonth{
            if actualDay > expectDay{
                fine = 15 * (actualDay - expectDay)
            }
        }else{
            fine = 500 * (actualMonth - expectMonth)
        }
    }else {
        fine = 0
    }
    fmt.Println(fine)
}


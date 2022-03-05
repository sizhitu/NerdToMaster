/*******************************************************************************/
Day 0: Hello, World
package main
import ("fmt"
        "bufio"
        "os"
        )

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan(){
        fmt.Println("Hello, World.\n"+scanner.Text())
    }
}

/********************************************************************************/
Day 1: Data Types
package main


import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
  
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)
    // Declare second integer, double, and String variables.
    
    var (
        secondInt uint32
		    secondDub float32
		    secondStr string
    )
    // Read and save an integer, double, and String to your variables.
    fmt.Scan(&secondInt)
    fmt.Scan(&secondDub)
    if scanner.Scan() {
        secondStr = scanner.Text()
    }
    // Print the sum of both integer variables on a new line.
    sum := secondInt + i
    fmt.Println(sum)
    // Print the sum of the double variables on a new line.
    sumDub := secondDub + d
    fmt.Println(fmt.Sprintf("%.1f",sumDub))
    // Concatenate and print the String variables on a new line
    // The 's' variable above should be printed first.
    fmt.Println(s+secondStr)
}

/*****************************************************************************/
Day 2: Operators
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'solve' function below.
 *
 * The function accepts following parameters:
 *  1. DOUBLE meal_cost
 *  2. INTEGER tip_percent
 *  3. INTEGER tax_percent
 */

func solve(meal_cost float64, tip_percent int32, tax_percent int32) {
    // Write your code here
    perct := tip_percent+tax_percent
    val := 1+float64(perct)*1e-2      //保留小数点后2位
    res := int32((meal_cost*val)+0.5) //对结果四舍五入
    fmt.Println(res)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    meal_cost, err := strconv.ParseFloat(strings.TrimSpace(readLine(reader)), 64)
    checkError(err)

    tip_percentTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    tip_percent := int32(tip_percentTemp)

    tax_percentTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    tax_percent := int32(tax_percentTemp)

    solve(meal_cost, tip_percent, tax_percent)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

/***********************************************************************************/
Day 3: Intro to Conditional Statements
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    N := int32(NTemp)
    switch {
        case N%2!=0:
        fmt.Println("Weird")
        case N%2==0:
            switch{
                case N >2 && N <= 5:
                    fmt.Println("Not Weird")
                case N >= 6 && N <= 20:
                    fmt.Println("Weird")
                case N > 20:
                    fmt.Println("Not Weird")
            }
    }
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

/**********************************************************************************/
Day 4: Class vs. Instance
package main
import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	//Add some more code to run some checks on initialAge
    if initialAge < 0{
        p.age = 0
        fmt.Println("Age is not valid, setting age to 0.")
    }else{
        p.age = initialAge
    }
	return p
}

func (p person) amIOld() {
	//Do some computation in here and print out the correct statement to the console
    if p.age < 13{
        fmt.Println("You are young.")
    }else if p.age < 18 && p.age >= 13{
        fmt.Println("You are a teenager.")
    }else if p.age >= 18{
        fmt.Println("You are old.")
    }
}

func (p person) yearPasses() person {
	//Increment the age of the person in here
    p.age += 1
	return p
}

func main() {
    var T,age int

    fmt.Scan(&T)

    for i := 0; i < T; i++ {
        fmt.Scan(&age)
        p := person{age: age}
        p = p.NewPerson(age)
        p.amIOld()
        for j := 0; j < 3; j++ {
            p = p.yearPasses()
        }
        p.amIOld()
        fmt.Println()
    }
}

/****************************************************************************************/
Day 5: Loops
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)
    for i:=1;i<=10;i++{
        res := n*int32(i)
        fmt.Printf("%v x %v = %v\n",n,i,res)
    }
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

/***********************************************************************/
Day 6: Let's Review
package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var (T int
         s string
         res string
         s1 string
         s2 string)
    
    fmt.Scan(&T)
    for i := 0;i < T;i++{
        fmt.Scan(&s)
        for k,v := range s{
            if k % 2 == 0{
                s1 = s1 + string(v)
            }else{
                s2 = s2 + string(v)
            }
        }
        res = s1 +" "+ s2
        fmt.Println(res)
        s1 = ""
        s2 = ""
    }
    
}
/*******************************************************************************/
Day 7: Arrays
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32
    var res string

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }
    for i:=len(arr)-1;i>=0;i-- {
        res = res + strconv.Itoa(int(arr[i]))+" "
    }
    fmt.Println(res)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
/**********************************************************************************************/
Day 8: Dictionaries and Maps
package main
import ("fmt"
        "bufio"
        "io"
        "os"
        "strings"
        "strconv")

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    reader := bufio.NewReaderSize(os.Stdin,1024*1024)
    nTemp , err := strconv.ParseInt(readLine(reader), 10, 64)
    n := int(nTemp)
    checkError(err)
    res := make(map[string]string)
    for i := 0; i < n; i++{
        arr := strings.Split(readLine(reader)," ")
        res[arr[0]] = arr[1]
    }
    for{
        name := readLine(reader)
        if name == ""{
            break
        }
        if val,ok := res[name];ok{
            fmt.Printf("%s=%s\n",name,val)
        }else{
            fmt.Println("Not found")
        }
    }
}

func readLine(reader *bufio.Reader) string{
    str, _, err := reader.ReadLine()
    if err == io.EOF{
        return ""
    }
    return strings.TrimRight(string(str),"\r\n")
}

func checkError(err error){
    if err != nil{
        panic(err)
    }
}
/****************************************************************************************/
Day 9: Recursion 3
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'factorial' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

func factorial(n int32) int32 {
    // Write your code here
    if n <= 1{
        return 1
    }
    return factorial(n-1)*n
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    result := factorial(n)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}


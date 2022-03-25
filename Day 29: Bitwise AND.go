package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

/*
 * Complete the 'bitwiseAnd' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER N
 *  2. INTEGER K
 */

func bitwiseAnd(N int32, K int32) int32 {
    // Write your code here
    var res,max,i,j int32
    max = math.MinInt32
    for i = 1; i < N; i++{
        for j = i + 1; j <= N && i & j < K; j++{
            res = i & j
            if res > max {
                max = res
            } 
        } 
    }
    return max
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        countTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
        checkError(err)
        count := int32(countTemp)

        limTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
        checkError(err)
        lim := int32(limTemp)

        res := bitwiseAnd(count, lim)

        fmt.Fprintf(writer, "%d\n", res)
    }

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

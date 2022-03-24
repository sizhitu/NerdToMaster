package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
    "regexp"
)



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    N := int32(NTemp)
    keys := make([]string, N)

    for NItr := 0; NItr < int(N); NItr++ {
        firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        firstName := firstMultipleInput[0]

        emailID := firstMultipleInput[1]
        
        re := regexp.MustCompile(string(`([a-zA-Z0-9]+)@gmail\.com`))
        
        match := re.FindString(emailID)
        if match != "" {
            keys = append(keys,firstName)
        }
    }
    
    sort.Strings(keys)
    for _, v := range keys {
        if v != ""{
            fmt.Println(v)
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

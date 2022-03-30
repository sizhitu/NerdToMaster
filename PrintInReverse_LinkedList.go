package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

type SinglyLinkedListNode struct {
    data int32
    next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
    head *SinglyLinkedListNode
    tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
    node := &SinglyLinkedListNode {
        next: nil,
        data: nodeData,
    }

    if singlyLinkedList.head == nil {
        singlyLinkedList.head = node
    } else {
        singlyLinkedList.tail.next = node
    }

    singlyLinkedList.tail = node
}

func printSinglyLinkedList(node *SinglyLinkedListNode, sep string) {
    for node != nil {
        fmt.Printf("%d", node.data)

        node = node.next

        if node != nil {
            fmt.Printf(sep)
        }
    }
}

/*
 * Complete the 'reversePrint' function below.
 *
 * The function accepts INTEGER_SINGLY_LINKED_LIST llist as parameter.
 */

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */

func reversePrint(llist *SinglyLinkedListNode) {
    // Write your code here
    var prev *SinglyLinkedListNode

    for llist != nil{
        cur := llist.next
        llist.next = prev
        //move prev
        prev = llist
        //move head
        llist = cur
    }
    for prev != nil{
        fmt.Println(prev.data)
        prev = prev.next
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    testsTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    tests := int32(testsTemp)

    for testsItr := 0; testsItr < int(tests); testsItr++ {
        llistCount, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)

        llist := SinglyLinkedList{}
        for i := 0; i < int(llistCount); i++ {
            llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
            checkError(err)
            llistItem := int32(llistItemTemp)
            llist.insertNodeIntoSinglyLinkedList(llistItem)
        }

        reversePrint(llist.head)
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

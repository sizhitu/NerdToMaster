package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

type DoublyLinkedListNode struct {
    data int32
    next *DoublyLinkedListNode
    prev *DoublyLinkedListNode
}

type DoublyLinkedList struct {
    head *DoublyLinkedListNode
    tail *DoublyLinkedListNode
}

func (doublyLinkedList *DoublyLinkedList) insertNodeIntoDoublyLinkedList(nodeData int32) {
    node := &DoublyLinkedListNode {
        next: nil,
        prev: nil,
        data: nodeData,
    }

    if doublyLinkedList.head == nil {
        doublyLinkedList.head = node
    } else {
        doublyLinkedList.tail.next = node
        node.prev = doublyLinkedList.tail
    }

    doublyLinkedList.tail = node
}

func printDoublyLinkedList(node *DoublyLinkedListNode, sep string, writer *bufio.Writer) {
    for node != nil {
        fmt.Fprintf(writer, "%d", node.data)

        node = node.next

        if node != nil {
            fmt.Fprintf(writer, sep)
        }
    }
}

/*
 * Complete the 'sortedInsert' function below.
 *
 * The function is expected to return an INTEGER_DOUBLY_LINKED_LIST.
 * The function accepts following parameters:
 *  1. INTEGER_DOUBLY_LINKED_LIST llist
 *  2. INTEGER data
 */

/*
 * For your reference:
 *
 * DoublyLinkedListNode {
 *     data int32
 *     next *DoublyLinkedListNode
 *     prev *DoublyLinkedListNode
 * }
 *
 */

func sortedInsert(llist *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {
    // Write your code here
    head := llist
    
    cur := &DoublyLinkedListNode{}
    cur.data = data
    //head
    if llist.data >= data{
        cur.next = llist
        llist.prev = cur        
        return cur
    } 
    for llist != nil {
        //mid
        if llist.data < data && llist.next != nil && llist.next.data >= data{
            tmp := llist.next
            llist.next = cur
            cur.prev = llist
            cur.next = tmp
        } 
        //tail 
        if llist.data < data && llist.next == nil {
                llist.next = cur
                cur.prev = llist
                cur.next = nil
        }
        llist = llist.next
    }
    return head
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        llistCount, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)

        llist := DoublyLinkedList{}
        for i := 0; i < int(llistCount); i++ {
            llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
            checkError(err)
            llistItem := int32(llistItemTemp)
            llist.insertNodeIntoDoublyLinkedList(llistItem)
        }

        dataTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        data := int32(dataTemp)

        llist1 := sortedInsert(llist.head, data)

        printDoublyLinkedList(llist1, " ", writer)
        fmt.Fprintf(writer, "\n")
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

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

/******************************************************************************************/
Day 10: Binary Numbers
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
    n := int64(nTemp)
    baseTwo := string(strconv.FormatInt(n,2))
    
    var local, max int
    for i := range baseTwo{
        if baseTwo[i] == '1'{
            local++
            if local > max{
                max = local
            }
        }else{
            local = 0
        }
    }
    fmt.Println(max)
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

/*****************************************************************************/
Day 11: 2D Arrays

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

    var arr [][]int32
    for i := 0; i < 6; i++ {
        arrRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var arrRow []int32
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != 6 {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }
    fmt.Print(findMaxHourGlasses(arr))
}

func findMaxHourGlasses(arr [][]int32) int32{
    var max int32
    max = -2147483648 //zero value will be the max of negetives
    for x := 1; x < len(arr)-1; x++{  //outer loop denotes column num, due to hourglass structure, the column only can move len(arr)-2 loops
        for y := 1; y < len(arr[x])-1; y++{  //inner loop denotes row num
            var total int32  //tmp value used for sum must in the inner loop to compare with max
            //up
            up := arr[y-1][x-1:x+2]
            for _,element := range up{
                total += element
            }
            //self
            center := arr[y][x]
            total += center
            
            //bottom
            bottom := arr[y+1][x-1:x+2]
            for _,element := range bottom{
                total += element
            }
            if total > max{
                max = total
            }
        }
    }
    return max
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

/*********************************************************************************/
Day 12: Inheritance （Golang does not have the property, here is python3 solution）

class Person:
	def __init__(self, firstName, lastName, idNumber):
		self.firstName = firstName
		self.lastName = lastName
		self.idNumber = idNumber
	def printPerson(self):
		print("Name:", self.lastName + ",", self.firstName)
		print("ID:", self.idNumber)

class Student(Person):
    #   Class Constructor
    #   
    #   Parameters:
    #   firstName - A string denoting the Person's first name.
    #   lastName - A string denoting the Person's last name.
    #   id - An integer denoting the Person's ID number.
    #   scores - An array of integers denoting the Person's test scores.
    #
    # Write your constructor here
    def __init__(self,firstName,lastName,idNumber,testScore):
        Person.__init__(Person,firstName,lastName,idNumber)
        self.testScore=testScore

    #   Function Name: calculate
    #   Return: A character denoting the grade.
    #
    # Write your function here
    def calculate(self):
        average=sum(self.testScore)/len(self.testScore)
        if average>=90 and average<=100:
            return "O"
        elif average>=80 and average<90:
            return "E"
        elif average>=70 and average<80:
            return "A"
        elif average>=55 and average<70:
            return "P"
        elif average>=40 and average<50:
            return "D"
        elif average<40:
            return "T"

line = input().split()
firstName = line[0]
lastName = line[1]
idNum = line[2]
numScores = int(input()) # not needed for Python
scores = list( map(int, input().split()) )
s = Student(firstName, lastName, idNum, scores)
s.printPerson()
print("Grade:", s.calculate())

/***********************************************************/
Day 13: Abstract Classes

from abc import ABCMeta, abstractmethod
class Book(object, metaclass=ABCMeta):
    def __init__(self,title,author):
        self.title=title
        self.author=author   
    @abstractmethod
    def display(): pass

#Write MyBook class
class MyBook(Book):
    def __init__(self, title,author,price):
            self.price = price
            super(MyBook,self).__init__(title,author)
        
        
    def display(self):
        print ("Title: %s" %(self.title))
        print ("Author: %s" %(self.author))
        print ("Price: %s" %(self.price))
        
title=input()
author=input()
price=int(input())
new_novel=MyBook(title,author,price)
new_novel.display()

/*******************************************************************************/
Day 14: Scope

class Difference:
    def __init__(self, a):
        self.__elements = a
        self.maximumDifference = None

    def computeDifference(self):
        self.maximumDifference = max(self.__elements) - min(self.__elements) 
	# Add your code here

# End of Difference class

_ = input()
a = [int(e) for e in input().split(' ')]

d = Difference(a)
d.computeDifference()

print(d.maximumDifference)

/***********************************************************************************/
Day 15: Linked List


class Node:
    def __init__(self,data):
        self.data = data
        self.next = None 
class Solution: 
    def display(self,head):
        current = head
        while current:
            print(current.data,end=' ')
            current = current.next

    def insert(self,head,data): 
    #Complete this method
        if head is None:
            return Node(data)
        head.next = self.insert(head.next,data)
        return head

mylist= Solution()
T=int(input())
head=None
for i in range(T):
    data=int(input())
    head=mylist.insert(head,data)    
mylist.display(head); 	  

/*****************************************************************/
Day 16: Exceptions - String to Integer

#!/bin/python3

import math
import os
import random
import re
import sys



S = input().strip()

try: 
    print (int(S))
except ValueError: 
    print ("Bad String")

/********************************************************************/
Day 17: More Exceptions

#Write your code here
class Calculator():
    def power(self, n, p):
        if  (n < 0) or (p < 0):
            raise Exception('n and p should be non-negative')
        return n**p

myCalculator=Calculator()
T=int(input())
for i in range(T):
    n,p = map(int, input().split())
    try:
        ans=myCalculator.power(n,p)
        print(ans)
    except Exception as e:
        print(e)   

/***********************************************************************/
Day 18: Queues and Stacks

import sys

class Solution:
    # Write your code here
# Declare two instance variables
    def __init__(self):
        self.stack = []
        self.queue = []
    
    # Adds an item to the end of the list
    def pushCharacter(self, item):
        self.stack.append(item)
    
    # Removes an item from the end of the list
    def popCharacter(self):
        return self.stack.pop()
        
    # Adds an item to the end of the list
    def enqueueCharacter(self, item):
        self.queue.append(item)
    
    # Removes an item from the front of the list
    def dequeueCharacter(self):
        return self.queue.pop(0)
# read the string s
s=input()
#Create the Solution class object
obj=Solution()   

l=len(s)
# push/enqueue all the characters of string s to stack
for i in range(l):
    obj.pushCharacter(s[i])
    obj.enqueueCharacter(s[i])
    
isPalindrome=True
'''
pop the top character from stack
dequeue the first character from queue
compare both the characters
''' 
for i in range(l // 2):
    if obj.popCharacter()!=obj.dequeueCharacter():
        isPalindrome=False
        break
#finally print whether string s is palindrome or not.
if isPalindrome:
    print("The word, "+s+", is a palindrome.")
else:
    print("The word, "+s+", is not a palindrome.")    
	       
/***************************************************************************************************/
Day 19: Interfaces

class AdvancedArithmetic(object):
    def divisorSum(n):
        raise NotImplementedError

class Calculator(AdvancedArithmetic):
    def divisorSum(self, n):
        factorslist=[]
        for i in range(1,n+1):
            if n%i==0:
                factorslist.append(i)
        return sum(factorslist)

n = int(input())
my_calculator = Calculator()
s = my_calculator.divisorSum(n)
print("I implemented: " + type(my_calculator).__bases__[0].__name__)
print(s)
	       
/******************************************************************************/
Day 20: Sorting Golang Solution

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

    aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var a []int32

    for i := 0; i < int(n); i++ {
        aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
        checkError(err)
        aItem := int32(aItemTemp)
        a = append(a, aItem)
    }
    
    // Write your code here
    bubbleSort(a)
}
func bubbleSort(arr []int32) []int32{
    n := len(arr)
    swapCount := 0
    for i := 0; i < n; i++ {
        // Track number of elements swapped during a single array traversal
        numberOfSwaps := 0
        
        for j := 0; j < n - 1; j++ {
            // Swap adjacent elements if they are in decreasing order
            if arr[j] > arr[j + 1] {
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
                numberOfSwaps++
                swapCount++
            }
        }
        
        // If no elements were swapped during a traversal, array is sorted
        if (numberOfSwaps == 0) {
            break;
        }
    }
    fmt.Printf("Array is sorted in %d swaps. \n",swapCount)
    fmt.Printf("First Element: %d \n",arr[0])
    fmt.Printf("Last Element: %d \n",arr[n-1])
    return arr
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

/******************************************************************/
 Day 21: Generics （C++ solution, Golang 1.8 version begin to support generics）
               
#include <iostream>
#include <vector>
#include <string>

using namespace std;

/**
*    Name: printArray
*    Print each element of the generic vector on a new line. Do not return anything.
*    @param A generic vector
**/

// Write your code here
template<class T>

void printArray(vector<T>arr){
        for (int i = 0; i < arr.size(); i++) {
        cout<<arr[i]<<endl;
    }
}

int main() {
	int n;
	
	cin >> n;
	vector<int> int_vector(n);
	for (int i = 0; i < n; i++) {
		int value;
		cin >> value;
		int_vector[i] = value;
	}
	
	cin >> n;
	vector<string> string_vector(n);
	for (int i = 0; i < n; i++) {
		string value;
		cin >> value;
		string_vector[i] = value;
	}

	printArray<int>(int_vector);
	printArray<string>(string_vector);

	return 0;
}     
/********************************************************************************************/
Day 22: Binary Search Trees

class Node:
    def __init__(self,data):
        self.right=self.left=None
        self.data = data
class Solution:
    def insert(self,root,data):
        if root==None:
            return Node(data)
        else:
            if data<=root.data:
                cur=self.insert(root.left,data)
                root.left=cur
            else:
                cur=self.insert(root.right,data)
                root.right=cur
        return root

    def getHeight(self,root):
        #Write your code here
        if root == None:
            return -1
        else:
            ldepth = self.getHeight(root.left)
            rdepth = self.getHeight(root.right)
            if ldepth>rdepth:
                return ldepth+1
            else:
                return rdepth+1
T=int(input())
myTree=Solution()
root=None
for i in range(T):
    data=int(input())
    root=myTree.insert(root,data)
height=myTree.getHeight(root)
print(height)       

/***************************************************************************/
Day 23: BST Level-Order Traversal
	       
import sys

class Node:
    def __init__(self,data):
        self.right=self.left=None
        self.data = data
class Solution:
    def insert(self,root,data):
        if root==None:
            return Node(data)
        else:
            if data<=root.data:
                cur=self.insert(root.left,data)
                root.left=cur
            else:
                cur=self.insert(root.right,data)
                root.right=cur
        return root

    def levelOrder(self,root):
        #Write your code here
        from collections import deque
        queue = deque([root])
        while len(queue):
            this = queue.popleft()
            print(this.data, end=" ")
            if this.left:
                queue.append(this.left)
            if this.right:
                queue.append(this.right)
T=int(input())
myTree=Solution()
root=None
for i in range(T):
    data=int(input())
    root=myTree.insert(root,data)
myTree.levelOrder(root)

//go solution 
func levelOrder(root *TreeNode) [][]int {
    // 通过上一层的长度确定下一层的元素
    result := make([][]int, 0)
    if root == nil {
        return result
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    for len(queue) > 0 {
        list := make([]int, 0)
        // 为什么要取length？
        // 记录当前层有多少元素（遍历当前层，再添加下一层）
        l := len(queue)
        for i := 0; i < l; i++ {
            // 出队列
            level := queue[0]
            queue = queue[1:]
            list = append(list, level.Val)
            if level.Left != nil {
                queue = append(queue, level.Left)
            }
            if level.Right != nil {
                queue = append(queue, level.Right)
            }
        }
        result = append(result, list)
    }
    return result
}

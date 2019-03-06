package main

import(
	"fmt"
	"bufio"
	"os"
)

var stack []string
var queue []string

func enqueue(char string, queue []string){
	queue = append(queue,char)
}

func dequeue(queue []string) string{
	if len(queue)<1{
		return ""
	}else{
		temp := queue[0]
		queue = queue[1:]
		return temp
	}
}

func push(char string, stack []string){
	stack = append(stack,char)
}

func pop(stack []string) string{
	if len(stack)<1{
		return ""
	}else{
		temp := stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		return temp
	}
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	text,_ := reader.ReadString('\n')
	l := len(text)
	for i:=0;i<l;i++{
		enqueue(string(text[i]),queue)
		push(string(text[i]),stack)
	}
	isPal :=1
	for j:=0;j<l/2;j++{
		if pop(stack)!=dequeue(queue){
			isPal = 0
		}
	}
	if isPal==1{
		fmt.Printf("Is a palindrome.\n")
	}else{
		fmt.Printf("Is not a palindrome.\n")
	}
}

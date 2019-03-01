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
        if head==None:
            head = Node(data)
            return head
        else:
            temp = head
            while head.next!=None:
                head=head.next;
            newnode = Node(data)
            head.next = newnode
            return temp
    #Complete this method
mylist= Solution()
T=int(input())
head=None
for i in range(T):
    data=int(input())
    head=mylist.insert(head,data)    
mylist.display(head); 	  

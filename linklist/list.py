from typing import List
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def ListPrint(l:ListNode)->List:
    node=[]
    while(l):
        node.append(l.val)
        l=l.next
    return node 

def CreateList(num:List)->ListNode:
    cur=ListNode()
    res=cur
    for _,num in enumerate(num):
        cur.next=ListNode(num)
        cur=cur.next
    return res.next

from list import ListNode,ListPrint,CreateList

class Solution(object):
    def reverseList(self, head:ListNode)->ListNode:
        prev=None
        while(head):
            next=head.next
            head.next=prev

            prev=head
            head=next
        return prev
    
    def reverseList2(self, head:ListNode)->ListNode:
        if head==None or head.next==None:
            return head
        
        newhead=self.reverseList2(head.next)
        head.next.next=head
        head.next=None
        return newhead
    
def main():
    c=Solution()
    l=c.reverseList2(CreateList([1,2,3,4,5]))
    print(ListPrint(l))

if __name__ =='__main__':
    main()

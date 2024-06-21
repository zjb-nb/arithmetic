from list import ListNode,CreateList,ListPrint
class Solution(object):
    def reverseKGroup(self, head:ListNode, k:int)->ListNode:
        if (not head) or (not head.next) :
            return head
        i=0
        start=head 
        while i<k-1 and head:
            head=head.next
            if not head:
                return start 
            i+=1

        next=head.next
        head.next=None
        res = self.swap(start)
        start.next = self.reverseKGroup(next,k)
        return res

    def swap(self,head:ListNode)->ListNode:
        if (not head) or (not head.next):
            return head 
        res = self.swap(head.next)
        head.next.next=head
        head.next=None
        return res 

def main():
    c=Solution()
    l=c.reverseKGroup(CreateList([1,2,3,4,5]),3)
    print( ListPrint(l) )

if __name__ =='__main__':
    main()
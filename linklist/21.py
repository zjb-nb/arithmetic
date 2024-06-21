from list import ListNode,ListPrint,CreateList

class Solution(object):
    def mergeTwoLists(self, list1:ListNode, list2:ListNode)->ListNode:
        if not list1 : return list2
        if not list2: return list1
        if list1.val>list2.val :
            list2.next=self.mergeTwoLists(list1,list2.next)
            return list2
        else:
            list1.next=self.mergeTwoLists(list2,list1.next)
            return list1
        
def main():
    c=Solution()
    l=c.mergeTwoLists(CreateList([1,2,4]),CreateList([1,3,4]))
    print( ListPrint(l) )

if __name__ =='__main__':
    main()
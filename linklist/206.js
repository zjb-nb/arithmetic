function ListNode(val, next) {
  this.val = (val===undefined ? 0 : val)
  this.next = (next===undefined ? null : next)
}

var reverseList = function(head) {
  if (head==undefined || head.next==undefined) {return head;}
  let newhead=reverseList(head.next);
  head.next.next=head;
  head.next=undefined;
  return newhead;
};
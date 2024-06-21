
function ListNode(val, next) {
  this.val = (val===undefined ? 0 : val)
  this.next = (next===undefined ? null : next)
}
 
/**
 * @param {ListNode} list1
 * @param {ListNode} list2
 * @return {ListNode}
 */
var mergeTwoLists = function(list1, list2) {
  if (list1==undefined) {
    return list2;
  }
  if (list2==undefined) {
    return list1;
  }
  if (list1.val>list2.val) {
    list2.next = mergeTwoLists(list1,list2.next)
    return list2
  }else{
    list1.next=mergeTwoLists(list2,list1.next)
    return list1
  }

};
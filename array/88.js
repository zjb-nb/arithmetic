var merge = function(nums1, m, nums2, n) {
  let count=m+n-1;
  while(n-1>=0){
    nums1[count] = (m-1>=0 && nums1[m-1]>nums2[n-1])?
      nums1[(m--)-1]:nums2[(n--)-1];
    count--;
  }
};

let t1=[1,2,3,0,0,0],t2=[2,5,6];
merge(t1,3,t2,3);
console.log(t1)
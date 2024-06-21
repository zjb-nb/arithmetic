var maxArea = function(height) {
  let s=0,l=0,r=height.length-1;
  while(l<r) {
    s = height[l]<height[r]?
      Math.max(s,(r-l)*height[l++] ):
      Math.max(s,(r-l)*height[r--]);
  }
  return s;
};
console.log(maxArea([1,8,6,2,5,4,8,3,7]));
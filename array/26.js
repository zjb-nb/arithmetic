var removeDuplicates = function(nums) {
  let i=0;
  if(nums.length<1) {
    return i;
  }
  for(let j=0;j<nums.length;j++){
    if (nums[j]>nums[i]) {
      i++;
      nums[i]=nums[j]
    }
  }
  return i+1
};

let test=[1,1,2];
console.log(removeDuplicates(test));
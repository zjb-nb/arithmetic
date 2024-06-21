var moveZeroes = function(nums) {
  let i=0,temp=0;
  for (let j=0;j<nums.length;j++) {
    if (nums[j]!=0) {
      temp=nums[j];
      nums[j]=nums[i];
      nums[i]=temp;
      i++;
    }
  }
};
let  test = [0,1,0,3,12]
moveZeroes(test)
console.log(test)
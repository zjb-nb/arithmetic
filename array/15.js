var threeSum = function(nums) {
  let res = [];
  if (nums.length<3){
    return res;
  }
  nums.sort();
  for (let i=0;i<nums.length-2;i++){
    if(nums[i]>0){
      return res;
    }
    if(i>0 && nums[i]==nums[i-1]){
      continue
    }
    let l=i+1,r=nums.length-1;
    while(l<r) {
      if (nums[i]+nums[l]+nums[r]==0){
        res.push([nums[i],nums[l],nums[r]])
        while(nums[l]==nums[l+1] && l<r) {
          l++
        }
        while(nums[r]==nums[r-1] && l<r) {
          r--
        }
        l++
        r--
      }else if(nums[i]+nums[l]+nums[r]>0){
        r--
      }else{
        l++
      }

    }
  }
  return res;
};

console.log(threeSum([-1,0,1,2,-1,-4]))
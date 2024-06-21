var twoSum = function(nums, target) {
  let hashmap = new Map();
  for (let i=0;i<nums.length;i++){
    if ( hashmap.has(target-nums[i]) ){
      return [i,hashmap.get(target-nums[i])];
    }
    hashmap.set(nums[i],i)
  }
  return []
};

console.log(twoSum([2,7,11,15],9));
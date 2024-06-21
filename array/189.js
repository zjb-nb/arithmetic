var rotate = function(nums, k) {
  let count=0,n=nums.length;
  k=k%n;
  for(let i=0;i<n;i++){
    let cur=nums[i],next_i=i,temp=0;
    while(true){
      next_i=(next_i+k)%n;
      temp = nums[next_i];
      nums[next_i] = cur;
      cur=temp;
      count++;
      if(next_i==i){
        break;
      }
    }
    if (count==n){
      return
    }
  }
};

let t=[1,2,3,4,5,6];
rotate(t,3);
console.log(t);
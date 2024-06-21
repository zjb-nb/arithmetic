var climbStairs = function(n) {
  if(n<3){
    return n;
  }
  let prev=1,cur=2,temp=0;
  for(let i=3;i<=n;i++){
    temp = cur ; 
    cur +=prev ;
    prev = temp;
  }
  return cur
};

console.log(climbStairs(4))
var plusOne = function(digits) {
  let res=[1];
  for (let j=digits.length-1;j>=0;j--){
    if (digits[j]!=9) {
      digits[j]++;
      return digits;
    }
    digits[j]=0;
  }
  return res.concat(digits)
};

console.log(plusOne([9,9]))
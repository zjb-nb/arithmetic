class Solution(object):
    def twoSum(self, nums, target):
        n = len(nums)
        for i in range(n):
            for j in range(i+1,n):
                if nums[j]+nums[i]==target:
                    return [i,j]
                
    
    def twoSum2(self,nums,target):
        hashmap={}
        for i,num in enumerate(nums):
            if hashmap.get(target-num) is not None:
                return [i,hashmap.get(target-num)]
            hashmap[num]=i

def main():
    c=Solution()
    print( c.twoSum2([2,7,11,15],9) )

if __name__ == '__main__':
    main()
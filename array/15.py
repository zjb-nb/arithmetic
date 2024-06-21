class Solution(object):
    def threeSum(self, nums):
        res = []
        if len(nums)<3 :
            return res
        nums=sorted(nums)
        for i in range (len(nums)-2):
            if nums[i]>0:
                return res 
            if i>0 and nums[i]==nums[i-1]:
                continue
            l,r=i+1,len(nums)-1
            while(l<r):
                if nums[i]+nums[l]+nums[r]==0:
                    res.append([nums[i],nums[l],nums[r]])
                    while nums[l]==nums[l+1] and l<r :
                        l+=1
                    while nums[r]==nums[r-1] and l<r :
                        r-=1
                    l+=1
                    r-=1
                elif nums[i]+nums[l]+nums[r]>0:
                    r-=1
                else :
                    l+=1
        return res
                  
def main():
    c=Solution()
    print(c.threeSum([-1,0,1,2,-1,-4]))

if __name__ == '__main__':
    main()
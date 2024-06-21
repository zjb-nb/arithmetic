class Solution(object):
    def removeDuplicates(self, nums):
        i=0
        if len(nums)<1:
            return i
        for j,num in enumerate(nums):
            if num!=nums[i]:
                i+=1
                nums[i]=num
        return i+1
    
def main():
    c=Solution()
    test=[1,1,2]
    print(c.removeDuplicates(test))
    print(test)

if __name__ == '__main__':
    main()
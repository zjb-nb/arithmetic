#nums = [0,1,0,3,12]  res=[1,3,12,0,0]
class Solution(object):
    #快慢双指针
    def moveZeroes(self, nums):
        i=0
        for j in range(len(nums)):
            if nums[j]!=0:
                nums[i],nums[j]=nums[j],nums[i]
                i+=1

    #滚雪球
    def moveZeroes1(self,nums):
        pass


def main():
    c = Solution()
    test = [0,1,0,3,12]
    c.moveZeroes(test)
    print( test )

if __name__ == '__main__':
    main()

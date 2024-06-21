from typing import List
from collections import deque
class Solution:

    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        if len(nums)<k :
            return nums
        deq=deque()
        for i in range(k):
            while deq and nums[i]>deq[-1] :
                deq.pop()
            deq.append(nums[i])
        res=[ deq[0] ]
        for i in range(k,len(nums) ):
            if deq[0]<=i-k:
                deq.popleft()
            while deque and nums[i]>deq[-1]:
                deq.pop()
            deq.append(nums[i])
            res.append( deq[0] )
        return res

def main():
    c=Solution()
    print(c.maxSlidingWindow([1,3,-1,-3,5,3,6,7],3))

if __name__ =='__main__':
    main()
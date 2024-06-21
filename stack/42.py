from typing import List
class Solution:
    '''[0,1,0,2]'''
    def trap(self, height: List[int]) -> int:
        area=0
        stack =[]
        for i in range(len(height)):
            while len(stack)>0 and height[i]>height[stack[len(stack)-1]]:
                right=i
                top=stack.pop()
                if len(stack)==0 :
                    break
                left = stack[len(stack)-1]
                h=min(height[right],height[left])-height[top]
                area+=h*(right-left-1)
            stack.append(i)
        return area
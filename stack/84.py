from typing import List
class Solution:
    '''[2,1,5,6]'''
    def largestRectangleArea(self, heights: List[int]) -> int:
        heights.append(-1)
        area = 0
        stack = []
        for i in range(len(heights)):
            while len(stack)>0 and heights[i]<heights[stack[len(stack)-1]]:
                right=i-1
                left=-1
                h = heights[stack[len(stack)-1]]
                stack.pop()
                while len(stack)>0 and h==heights[ stack[len(stack)-1] ] :
                    stack = stack.pop()
                if len(stack)>0 :
                    left = stack[len(stack)-1]  
                area =max(area,h*(right-left))
            stack.append(i)
        return area

def main():
    c=Solution()
    print(c.largestRectangleArea([2,1,5,6,2,3]))

if __name__ == '__main__':
    main()
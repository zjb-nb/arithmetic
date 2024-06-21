# review [1,8,6,2,5,4,8,3,7]
#s= h * w = min(height[i],height[j])*(j-i)
class Solution(object):
    def maxArea(self, height):
        #todo 核心还是移动矮的边
        s,l,r=0,0,len(height)-1
        while l<r:
            if height[l]<height[r]:
                s=max(s,height[l]*(r-l))
                l+=1
            else:
                s=max(s,height[r]*(r-l))
                r-=1
        return s
    
def main():
    c = Solution()
    print(c.maxArea([1,8,6,2,5,4,8,3,7]))

if __name__ == '__main__':
    main()
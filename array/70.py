#n阶楼梯
class Solution(object):
    def climbStairs(self, n):
        if n<3 :
            return n
        prev,cur = 1,2
        for i in range(3,n+1):
            prev,cur = cur,cur+prev
        return cur
        
def main():
    c = Solution()
    print(c.climbStairs(2))
    print(c.climbStairs(3))
    print(c.climbStairs(4))

if __name__=='__main__':
    main()
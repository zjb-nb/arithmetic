class Solution:
    def rotate(self, nums, k) :
        count,n=0,len(nums)
        k=k%n
        for i in range(n):
            next,temp=i,nums[i]
            while(True):
                next = (next+k)%n 
                temp,nums[next] = nums[next],temp
                count+=1
                if next==i:
                    break
            if count==n :
                break


def main():
    c=Solution()
    l=[1,2,3,4,5,6,7]
    c.rotate(l,3)
    print(l)

if __name__ =='__main__':
    main()
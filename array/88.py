class Solution(object):
    def merge(self, nums1, m, nums2, n):
        count=m+n-1
        while(n-1>=0):
            if m-1>=0 and nums1[m-1]>nums2[n-1]:
                nums1[count]=nums1[m-1]
                m-=1
            else:
                nums1[count] = nums2[n-1]
                n-=1
            count-=1

def main():
    c=Solution()
    t1,t2=[1,2,3,0,0,0],[2,5,6]
    c.merge(t1,3,t2,3)
    print(t1)

if __name__ == '__main__':
    main()
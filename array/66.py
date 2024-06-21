class Solution(object):
    def plusOne(self, digits):
        for i in range(len(digits)-1,-1,-1):
            if digits[i]!=9 :
                digits[i]+=1
                return digits
            digits[i]=0
        return [1]+digits
        "return [1]+[0]*digits"
    
def main():
    c=Solution()
    t1,t2=[9,9],[1,2,9]
    print( c.plusOne(t1) )
    print( c.plusOne(t2) )


if __name__ =='__main__':
    main()
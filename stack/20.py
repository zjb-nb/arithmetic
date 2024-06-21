class Solution(object):
    def __init__(self) -> None:
        self.hashmap = {")":"(","]":"[","}":"{"}
    def isValid(self, s:str)->bool:
        if len(s)%2>0 :
            return False
        stack=[]
        for _,ch in enumerate( s) :
            if self.hashmap.get(ch):
                if len(stack)==0 or stack[len(stack)-1]!=self.hashmap.get(ch):
                    return False
                stack.pop() 
                continue
            stack.append(ch)
        return len(stack)==0
    
def main():
    c=Solution()
    print(  c.isValid("()[]") )
    print(  c.isValid("([])") )
    print(  c.isValid("([](") )

if __name__=='__main__':
    main()

        
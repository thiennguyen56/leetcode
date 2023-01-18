class Solution:
    def isHappy(self, n: int) -> bool:
        while True:
            if n == 1:
                return True
            n = self.sumOfSquares(n)   

            if n == 4:
                return False
        return False

    def sumOfSquares(self, input) -> int:
        sum = 0
        for i in iter(str(input)):
            sum = sum + (int(i)**2)
        return sum

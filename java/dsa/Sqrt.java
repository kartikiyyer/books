class Solution {

    public int mySqrt(int x) {
        int i =1;
        int count = 0;
        if (x < 1) {
            return 0;
        }
        
        while(x >= i) {
            x -= i;
            i += 2;
            count++;
        }
        return count;
    }
}
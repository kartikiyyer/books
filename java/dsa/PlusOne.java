class Solution {
    
    public int[] plusOne(int[] digits) {
    
    int carry = 1;
    
    for (int i = digits.length-1; i>= 0; i--) {
        digits[i] += carry;
        if (digits[i] != 10) // early return 
            return digits;
        digits[i] = 0;
    }
    
    //if upper return statement didn't run, it means [9],[9,9] type of test case is present.
    
    int[] ans = new int[digits.length+1]; 
    ans[0] = 1; //As other elements will be 0, therefore only updating the 0th element
    return ans;
    
    }

    public int[] plusOne(int[] digits) {
        int carry = 1;
        for(int i = digits.length-1;i>=0;i--) {
            int sum = digits[i] + carry;
            if(sum == 10) {
                digits[i] = 0;
            } else {
                carry = 0;
                digits[i] = sum;
            }   
        }
        if(carry == 1) {
            int[] newdigits = new int[digits.length+1];
            newdigits[0] = 1;
            for(int i = 0;i < digits.length;i++) {
                newdigits[i+1] = digits[i]; 
            }
            
            return newdigits;
        }
        return digits;
    }
}
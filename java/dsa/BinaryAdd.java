class Solution {
    
     public String addBinary(String a, String b) {
        StringBuilder sb = new StringBuilder();
        int i = a.length() - 1, j = b.length() -1, carry = 0;
        while (i >= 0 || j >= 0) {
            int sum = carry;
            if (j >= 0) sum += b.charAt(j--) - '0';
            if (i >= 0) sum += a.charAt(i--) - '0';
            sb.append(sum % 2);
            carry = sum / 2;
        }
        if (carry != 0) sb.append(carry);
        return sb.reverse().toString();
    }


    public String addBinary(String a, String b) {
        if(a.length()==0&&b.length()==0) {
            return "0";
        }
        if(a.length()==0) {
            return b;
        }
        if(b.length()==0) {
            return a;
        }
        
        int i = a.length() - 1;
        int j = b.length() -1;
        int carry =0;
        char ac = 0;
        char bc = 0;
        String sum = new String();
        while(j >= 0 && i >=0) {
            ac = a.charAt(i);
            bc = b.charAt(j);
            if(ac == '1' && bc == '1') {
                if(carry == 1) {
                    sum = "1" + sum;
                } else {
                    sum = "0" + sum;
                }
                carry = 1;
            }
            else if(ac == '0' && bc == '0') {
                if(carry == 1) {
                    sum = "1" + sum;
                } else {
                    sum = "0" + sum;
                }
                carry = 0;
            }
            else {
                if(carry == 1) {
                    sum = "0" + sum;
                    carry = 1;
                } else {
                    sum = "1" + sum;
                    carry = 0;
                }
                
            }
            System.out.println(sum+ " " +carry);
            i--;
            j--;
        }
        
        while(i>=0){
            ac = a.charAt(i);
            if(ac == '1'){     
                if (carry ==1) {
                    sum = "0" + sum;
                    carry = 1;
                } else {
                    sum = "1" + sum;
                    carry = 0;
                }
            } else {
                if(carry == 1) {
                    sum = "1" + sum;
                    carry = 0;
                } else {
                    sum = "0" + sum;
                }
                carry = 0;
            }
            i--;
        } 
        while(j>=0){
            bc = b.charAt(j);
            if(bc == '1'){     
                if (carry ==1) {
                    sum = "0" + sum;
                    carry = 1;
                } else {
                    sum = "1" + sum;
                    carry = 0;
                }
            } else {
                if(carry == 1) {
                    sum = "1" + sum;
                } else {
                    sum = "0" + sum;
                }
                carry = 0;
            }
            j--;
        } 
        
        if(carry == 1) {
            sum = "1" + sum;
        }
        return sum;
    }
}
class Solution {
    public boolean isPalindrome(int x) {
        // Using string builder
        StringBuilder sb = new StringBuilder();
        StringBuilder sbr = new StringBuilder();
        sb.append(x);
        sbr.append(x);
        sbr.reverse();
        return sbr.toString().equals(sb.toString());
    }

    public boolean isPalindrome(int x) {
        // Using string
        String s = "" + x;
        System.out.println(s);
  
        String sr = new String();
        for (int i=s.length()-1;i>=0;i--) {
            sr += s.charAt(i);
        }
        return sr.equals(s);
    }
}
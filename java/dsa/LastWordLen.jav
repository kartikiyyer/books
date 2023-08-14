class Solution {
    public int lengthOfLastWord(String s) {
        String[] ss = s.split(" ");
        return ss[ss.length-1].length();
    }

    public int lengthOfLastWord(String s) {
        int len = 0;
        boolean flag = false;
        for(int i = s.length()-1;i>=0;i--) {
            if(s.charAt(i) == ' ') {
                if(flag) {
                    break;
                }
            } else {
                flag = true;
                len++;
            }
        }
        return len;
    }
}
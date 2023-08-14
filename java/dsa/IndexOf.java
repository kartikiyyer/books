class Solution {
    public int strStr(String haystack, String needle) {
        return haystack.indexOf(needle);
    }

    public int strStr(String haystack, String needle) {
        for(int i=0;i<=haystack.length()-needle.length();i++){
            if(haystack.substring(i,i+needle.length()).equals(needle)){
                return i;
            }
        }
        return -1;
    }

    public int strStr(String haystack, String needle) {
        if(needle.isEmpty()) {
            return 0;
        }
        
        int curr;
        for(int i=0;i<haystack.length();i++) {
            curr = i;
            boolean flag = true;
            int j = 0;
            while(j<needle.length()) {
                
                if(curr >= haystack.length() || (haystack.charAt(curr) != needle.charAt(j))) {
                    flag = false;
                }
                j++;
                curr++;
            }
            if(flag) {
                return i;
            }
        }
        return -1;
    }
}
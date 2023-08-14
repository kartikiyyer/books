class Solution {
    //Bruteforce
    public boolean isValid(String s) {
        Stack<Character>stack = new Stack<Character>();
        for(int i=0;i<s.length();i++) {
            if(s.charAt(i) == ')') {
                if(stack.empty()) {
                    return false;
                }
                int out = stack.pop();
                if(out != '(') {
                    return false;
                }
            } else if(s.charAt(i) == ']') {
                if(stack.empty()) {
                    return false;
                }
                int out = stack.pop();
                if(out != '[') {
                    return false;
                }
            } else if(s.charAt(i) == '}') {
                if(stack.empty()) {
                    return false;
                }
                int out = stack.pop();
                if(out != '{') {
                    return false;
                }
            } else {
                stack.push(s.charAt(i));
            }
        }
        if(!stack.empty()) {
            return false;
        }
        return true;
    }
}
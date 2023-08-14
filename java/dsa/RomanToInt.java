class Solution {
    public int romanToInt(String s) {
    int result = 0;
    if (s == null || s.isBlank() || s.length() < 1 || s.length() > 15) {
      return result;
    }

    int previous = 0;
    for (int i = s.length() - 1; i >= 0; i--) {
      int current = switch (s.charAt(i)) {
        case 'I' -> 1;
        case 'V' -> 5;
        case 'X' -> 10;
        case 'L' -> 50;
        case 'C' -> 100;
        case 'D' -> 500;
        case 'M' -> 1000;
        default -> 0;
      };
      result = previous > current ? result - current : result + current;
      previous = current;
    }

    return result;
  }

    public int romanToInt(String s) {
        HashMap<Character,Integer> m = new HashMap<Character,Integer>();
        m.put('I',1);
        m.put('V',5);
        m.put('X',10);
        m.put('L',50);
        m.put('C',100);
        m.put('D',500);
        m.put('M',1000);
        
        char prev = 0;
        int value = 0;
        for (char c : s.toCharArray()) {
            System.out.println(c);
            value += m.get(c);
            if (prev != 0 && m.get(prev) < m.get(c)) {
                System.out.println("Inside");
                value -= (m.get(prev) * 2);
            }
            
            prev = c;
        }
        return value;
    }

        public int romanToInt(String s) {
        HashMap<Character,Integer> m = new HashMap<Character,Integer>();
        m.put('I',1);
        m.put('V',5);
        m.put('X',10);
        m.put('L',50);
        m.put('C',100);
        m.put('D',500);
        m.put('M',1000);
        
        int prevValue = 0,curValue = 0;
        int value = 0;
        for (char c : s.toCharArray()) {
            System.out.println(c);
            curValue = m.get(c);
            value += curValue;
            if (prevValue != 0 && prevValue < curValue) {
                System.out.println("Inside");
                value -= (prevValue * 2);
            }
            
            prevValue = curValue;
        }
        return value;
    }


}
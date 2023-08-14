import java.util.*;

class Solution {
    public static int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> m = new HashMap<Integer, Integer>();
        for(int i=0;i<nums.length;i++) {
            int diff = target - nums[i];
            if(m.containsKey(diff)) {
                return new int[]{m.get(diff), i};
            }
            m.put(nums[i],i);
        }
        return new int[]{};
    }
    public static void main(String[] args) {
        System.out.println(Solution.twoSum(new int[]{2,7,11,15}, 9));
    }
}
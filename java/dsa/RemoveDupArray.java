class Solution {
    public int removeDuplicates(int[] nums) {
        if(nums.length == 0 || nums.length == 1) {
            return nums.length;
        }
        
        int i=0,j=1,temp;
        
        while(j < nums.length) {
            if(nums[i]!= nums[j]) {
                i++;
                temp = nums[i];
                nums[i] = nums[j];
                nums[j] = temp;
            }
            j++;
        }
        return i+1;
        
    }
}
class Solution {

        public int searchInsert(int[] nums, int target) {
        if(target < nums[0]) {
            return 0;
        }
        
        if(target > nums[nums.length-1]) {
            return nums.length;
        }
        
        int low = 0;
        int high = nums.length;
        int mid = 0;
        
        while(low < high) {
            mid = low + (high - low) /2;
            
            if(nums[mid] == target) {
                return mid;
            }
            else if(nums[mid] > target) {
                high = mid;
            }
            else if(nums[mid] < target) {
                low = mid + 1;
            }
        }
        return low;
    }



    // Bruteforce
    public int searchInsert(int[] nums, int target) {
        if(target < nums[0]) {
            return 0;
        }
        
        if(target > nums[nums.length-1]) {
            return nums.length;
        }
        
        for(int i=0;i<nums.length;i++) {
            if(nums[i] == target) {
                return i;
            }
            if(nums[i] > target) {
                return i;
            }
        }
        return nums.length;
    }
}
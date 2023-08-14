class Solution {


//Approach2 merge sort
    int[] mix=new int[m+n];
    int i=0;
    int j=0;
    int k=0;
    while(i<m&&j<n){
        if(nums1[i]<nums2[j]){
            mix[k]=nums1[i];
            i++;
            
        }else{
            mix[k]=nums2[j];
            j++;
        }
        k++;
    }
    while(i<m){
        mix[k]=nums1[i];
        i++;
        k++;
    }
    while(j<n){
        mix[k]=nums2[j];
        j++;
        k++;
    }
    for(int x=0;x<mix.length;x++){
        nums1[x]=mix[x];
    }
    
    //Approach1
     int j=0;
     for(int i=m;i<nums1.length;i++)
     {
         nums1[i]=nums2[j++];
     }
    Arrays.sort(nums1);



    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int i =0, j = 0;
        while(i<m && j<n) {
            if(nums1[i] > nums2[j]) {
                // Make space.
                for(int k=m;k>i;k--) {
                    nums1[k] = nums1[k-1];
                }
                nums1[i] = nums2[j];
                j++;
                m++;
            }
            i++;
        }
        while(j<n){
            nums1[i++] = nums2[j++];
        }
    }
}
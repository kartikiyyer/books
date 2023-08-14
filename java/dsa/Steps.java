class Solution {



    public int climbStairs(int n) {
        var tab = new int[n+1];
        tab[0] = 1;
        tab[1] = 1;
        for(var i = 2; i<=n ; i++)
        {
            tab[i] = tab[i-1] + tab[i-2];
        }
        return tab[n];
    }

    
    int[] dp = new int[46];
    
    public int climbStairs(int n) {
        Arrays.fill(dp,-1);
        int i=0;
        int x = climb(n,0);
        return x;
    }
    public int climb(int n, int i){
        if(i > n)
            return 0;
        if(i == n)
            return 1;
        if(dp[i] != -1)
            return dp[i];
        return dp[i]=climb(n,i+1)+climb(n,i+2);
    }
}
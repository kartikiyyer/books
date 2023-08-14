class Solution {
    List<Integer> out= new ArrayList<Integer>();
    void inorder(TreeNode n) {
        if(n == null) {
            return;
        }
        inorder(n.left);
        out.add(n.val);
        inorder(n.right);
    }
    
    public List<Integer> inorderTraversal(TreeNode root) {
        inorder(root);
        return out;
    }
}
package template;

public class RBTree<K extends Comparable<K>, V>{
    RBNode<K, V> root;
    RBTree(){
        root = null;
    }
    public RBNode<K, V> get(K key){
        if(key == null) throw new IllegalArgumentException("First argument to put() is null.");
        return get(root,key);
    }
    public void put(K key, V val){
        if(key == null) throw new IllegalArgumentException("First argument to put() is null.");
        if(val == null) throw new IllegalArgumentException("Second argument to put() is null.");
        put(root, key, val);
    }
    public boolean contains(K key) { 
        if (key == null) throw new IllegalArgumentException("First argument to contains() is null.");
        return get(root, key) != null;
    }
    public void delete(K key){
        if (key == null) throw new IllegalArgumentException("First argument to delete() is null.");
        RBNode<K, V> x = get(key);
        if(x != null) delete(x);
    }
    private RBNode<K, V> get(RBNode<K, V> node, K key){
        while (node != null) {
            int compareRes = key.compareTo(node.key);
            if (compareRes < 0) node = node.left; // 在左子树中寻找
            else if (compareRes > 0) node = node.right; // 在右子树中寻找
            else return node; // 找到返回该目标 key 的 val
        }
        return null;
    }
    private void put(RBNode<K, V> node, K key, V val){//实际的put函数
        int comRes;
        RBNode<K, V> cur = node, p = null;
        while(cur != null){ //寻找插入位置
            p = cur;
            comRes = key.compareTo(cur.key);
            if(comRes == 0){
                cur.val = val;
                return;
            } 
            cur = comRes < 0 ? cur.left : cur.right;
        }
        RBNode<K,V> newNode = new RBNode<>(key, val, RBNode.RED);
        newNode.parent = p;
        if(p == null) this.root = newNode; //若无父节点则为根节点
        else{ //插入节点
            comRes = newNode.key.compareTo(p.key);
            if (comRes < 0) p.left = newNode;
            else p.right = newNode;
        }
        fixAfterPut(newNode); //调整红黑树
    }
    private void delete(RBNode<K, V> node){
        if(node.left != null && node.right != null){
            RBNode<K, V> s = node.getSuccessNode();
            node.key = s.key;
            node.val = s.val;
            node = s;
        }
        RBNode<K, V> r = node.left != null ? node.left : node.right;
        if(r != null){
            r.parent = node.parent;
            if(r.parent == root) root = r;
            else if(node == node.parent.left) node.parent.left = r;
            else node.parent.right = r;
            node.left = node.right = node.parent = null;
            r.setColor(RBNode.BLACK);
        }else if(node.parent == null) root = null;
        else{
            if(isBlack(node)) fixAfterDelete(node);
            if(node.parent != null){
                if(node == node.parent.left) node.parent.left = null;
                else if(node == node.parent.right) node.parent.right = null;
                node.parent = null;
            }
        }
    }
    public void printTree(){
        printTree(root);
    }
    private void printTree(RBNode<K, V> t){
        if(t != null) {
            printTree(t.left);
            System.out.print(t.key + " ");
            printTree(t.right);
        }
    }
    private void fixAfterPut(RBNode<K, V> node){
        RBNode<K, V> p, g, u;
        while(((p = node.getParent()) != null) && isRed(p)){
            g = p.getParent();
            u = p.getBrother();
            if(p == g.left){
                if(isRed(u)){
                    toggleColors(g);
                    node = g;
                }else{
                    if(node == p.right){
                        node = p;
                        rotateLeft(node);
                    }
                    p.toggleColor();
                    g.toggleColor();
                    rotateRight(g);
                }
            }else{
                if(isRed(u)){
                    toggleColors(g);
                    node = g;
                }else{
                    if(node == p.left){
                        node = p;
                        rotateRight(node);
                    }
                    p.toggleColor();
                    g.toggleColor();
                    rotateLeft(g);
                }
            }
        }
        if(node == root) node.setColor(RBNode.BLACK);
    }
    public void fixAfterDelete(RBNode<K, V> node){
        RBNode<K, V> p = parentOf(node);
        while(isBlack(node) && node != root){
            if(node == leftOf(p)){
                RBNode<K, V> x = rightOf(p);
                if(isRed(x)){
                    p.setColor(RBNode.RED);
                    x.setColor(RBNode.BLACK);
                    rotateLeft(p);
                }
                if(isBlack(leftOf(x)) && isBlack(rightOf(x))){
                    if(x != null) x.setColor(RBNode.RED);
                    node = p;
                    p = parentOf(node);
                }else{
                    if(isBlack(rightOf(x))){
                        x.left.setColor(RBNode.BLACK);
                        x.setColor(RBNode.RED);
                        rotateRight(x);
                    }
                    if(x != null) x.setColor(colorOf(p));
                    if(p != null) p.setColor(RBNode.BLACK);
                    if(rightOf(x) != null) x.left.setColor(RBNode.BLACK);
                    rotateLeft(p);
                    x = root;
                }
            }else{
                RBNode<K, V> x = leftOf(p);
                if(isRed(x)){
                    p.setColor(RBNode.RED);
                    x.setColor(RBNode.BLACK);
                    rotateRight(p);
                }
                if(isBlack(leftOf(x)) && isBlack(rightOf(x))){
                    if(x != null) x.setColor(RBNode.RED);
                    node = p;
                    p = parentOf(node);
                }else{
                    if(isBlack(leftOf(x))){
                        x.right.setColor(RBNode.BLACK);
                        x.setColor(RBNode.RED);
                        rotateLeft(x);
                    }
                    if(x != null) x.setColor(colorOf(p));
                    if(p != null) p.setColor(RBNode.BLACK);
                    if(leftOf(x) != null) x.left.setColor(RBNode.BLACK);
                    rotateRight(p);
                    x = root;
               }
            }
        }
    }
    private void toggleColors(RBNode<K, V> node){
        node.toggleColor();
        node.left.toggleColor();
        node.right.toggleColor();
    }
    public void rotateLeft(RBNode<K, V> h) {
        RBNode<K, V> x = h.right; 
        h.right = x.left; 
        if (x.left != null) x.left.parent = h;
        x.parent = h.parent;
        if (h.parent == null) this.root = x;
        else {
            if (h == h.parent.left) h.parent.left = x;
            else h.parent.right = x;
        }
        x.left = h;
        h.parent = x;
    }
    public void rotateRight(RBNode<K, V> h) {
        RBNode<K, V> x = h.left;
        h.left = x.right;
        if (x.right != null) x.right.parent = h;
        x.parent = h.parent;
        if (h.parent == null) this.root = x;
        else {
            if (h == h.parent.right) h.parent.right = x;
            else h.parent.left = x;
        }
        x.right = h;
        h.parent = x;
    }
    private RBNode<K,V> leftOf(RBNode<K,V> p) { // 返回 p 的左子结点
        return (p == null) ? null: p.left;
    }

    private RBNode<K,V> rightOf(RBNode<K,V> p) { // 返回 p 的右子结点
        return (p == null) ? null: p.right;
    }
    private RBNode<K, V> parentOf(RBNode<K, V> p) { // 返回 p 的父结点
        return (p != null) ? p.parent : null;
    }
    private int colorOf(RBNode<K, V> node) { // 返回 node 的颜色
        if (node != null) return node.getColor();
        return RBNode.BLACK;
    }
    private boolean isRed(RBNode<K, V> node) { // 判断 node 是否为红
        return (node != null && node.getColor() == RBNode.RED) ? true : false;
    }
    private boolean isBlack(RBNode<K, V> node) { // 判断 node 是否为黑
        return !isRed(node);
    }
}

class RBNode<K extends Comparable<K>, V>{
    K key;
    V val;
    RBNode<K, V> parent;
    RBNode<K, V> left;
    RBNode<K, V> right;
    boolean color;
    static final int RED = 0;
    static final int BLACK = 1;
    RBNode(){}
    RBNode(K key, V val, int color){
        this.key = key;
        this.val = val;
        this.color = color == 1;
    }
    public RBNode<K, V> getParent() {
        return parent;
    }
    public RBNode<K, V> getGrandParent(){
        if(parent != null) return parent.getParent();
        return null;
    }
    public RBNode<K, V> getBrother(){
        return parent.left == this ? parent.right : parent.left;
    }
    public int getColor(){
        return color ? BLACK : RED;
    }
    public void setColor(int color) {
        this.color = color == 1;
    }
    public void toggleColor(){
        color = !color;
    }
    private RBNode<K, V> minNode() { // 返回最小结点
        if (left == null) return this;
        else return left.minNode();
    }
    public RBNode<K, V> getSuccessNode(){
        if(right != null ) return right.minNode();
        RBNode<K, V> n = this, p = parent;
        while((p != null) && (p.right == n)){
            n = p;
            p = p.parent;
        }
        return p;
    }
}
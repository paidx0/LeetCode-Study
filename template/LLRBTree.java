package template;

public class LLRBTree<K extends Comparable<K>, V> {
    int size;
    RBNode<K, V> root;
    LLRBTree(){
        root = null;
        size = 0;
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
    private RBNode<K, V> get(RBNode<K, V> x, K key) {
        while (x != null) {
            int compareRes = key.compareTo(x.key);
            if (compareRes < 0) x = x.left;
            else if (compareRes > 0) x = x.right;
            else return x;
        }
        return null;
    }
    boolean isRed(RBNode<K, V> node){
        if(node == null) return false;
        return node.color == false;
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
}
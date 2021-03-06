### 堆

堆就是用数组实现的二叉树，所有它没有使用父指针或者子指针。

在堆底插入元素，在堆顶取出元素，但是堆中元素的排列不是按照到来的先后顺序，而是按照一定的优先顺序排列的堆根据“堆属性”来排序

> 堆的特点

- 必须是完全二叉树
- 用数组实现
- 任一结点的值是其子树所有结点的最大值或最小值
    - 最大值时，称为“最大堆”，也称大顶堆；
    - 最小值时，称为“最小堆”，也称小顶堆。

> 堆的基本操作

|操作|描述|时间复杂度|
|---|---|---|
|create|创建一个空堆|O(n)|
|push|向堆中插入一个新元素|O(log n)|
|peek|获取堆顶元素的值|O(1)|
|pop|删除堆顶元素|O(log n)|
|len|获取堆的元素个数|O(1)|

> 堆和二叉搜索树的区别

- 堆并不能取代二叉搜索树

   两者之间的主要差别是节点顺序。在二叉搜索树中，左子节点必须比父节点小，右子节点必须必比父节点大。但是在堆中并非如此。在最大堆中两个子节点都必须比父节点小，而在最小堆中，它们都必须比父节点大。

- 内存占用

    普通树占用的内存空间比它们存储的数据要多。你必须为节点对象以及左/右子节点指针分配额为是我内存。堆仅仅使用一个数据来村塾数组，且不使用指针。

- 平衡
    
    二叉搜索树必须是“平衡”的情况下，其大部分操作的复杂度才能达到O(log n)。你可以按任意顺序位置插入/删除数据，或者使用 AVL 树或者红黑树，但是在堆中实际上不需要整棵树都是有序的。我们只需要满足对属性即可，所以在堆中平衡不是问题。因为堆中数据的组织方式可以保证O(log n) 的性能。

- 搜索
    
    在二叉树中搜索会很快，但是在堆中搜索会很慢。在堆中搜索不是第一优先级，因为使用堆的目的是将最大（或者最小）的节点放在最前面，从而快速的进行相关插入、删除操作。

> 堆的应用场景

- 优先队列 
- 堆排序
- 海量或流数据中的K大和K小

参考链接：

- [堆-维基百科](https://zh.wikipedia.org/wiki/堆積)
- [数据结构-堆-掘金](https://juejin.im/post/59fc75f76fb9a0452206dd15)
- [数据结构-堆-简书](https://www.jianshu.com/p/6b526aa481b1)
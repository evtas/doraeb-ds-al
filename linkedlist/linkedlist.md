# 链表
常见的链表类型：
- 单链表
- 循环链表
- 双向链表
  - 在删除给定指针指向的节点时只需要O(1)的时间复杂度
  - 在指定节点前插入节点也只需要O(1)的时间复杂度
  - 空间换时间思想

## 特点
- 内存空间不连续，跟数组相比需要额外的空间保存后续节点的指针
- 插入、删除操作的时间复杂度为O(1)，随机访问的时间复杂度为O(n)

## 问题
1. 带头链表的作用——方便操作
2. 循环链表是带头链表吗？——不是，尾节点就代替了空头结点的作用

## 链表跟数组的性能比较
1. CPU缓存机制可以预读，对数组更友好
2. 数组大小固定，占用整块连续的内存空间，如果需要扩容，还会有全量复制的操作
3. 链表占用更多的内存，而且容易造成内存碎片，引发gc

## 注意
- 指针丢失跟内存泄漏
- 带头节点是哨兵思想
> 哨兵思想: "哨兵"思想是指在循环中设置一个特殊的元素（称为哨兵），以便在循环过程中能够更高效地处理某些边界情况或结束条件。
- 边界条件
  - 链表为空
  - 只包含一个节点
  - 只包含两个节点
  - 处理头结点跟尾结点
  - 等等

## 使用场景
- LRU算法
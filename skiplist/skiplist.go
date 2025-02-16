package main

import (
	"math/rand"
	"time"
)

// 跳表
// 时间复杂度O(logn)， 空间复杂度O(n)
// 通过随机函数保持平衡性
// 区间查询会比红黑树更高效方便，更加灵活，可以平衡执行效率和内存消耗

// copy from https://leetcode.cn/problems/design-skiplist/solutions/1699646/chao-xiang-xi-golangshi-xian-tiao-biao-b-vzmd/
// author: 风月
// 仅供学习使用
type Skiplist struct {
	head  *Node
	size  int
	level int // 当前层级: 指的是当前跳表中有多少层
}

type Node struct {
	Score int
	value interface{} // 实践中可把int64改成其他数据类型
	next  []*Node
}

const (
	maxLevel   = 16 // 最大层级: 指的是跳表最多能有多少层
	p          = 0.25
	ERROR_CODE = -43962200
)

func newNode(Score int, value interface{}, maxLevel int) *Node {
	return &Node{
		Score: Score,
		value: value,
		next:  make([]*Node, maxLevel),
	}
}

/**
 * 获取随机level
 * 这里参考了skiplist论文中的获取随机level的方法
 */
func randomLevel() int {
	var level = 1
	for level < maxLevel && rands() < p {
		level++
	}
	return level
}

func rands() float32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32()
}

func Constructor() Skiplist {
	return Skiplist{
		head: &Node{next: make([]*Node, maxLevel)},
	}
}

func (sl *Skiplist) Search(searchScore int) bool {
	// 从最高层开始从左向右找（最上面一层可以最快定位到我们想找的元素大概位置）
	c := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		// 退出循环的条件就是next元素大于指定的元素, 就往下一层开始找
		for c.next[i] != nil && c.next[i].Score < searchScore {
			// 当前节点在这一层直接向前
			c = c.next[i]
		}
		// 判断下一个元素是否满足条件, 任何一层, 找到就直接返回对应的值
		if c.next[i] != nil && c.next[i].Score == searchScore {
			return true
		}
	}
	// 找到最下面一层还没有找到则说明元素不存在。
	return false
}

func (sl *Skiplist) Add(score int) {
	//正常来说value应该是传进来的,即Add(score int, value interface{}), 这里根据题目改了一下
	value := 1
	// update用来记录每一层最后一个小于score的节点
	update := make([]*Node, maxLevel)
	c := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for c.next[i] != nil && c.next[i].Score < score {
			// 当前节点在这一层直接向前
			c = c.next[i]
		}
		update[i] = c
	}
	c = c.next[0]
	// 如果这个元素本身就存在跳表里, 则直接更新, 如果允许有重复元素, 注释下面这一个if即可
	if c != nil && c.Score == score {
		c.value = value
		return
	}

	/** 如果随机到的level比当前的level要高, 就更新层级
	 * 一种实现方式是可以把rdmlevel设置current level + 1, 方便查找
	 * 但是破坏了算法的随机性
	 * // The pseudo-code
	 * if rdmlevel > sl.level {
	 * 		rdmlevel = sl.level + 1
	 *		update[sl.level] = sl.head
	 *		sl.level = rdmlevel
	 *	}
	 */
	rdmlevel := randomLevel()
	if rdmlevel > sl.level {
		for i := sl.level; i < rdmlevel; i++ {
			update[i] = sl.head
		}
		sl.level = rdmlevel
	}
	e := newNode(score, value, rdmlevel)
	// 处理每一层的情况
	// o->last => o->update->last
	for i := 0; i < rdmlevel; i++ {
		e.next[i] = update[i].next[i]
		update[i].next[i] = e
	}
	sl.size++
}

func (sl *Skiplist) Erase(searchScore int) bool {
	c := sl.head
	update := make([]*Node, maxLevel)
	for i := sl.level - 1; i >= 0; i-- {
		for c.next[i] != nil && c.next[i].Score < searchScore {
			c = c.next[i]
		}
		// 设置每一层对应的元素信息
		update[i] = c
	}
	// 最下面一层的第一个指向的元素, 即即将被删除的元素
	c = c.next[0]

	if c != nil && c.Score == searchScore {
		for i := 0; i < sl.level; i++ {
			// 处理每一层的情况
			// o->update->last => o->last
			if update[i].next[i] != c {
				break
			}
			update[i].next[i] = c.next[i]
		}
		sl.size--
		return true
	}
	return false
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */

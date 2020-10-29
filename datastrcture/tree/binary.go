package tree

import (
	"github.com/dingkegithub/golanguage/datastrcture/queueandstack"
)

type Node struct {
	Left   *Node
	Right  *Node
	Parent *Node
	Data   interface{}
}

type Binary struct {
	root    *Node
	Compare func(interface{}, interface{}) int
}

func (b *Binary) Empty() bool {
	if b.root == nil {
		return true
	}

	return false
}

func (b *Binary) BuildWithPreMidOrder(pre []interface{}, mid []interface{}) {
	var idx int = 0
	b.root = b.buildWithPreMidOrder(pre, mid, 0, len(pre)-1, &idx)
}

func (b *Binary) buildWithPreMidOrder(pre, mid []interface{}, start, end int, idx *int) *Node {
	if start > end {
		return nil
	}

	// 1> pick and element from pre order and then create new node
	node := &Node{}
	node.Data = pre[*idx]
	// fmt.Println("idx: ", *idx, "start: ", start, "end: ", end, "node: ", node.Data)

	// 2> increment index to pick next element in next recursive call
	*idx += 1

	if start == end {
		return node
	}

	// 3> find index of new node in mid order, we call this index as mid-index
	midIdx := b.search(mid, node.Data, start, end)

	// 4> call _buildTreeWithPreMidOrder for elements before mid-index and make
	// built tree as left subtree of new node
	node.Left = b.buildWithPreMidOrder(pre, mid, start, midIdx-1, idx)

	// 5> call _buildTreeWithPreMidOrder for elements after mid-index and make
	//    built tree as right subtree of new node
	node.Right = b.buildWithPreMidOrder(pre, mid, midIdx+1, end, idx)

	// fmt.Println("----------------------------------------------------")
	// fmt.Println("root->", node.Data)

	// if node.Right != nil {
	// 	fmt.Println("Right-->", node.Right.Data)
	// } else {
	// 	fmt.Println("Right-->", nil)
	// }

	// if node.Left != nil {
	// 	fmt.Println("LEFT-->", node.Left.Data)
	// } else {
	// 	fmt.Println("LEFT-->", nil)
	// }
	// fmt.Println("----------------------------------------------------")

	return node
}

func (b *Binary) BuildWithPostMidOrder(post []interface{}, mid []interface{}) {
	var idx int = 0
	b.root = b.buildWithPreMidOrder(pre, mid, 0, len(pre)-1, &idx)
}

func (b *Binary) buildWithPostMidOrder(pre, mid []interface{}, start, end int, idx *int) *Node {
	if start > end {
		return nil
	}

	// 1> pick and element from pre order and then create new node
	node := &Node{}
	node.Data = pre[*idx]
	// fmt.Println("idx: ", *idx, "start: ", start, "end: ", end, "node: ", node.Data)

	// 2> increment index to pick next element in next recursive call
	*idx += 1

	if start == end {
		return node
	}

	// 3> find index of new node in mid order, we call this index as mid-index
	midIdx := b.search(mid, node.Data, start, end)

	// 4> call _buildTreeWithPreMidOrder for elements before mid-index and make
	// built tree as left subtree of new node
	node.Left = b.buildWithPreMidOrder(pre, mid, start, midIdx-1, idx)

	// 5> call _buildTreeWithPreMidOrder for elements after mid-index and make
	//    built tree as right subtree of new node
	node.Right = b.buildWithPreMidOrder(pre, mid, midIdx+1, end, idx)

	// fmt.Println("----------------------------------------------------")
	// fmt.Println("root->", node.Data)

	// if node.Right != nil {
	// 	fmt.Println("Right-->", node.Right.Data)
	// } else {
	// 	fmt.Println("Right-->", nil)
	// }

	// if node.Left != nil {
	// 	fmt.Println("LEFT-->", node.Left.Data)
	// } else {
	// 	fmt.Println("LEFT-->", nil)
	// }
	// fmt.Println("----------------------------------------------------")

	return node
}

func (b *Binary) search(mid []interface{}, data interface{}, start, end int) int {
	for i := start; i <= end; i++ {
		if b.Compare(mid[i], data) == 0 {
			return i
		}
	}

	return -1
}

func (b *Binary) PreTraverse(res *[]interface{}) {
	b.preTraverse(b.root, res)
}

func (b *Binary) preTraverse(node *Node, res *[]interface{}) {
	if node == nil {
		return
	}

	*res = append(*res, node.Data)
	b.preTraverse(node.Left, res)
	b.preTraverse(node.Right, res)
}

func (b *Binary) MidTraverse(res *[]interface{}) {
	b.midTraverse(b.root, res)
}

func (b *Binary) midTraverse(node *Node, res *[]interface{}) {
	if node == nil {
		return
	}

	b.midTraverse(node.Left, res)
	*res = append(*res, node.Data)
	b.midTraverse(node.Right, res)
}

func (b *Binary) PostTraverse(res *[]interface{}) {
	b.postTraverse(b.root, res)
}

func (b *Binary) postTraverse(node *Node, res *[]interface{}) {
	if node == nil {
		return
	}

	b.postTraverse(node.Left, res)
	b.postTraverse(node.Right, res)
	*res = append(*res, node.Data)
}

func (b *Binary) LevelTraverse(res *[]interface{}) {
	q := queueandstack.New(queueandstack.WithScale())

	q.Put(b.root)

	node, err := q.Get()

	for err == nil {
		*res = append(*res, node.(*Node).Data)
		treeNode := node.(*Node)

		if treeNode.Left != nil {
			q.Put(treeNode.Left)
		}

		if treeNode.Right != nil {
			q.Put(treeNode.Right)
		}
		node, err = q.Get()
	}
}

func (b *Binary) IsFullTree() bool {
	return b.isFullTree(b.root)
}

// 满树判断
// 条件：每个节点要么有两个子节点，要么没有子节点
func (b *Binary) isFullTree(n *Node) bool {
	if n == nil {
		return true
	}

	if n.Left == nil && n.Right == nil {
		return true
	}

	if n.Left != nil && n.Right != nil {
		return b.isFullTree(n.Left) && b.isFullTree(n.Right)
	}

	return false
}

func (b *Binary) IsPerfectTree() bool {
	return b.isPerfectTree(b.root, 0, b.Depth())
}

// 完美二叉树
// 每个节点含有两个子节点，且所以子节点在同一层上
func (b *Binary) isPerfectTree(n *Node, level, depth int) bool {
	if n == nil {
		return true
	}

	if n.Left == nil && n.Right == nil {
		return depth == level+1
	}

	if n.Left == nil || n.Right == nil {
		return false
	}

	return b.isPerfectTree(n.Left, level+1, depth) && b.isPerfectTree(n.Right, level+1, depth)
}

func (b *Binary) Depth() int {
	return b.depth(b.root)
}

func (b *Binary) depth(n *Node) int {
	if n == nil {
		return 0
	}

	lD := b.depth(n.Left)
	rD := b.depth(n.Right)

	if lD > rD {
		return lD + 1
	} else {
		return rD + 1
	}
}

func (b *Binary) NodeCount() int {
	return b.nodeCount(b.root)
}

func (b *Binary) nodeCount(n *Node) int {
	if n == nil {
		return 0
	}

	return 1 + b.nodeCount(n.Left) + b.nodeCount(n.Right)
}

func (b *Binary) IsCompleteTree() bool {
	return b.isCompleteTree(b.root, 0, b.NodeCount())
}

func (b *Binary) isCompleteTree(n *Node, idx, numbers int) bool {
	if n == nil {
		return true
	}

	if idx > numbers {
		return false
	}

	return b.isCompleteTree(n.Left, 2*idx+1, numbers) && b.isCompleteTree(n.Right, 2*idx+2, numbers)
}

func (b *Binary) isBalanceTree(node *Node, height *int) bool {
	var lH, rH int
	var l, r bool

	if node == nil {
		*height = 0
		return true
	}

	l = b.isBalanceTree(node.Left, &lH)
	r = b.isBalanceTree(node.Right, &rH)

	if lH > rH {
		*height = lH + 1
	} else {
		*height = rH + 1
	}

	if lH-rH >= 2 || rH-lH >= 2 {
		return false
	} else {
		return l && r
	}

}

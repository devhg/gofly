package gofly

import "strings"

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 孩子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// matchChild 主要用于插入时，匹配part是否在当前节点的孩子节点群中是否存在
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child // part 是n的孩子，返回孩子节点
		}
	}
	return nil
}

//
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert defines the method to add Elem to trie Tree
func (n *node) insert(pattern string, parts []string, height int) {
	// 走到最后一个， parts长度，就是树深度（此路由的深度）
	if len(parts) == height {
		// 只有完全匹配的最后节点保存pattern全部 /go/:lang/doc。
		n.pattern = pattern
		return
	}

	part := parts[height] // 当前part，比如go，:lang
	// 判断part是不是n的孩子
	child := n.matchChild(part)

	// part不在n的孩子节点中，加入到n的孩子节点群中
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	// 继续走，遍历下一个part
	child.insert(pattern, parts, height+1)
}

// search defines the method to match route
func (n *node) search(parts []string, height int) *node {
	// 1.走到末尾返回，末尾节点保存了整个路由规则的全部内容
	// 2.匹配到*返回
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		search := child.search(parts, height+1)
		if search != nil {
			return search
		}
	}
	return nil
}

func (n *node) travel(list *[]*node) {
	if n.pattern != "" {
		*list = append(*list, n)
	}

	for _, child := range n.children {
		child.travel(list)
	}
}

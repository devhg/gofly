package gofly

import (
	"fmt"
	"testing"
)

func Test_node_insert(t *testing.T) {
	root := &node{}
	root.insert("/doc/:lang", parsePattern("/doc/:lang"), 0)
	root.insert("/doc/zhcn/go", parsePattern("/doc/zhcn/go"), 0)

	nodes := make([]*node, 0)
	root.travel(&nodes)
	for _, v := range nodes {
		fmt.Printf("%p, %v\n", v, v)
	}
	//0xc0000129c0, &{/doc/:lang :lang [0xc000012a80] true}
	//0xc000012a80, &{/doc/zhcn/go go [] false}  // 上面的孩子

	search := root.search(parsePattern("/doc/en"), 0)
	fmt.Println(search)
	search2 := root.search(parsePattern("/a"), 0)
	fmt.Println(search2)
	search3 := root.search(parsePattern("/doc/zhcn/go"), 0)
	fmt.Println(search3)
	search4 := root.search(parsePattern("/doc/zhcn/java"), 0)
	fmt.Println(search4)
}

func Test_node_search(t *testing.T) {
	root := &node{}
	root.insert("/doc/:lang", parsePattern("/doc/:lang"), 0)
	root.insert("/doc/zhcn/go", parsePattern("/doc/zhcn/go"), 0)

	search := root.search(parsePattern("/doc/en"), 0)
	fmt.Println(search)
	search2 := root.search(parsePattern("/a"), 0)
	fmt.Println(search2)
	search3 := root.search(parsePattern("/doc/zhcn/go"), 0)
	fmt.Println(search3)
	search4 := root.search(parsePattern("/doc/zhcn/java"), 0)
	fmt.Println(search4)
}

func Test_node_travel(t *testing.T) {

}

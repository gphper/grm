/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-23 21:07:01
 */
package common

import "sort"

// 字典树节点
type TrieNode struct {
	Children map[string]*TrieNode
	Types    string
	IsEnd    bool
}

// 构造字典树节点
func newTrieNode() *TrieNode {
	return &TrieNode{Children: make(map[string]*TrieNode), IsEnd: false, Types: "none"}
}

// 字典树
type Trie struct {
	Root *TrieNode
}

// 构造字典树
func NewTrie() *Trie {
	return &Trie{Root: newTrieNode()}
}

// 向字典树中插入一个单词
func (trie *Trie) Insert(word []string, types string) {
	node := trie.Root
	for i := 0; i < len(word); i++ {
		_, ok := node.Children[word[i]]
		if !ok {
			node.Children[word[i]] = newTrieNode()
		}
		node = node.Children[word[i]]
	}

	node.IsEnd = true
	node.Types = types
}

type Node struct {
	Title    string `json:"label"`
	All      string `json:"id"`
	Children []Node `json:"children"`
}

func GetOne(children map[string]*TrieNode, pre string) []Node {
	items := make([]Node, 0)
	for k, v := range children {
		if v.IsEnd && len(v.Children) != 0 {
			tmp1 := Node{Title: k, All: pre + k, Children: nil}
			items = append(items, tmp1)
		}
		tmp := Node{Title: k, All: pre + k, Children: nil}
		if len(v.Children) != 0 {
			tmp.Children = GetOne(v.Children, pre+k+":")
		}
		items = append(items, tmp)
	}

	sort.Slice(items, func(i, j int) bool {

		if items[i].Title == items[j].Title {
			return len(items[i].Children) > len(items[j].Children)
		} else {
			return items[i].Title < items[j].Title
		}

	})

	return items

}

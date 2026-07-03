package tree

import (
	"path/filepath"
	"sort"

	"github.com/codexgenome/codex-genome/internal/scanner"
)

// Node is a single entry in the project tree.
//
// Directory nodes have a non-empty Children slice and a nil File pointer.
// File nodes are leaf nodes: Children is always nil and File is non-nil.
// The root node has a nil Parent.
type Node struct {
	Name         string
	RelativePath string
	Depth        int
	Parent       *Node
	Children     []*Node
	File         *scanner.File // non-nil only for file leaf nodes
}

// IsFile reports whether n is a file leaf node.
func (n *Node) IsFile() bool { return n.File != nil }

// Build constructs an in-memory project tree from a scanned Project.
//
// The returned root node represents the scanned root directory.
// Children at every level are sorted alphabetically by name.
// Construction is O(n) in the number of entries.
func Build(project *scanner.Project) *Node {
	root := &Node{
		Name:         filepath.Base(project.RootPath),
		RelativePath: ".",
		Depth:        0,
	}

	// index maps RelativePath → *Node for O(1) parent lookups.
	index := map[string]*Node{".": root}

	for i := range project.Directories {
		dir := &project.Directories[i]
		node := &Node{
			Name:         dir.Name,
			RelativePath: dir.RelativePath,
			Depth:        dir.Depth,
		}
		link(index[parentOf(dir.RelativePath)], node)
		index[dir.RelativePath] = node
	}

	for i := range project.Files {
		f := &project.Files[i]
		node := &Node{
			Name:         f.Name,
			RelativePath: f.RelativePath,
			Depth:        f.Depth,
			File:         f,
		}
		link(index[parentOf(f.RelativePath)], node)
	}

	sortChildren(root)
	return root
}

// link attaches child to parent, setting the bidirectional relationship.
func link(parent, child *Node) {
	child.Parent = parent
	parent.Children = append(parent.Children, child)
}

// parentOf returns the relative path of the parent directory.
// filepath.Dir handles both Unix ("/") and Windows ("\") separators and
// returns "." for top-level entries, which maps to the root node in the index.
func parentOf(rel string) string {
	return filepath.Dir(rel)
}

// sortChildren recursively sorts each node's children alphabetically by name.
func sortChildren(n *Node) {
	sort.Slice(n.Children, func(i, j int) bool {
		return n.Children[i].Name < n.Children[j].Name
	})
	for _, child := range n.Children {
		sortChildren(child)
	}
}

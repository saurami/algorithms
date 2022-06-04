package tree

import "testing"

func TestInsert(t *testing.T) {

	/*

	  Tree Structure

			10
		   /  \
		  /    \
		 8     20
		/ \      \
	   /   \      \
	  5     9     22

	*/

	tree := &BinarySearchTree{}
	t.Run("root node", func(t *testing.T) {
		got := tree.insert(10)
		want := 10
		if got.root.data != want {
			t.Errorf("Expected to receive %v; but got %v", got, want)
		}
	})

	t.Run("left node level 1", func(t *testing.T) {
		got := tree.insert(8)
		want := 8
		if got.root.left.data != want {
			t.Errorf("Expected to receive %v; but got %v", got, got)
		}
	})

	t.Run("left node level 2", func(t *testing.T) {
		got := tree.insert(5)
		want := 5
		if got.root.left.left.data != want {
			t.Errorf("Expected to receive %v; but got %v", got, got)
		}
	})

	t.Run("right node level 2", func(t *testing.T) {
		got := tree.insert(9)
		want := 9
		if got.root.left.right.data != want {
			t.Errorf("Expected to receive %v; but got %v", got, got)
		}
	})

	t.Run("right node level 1", func(t *testing.T) {
		got := tree.insert(20)
		want := 20
		if got.root.right.data != want {
			t.Errorf("Expected to receive %v; but got %v", got, got)
		}
	})

	t.Run("right node level 2", func(t *testing.T) {
		got := tree.insert(22)
		want := 22
		if got.root.right.right.data != want {
			t.Errorf("Expected to receive %v; but got %v", got, got)
		}
	})

}

func TestSearch(t *testing.T) {
	tree := &BinarySearchTree{}
	tree.insert(10)
	tree.insert(8)
	tree.insert(5)
	tree.insert(9)
	tree.insert(20)

	t.Run("value exists", func(t *testing.T) {
		got := tree.search(10)
		want := true
		if got != want {
			t.Errorf("Value is present; got %v, wanted %v", got, want)
		}
	})

	t.Run("value does not exist", func(t *testing.T) {
		got := tree.search(2)
		want := false
		if got != want {
			t.Errorf("Value doesn't exist; got %v; wanted %v", got, want)
		}
	})

	t.Run("value does not exist", func(t *testing.T) {
		got := tree.search(23)
		want := false
		if got != want {
			t.Errorf("Value doesn't exist; got %v; wanted %v", got, want)
		}
	})
}

// Edge case - search empty tree
func TestSearchEmptyTree(t *testing.T) {
	tree := &BinarySearchTree{}
	got := tree.search(10)
	want := false
	if got != want {
		t.Errorf("Tree doesn't exist; got %v; wanted %v", got, want)
	}
}

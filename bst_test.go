package datastructures_test

import (
	"testing"

	ds "github.com/andreev1024/go-datastructures"
)

const TEST_FAILED = "Test failed"

func TestInsert(t *testing.T) {

	bst := ds.BST{}
	bst.Insert(5)
	bst.Insert(4)
	bst.Insert(7)
	bst.Insert(6)
	bst.Insert(8)

	if bst.Root.Val != 5 {
		t.Error(TEST_FAILED)
	}

	if bst.Root.Left.Val != 4 {
		t.Error(TEST_FAILED)
	}

	if bst.Root.Right.Val != 7 {
		t.Error(TEST_FAILED)
	}

	if bst.Root.Right.Left.Val != 6 {
		t.Error(TEST_FAILED)
	}

	if bst.Root.Right.Right.Val != 8 {
		t.Error(TEST_FAILED)
	}
}

func TestSearch(t *testing.T) {
	bst := ds.BST{}
	bst.Insert(5)
	bst.Insert(4)
	bst.Insert(7)
	bst.Insert(6)
	bst.Insert(8)

	if bst.Search(7).Val != 7 {
		t.Error(TEST_FAILED)
	}
	if bst.Search(1) != nil {
		t.Error(TEST_FAILED)
	}
}

func TestFloor(t *testing.T) {
	bst := ds.BST{}
	bst.Insert(10)
	bst.Insert(12)
	bst.Insert(14)
	bst.Insert(5)

	if bst.Floor(10).Val != 10 {
		t.Error(TEST_FAILED)
	}
	if bst.Floor(7).Val != 5 {
		t.Error(TEST_FAILED)
	}
	if bst.Floor(3) != nil {
		t.Error(TEST_FAILED)
	}
	if bst.Floor(13).Val != 12 {
		t.Error(TEST_FAILED)
	}
}

func TestCeil(t *testing.T) {
	bst := ds.BST{}
	bst.Insert(10)
	bst.Insert(12)
	bst.Insert(14)
	bst.Insert(5)

	if bst.Ceil(10).Val != 10 {
		t.Error(TEST_FAILED)
	}
	if bst.Ceil(3).Val != 5 {
		t.Error(TEST_FAILED)
	}
	if bst.Ceil(11).Val != 12 {
		t.Error(TEST_FAILED)
	}
	if bst.Ceil(15) != nil {
		t.Error(TEST_FAILED)
	}
}

func TestDelete(t *testing.T) {
	bst := ds.BST{}
	bst.Insert(10)
	bst.Delete(10)

	if bst.Root != nil {
		t.Error(TEST_FAILED)
	}

	bst.Insert(10)
	bst.Insert(9)
	bst.Delete(9)

	if bst.Root.Left != nil {
		t.Error(TEST_FAILED)
	}

	bst.Insert(9)
	bst.Insert(20)
	bst.Insert(14)
	bst.Insert(15)
	bst.Insert(30)
	bst.Delete(10)

	if bst.Root.Val != 14 {
		t.Error(TEST_FAILED)
	}
}

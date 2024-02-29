package quickunion_test

import (
	"testing"

	qu "github.com/andreev1024/go-datastructures/unionfind/quickunion"
)

const TEST_FAILED = "Test failed"

func TestQuickUnionFind(t *testing.T) {
	uf := qu.CreateUnionFind(10)
	// 1-2-5-6-7 3-8-9 4
	uf.Union(1, 2)
	uf.Union(2, 5)
	uf.Union(5, 6)
	uf.Union(6, 7)
	uf.Union(3, 8)
	uf.Union(8, 9)

	if !uf.Connected(1, 5) {
		t.Error(TEST_FAILED)
	}

	if !uf.Connected(5, 7) {
		t.Error(TEST_FAILED)
	}

	if uf.Connected(4, 9) {
		t.Error(TEST_FAILED)
	}

	uf.Union(9, 4)

	if !uf.Connected(4, 9) {
		t.Error(TEST_FAILED)
	}
}

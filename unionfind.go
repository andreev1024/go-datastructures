package datastructures

//https://leetcode.com/explore/learn/card/graph/618/disjoint-set/3843/

// Optimized “disjoint set” with Path Compression and Union by Rank
type UnionFind struct {
	root []int // array value - root vertex, array index - vertex
	rank []int
}

func (this *UnionFind) Find(x int) int {
	if x == this.root[x] {
		return x
	}

	this.root[x] = this.Find(this.root[x])
	return this.root[x]
}

func (this *UnionFind) Union(x, y int) {
	rootX := this.Find(x)
	rootY := this.Find(y)
	if rootX == rootY {
		return
	}
	if this.rank[rootX] > this.rank[rootY] {
		this.root[rootY] = rootX
	} else if this.rank[rootX] < this.rank[rootY] {
		this.root[rootX] = rootY
	} else {
		this.root[rootY] = rootX
		this.rank[rootX]++
	}
}

func (this *UnionFind) Connected(x, y int) bool {
	return this.Find(x) == this.Find(y)
}

func CreateUnionFind(size int) UnionFind {
	uf := UnionFind{make([]int, size), make([]int, size)}
	for i := 0; i < size; i++ {
		uf.rank[i] = 1 // one!
		uf.root[i] = i
	}
	return uf
}

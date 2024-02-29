package quickunion

//https://leetcode.com/explore/learn/card/graph/618/disjoint-set/3879/

type UnionFind struct {
	root []int
}

func (this *UnionFind) Find(x int) int {
	for x != this.root[x] {
		x = this.root[x]
	}
	return x
}

func (this *UnionFind) Union(x, y int) {
	rootX := this.Find(x)
	rootY := this.Find(y)
	if rootX != rootY {
		this.root[rootY] = rootX
	}
}

func (this *UnionFind) Connected(x, y int) bool {
	return this.Find(x) == this.Find(y)
}

func CreateUnionFind(size int) UnionFind {
	uf := UnionFind{make([]int, size)}
	for i := 0; i < size; i++ {
		uf.root[i] = i
	}
	return uf
}

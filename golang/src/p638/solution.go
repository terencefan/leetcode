package main

type Package struct {
	items []int
	price int
}

type Solution struct {
	packages []Package
	prices   []int
	max      int
}

func capable(remain []int, pkg Package) bool {
	for i, c := range pkg.items {
		if c > remain[i] {
			return false
		}
	}
	return true
}

func (self *Solution) dfs(remain []int, current int) {
	if current > self.max {
		return
	}
	for _, pkg := range self.packages {
		if !capable(remain, pkg) {
			continue
		}
		for i, item := range pkg.items {
			remain[i] -= item
		}
		self.dfs(remain, current+pkg.price)
		for i, item := range pkg.items {
			remain[i] += item
		}
	}
	for i, c := range remain {
		current += self.prices[i] * c
	}
	if current < self.max {
		self.max = current
	}
}

func ShoppingOffers(price []int, special [][]int, needs []int) int {

	var s = Solution{
		make([]Package, len(special)),
		price,
		99999,
	}

	for i, l := range special {
		s.packages[i] = Package{
			l[:len(l)-1],
			l[len(l)-1],
		}
	}
	s.dfs(needs, 0)
	return s.max
}

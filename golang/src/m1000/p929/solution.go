package main

import (
	"fmt"
)

func parse(email string) string {
	i, bs := 0, make([]byte, 0)

	skip := false
	for ; i < len(email); i++ {
		if email[i] == '@' {
			break
		} else if skip || email[i] == '.' {
			continue
		} else if email[i] == '+' {
			skip = true
		} else {
			bs = append(bs, email[i])
		}
	}

	for ; i < len(email); i++ {
		bs = append(bs, email[i])
	}
	return string(bs)
}

func numUniqueEmails(emails []string) int {
	var m = make(map[string]bool)
	for _, email := range emails {
		m[parse(email)] = true
	}
	fmt.Println(m)
	return len(m)
}

func main() {
	emails := []string{
		"j+ezsorqlmc@rfpycgjuf.com",
		"j+k+ri+rigt.ad@rfpycgjuf.com",
		"hfmeqzk.isx+i@fiidmdrsu.com",
		"j+p+h+d+d+p+z.j.k.g@rfpycgjuf.com",
		"zygekdy.l.w+s@snh.owpyu.com",
		"j+m.l+ia+qdgv+w@rfpycgjuf.com",
		"r+hwbjwefkp@wcjaki.n.com",
		"zygekdy.l.w+d@snh.owpyu.com",
		"bizdm+sz.f.a.k@il.cjjlp.com",
		"hfmeqzk.isx+t@fiidmdrsu.com",
		"hfmeqzk.isx+i@fiidmdrsu.com",
		"bizdm+f+j+m.l.l@il.cjjlp.com",
		"zygekdy.l.w+g@snh.owpyu.com",
		"r+evgvxmksf@wcjaki.n.com",
		"hfmeqzk.isx+h@fiidmdrsu.com",
		"bizdm+lov+cy@il.cjjlp.com",
		"hfmeqzk.isx+o@fiidmdrsu.com",
		"bizdm+hs+qao@il.cjjlp.com",
		"r+v+z+rcuznrj@wcjaki.n.com",
		"j+r.kn+h.w.a.h+bh@rfpycgjuf.com",
		"hfmeqzk.isx+t@fiidmdrsu.com",
		"hfmeqzk.isx+a@fiidmdrsu.com",
		"zygekdy.l.w+o@snh.owpyu.com",
		"zygekdy.l.w+i@snh.owpyu.com",
		"r+vy.u.y+f.er+aw@wcjaki.n.com",
		"zygekdy.l.w+c@snh.owpyu.com",
		"bizdm+wztzg@il.cjjlp.com",
		"j+h.fwbsr+chp@rfpycgjuf.com",
		"zygekdy.l.w+s@snh.owpyu.com",
		"zygekdy.l.w+d@snh.owpyu.com",
		"bizdm+qq.o.q+p@il.cjjlp.com",
		"zygekdy.l.w+o@snh.owpyu.com",
		"j+c+zqbq+h.dyt@rfpycgjuf.com",
		"r+hl.b.i+fz.hz.t@wcjaki.n.com",
		"r+cbabpf.k+w+e@wcjaki.n.com",
	}
	fmt.Println(numUniqueEmails(emails))
}

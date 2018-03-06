/*
题目描述
查找和排序

题目：输入任意（用户，成绩）序列，可以获得成绩从高到低或从低到高的排列,相同成绩
      都按先录入排列在前的规则处理。

   例示：
jack 70
peter 96
Tom 70
smith 67

   从高到低  成绩
   peter     96
   jack      70
   Tom       70
   smith     67

   从低到高

   smith     67

   Tom       70
   jack      70
   peter     96

输入描述:
输入多行，先输入要排序的人的个数，然后输入排序方法0（降序）或者1（升序）再分别输入他们的名字和成绩，以一个空格隔开

输出描述:
按照指定方式输出名字和成绩，名字和成绩之间以一个空格隔开

示例1
输入
3
0
fang 90
yang 50
ning 70
输出
fang 90
ning 70
yang 50
*/

package main

import (
	"fmt"
)

type user struct {
	name  string
	score int
	next  *user
}

func (u *user) Output() {
	cu := u
	for cu != nil {
		fmt.Printf("%s %d\n", cu.name, cu.score)
		cu = cu.next
	}
}

type tree struct {
	l, r *tree
	u    *user
}

func (t *tree) add(u *user) bool {
	if t.u == nil {
		t.u = u
		return true
	}
	if t.u.score > u.score {
		if t.l == nil {
			t.l = &tree{u: u}
			return true
		}
		return t.l.add(u)
	}
	if t.r == nil {
		t.r = &tree{u: u}
		return true
	}
	return t.r.add(u)
}

func (t *tree) minfirst() {
	if t.l != nil {
		t.l.minfirst()
	}
	t.u.Output()
	if t.r != nil {
		t.r.minfirst()
	}
}

func (t *tree) maxfirst() {
	if t.r != nil {
		t.r.maxfirst()
	}
	t.u.Output()
	if t.l != nil {
		t.l.maxfirst()
	}
}

func main() {
	var n, s int
	for k, err := fmt.Scanf("%d\n%d", &n, &s); k == 2 && err == nil; {
		// k, err := fmt.Scanf("%d\n%d", &n, &s)
		if k <= 0 || err != nil {
			fmt.Println(k, err)
			return
		}
		us := make([]*user, n)
		root := new(tree)
		for i := 0; i < n; i++ {
			u := &user{}
			fmt.Scanf("%s %d", &(u.name), &(u.score))
			us[i] = u
			root.add(u)
		}
		// fmt.Println("n:", n)

		if s == 1 {
			root.minfirst()
		} else {
			root.maxfirst()
		}
	}
}

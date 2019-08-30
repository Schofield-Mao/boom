package heap

import (
	"testing"
	"strconv"
	"sort"
)

func TestInsert(t *testing.T){
	
	Tests := []struct{
		m map[string]int
		a []string
		ans []string
	}{
		{
			m : map[string]int{
				"a": 0,
				"b": 1,
				"c": 2,
				"d": 3,
				"e": 4,
				"f": 5,
				"g": 6,
				"h": 7,
				"i": 8,
				"j": 9,
				"TEST":-1,
			},
			a :[]string{"a","b","c","d","e","f","g","h","i","j"},
			ans : []string{"TEST","a","c","d","b","f","g","h","i","j","e"},
		},
		{
			m : map[string]int{
				"a": 0,
				"b": 1,
				"c": 2,
				"d": 3,
				"e": 4,
				"f": 5,
				"g": 6,
				"h": 7,
				"i": 8,
				"j": 9,
				"TEST":10,
			},
			a : []string{"a","b","c","d","e","f","g","h","i","j"},
			ans : []string{"a","b","c","d","e","f","g","h","i","j","TEST"},
		},
		{
			m : map[string]int{
				"a": 0,
				"b": 1,
				"c": 2,
				"d": 3,
				"e": 4,
				"f": 5,
				"g": 6,
				"h": 7,
				"i": 8,
				"j": 9,
				"TEST":2,
			},
			a : []string{"a","b","c","d","e","f","g","h","i","j"},
			ans : []string{"a","b","c","d","TEST","f","g","h","i","j","e"},
		},
	}
	
	for _,test := range Tests{
		insert("TEST",&test.a ,test.m)
		for idx,v := range test.ans{
			if v != test.a[idx]{
				t.Error(test.ans)
				t.Error(test.a)
				t.Error("error")
				return
			}
		}
	}
}

func TestRemove(t *testing.T){
	m := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
		"i": 8,
		"j": 9,
	}
	a := []string{"a","b","c","d","e","f","g","h","i","j"}
	ans := []string{"b","d","c","h","e","f","g","j","i"}
	remove(&a,m)
	for idx,v := range ans{
		if v != a[idx]{
			t.Error(ans)
			t.Error(a)
			t.Error("error")
			return
		}
	}
}

func TestGetTopK(t *testing.T){
	m := make(map[string]int)
	res := make([]string,0)
	for i:=0;i<50;i++{
		res = append(res,strconv.Itoa(i))
		m[strconv.Itoa(i)] = 100
	}
	for i:=51;i<10000;i++{
		m[strconv.Itoa(i)] = 30
	}
	for i:=10000;i<10050;i++{
		res = append(res,strconv.Itoa(i))
		m[strconv.Itoa(i)] = 100
	}

	a := GetTopK(m,100)
	sort.Sort(sort.StringSlice(a))
	sort.Sort(sort.StringSlice(res))
	for i,v := range res{
		if v != a[i]{
			t.Error("error")
			break
		}
	}
}


package iutil

import (
	"sort"
	"testing"
)

func TestHanString(t *testing.T) {
	han := "一二三四五六七"
	for _, s := range han {
		t.Log(string(s)) // 一 二 三 四 五 六 七
	}
	t.Log([]rune(han)) // [19968 20108 19977 22235 20116 20845 19971]
}

func TestHanLess(t *testing.T) {
	t.Log(HanLess("啊", "阿")) // true
	t.Log(HanLess("阿", "啊")) // false
}

func TestSort(t *testing.T) {
	han := []string{"一", "一曲", "一破", "已排", "二", "三", "四", "五", "六", "七"}
	sort.Slice(han, func(i, j int) bool {
		return HanLess(han[i], han[j])
	})
	t.Log(han) // [二 六 七 三 四 五 一 一破 一曲 已排]
}

func TestMaxInt(t *testing.T) {
	t.Log(MaxInt(1, 4)) // 4
	t.Log(MaxInt(9, 4)) // 9
}

func TestMinInt(t *testing.T) {
	t.Log(MinInt(1, 4)) // 1
	t.Log(MinInt(9, 4)) // 4
}

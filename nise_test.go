package main

import (
	"testing"
)

func TestFilter(t *testing.T) {
	cases := []struct {
		Sentence string
		Expected string
	}{
		{
			Sentence: "私はお酒が飲みたい",
			Expected: "我酒飲希望",
		},
		{
			Sentence: "私は本日定時退社します",
			Expected: "我本日定時退社也",
		},
		{
			Sentence: "私は理解しました",
			Expected: "我理解了",
		},
		{
			Sentence: "私は昨日、日本の料理を食べました",
			Expected: "我昨日、日本的料理食了",
		},
		{
			Sentence: "今日は10時に仕事が終わります",
			Expected: "今日10時仕事終了",
		},
		{
			Sentence: "あなたは何時に終わりますか？",
			Expected: "你何時終了如何?", // 君何時終了如何?
		},
	}
	for i, v := range cases {
		if got := Filter(v.Sentence); got != v.Expected {
			t.Errorf("expectded %s, got %s (%d)", v.Expected, got, i)
		}
	}
}

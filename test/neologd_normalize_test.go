package test

import (
	"fmt"
	"github.com/shigenobu/mysql_neologd_normalize/func"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalize(t *testing.T) {
	inputs := [] string{
		"０１２３４５６７８９",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ",
		"ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ",
		"！”＃＄％＆’（）＊＋，－．／：；＜＞？＠［￥］＾＿｀｛｜｝",
		"＝。、・「」",
		"ﾊﾝｶｸ",
		"o₋o",
		"majika━",
		"わ〰い",
		"スーパーーーー",
		"!#",
		"ゼンカク　スペース",
		"お             お",
		"      おお",
		"おお      ",
		"検索 エンジン 自作 入門 を 買い ました!!!",
		"アルゴリズム C",
		"　　　ＰＲＭＬ　　副　読　本　　　",
		"Coding the Matrix",
		"南アルプスの　天然水　Ｓｐａｒｋｉｎｇ　Ｌｅｍｏｎ　レモン一絞り",
		"南アルプスの　天然水-　Ｓｐａｒｋｉｎｇ*　Ｌｅｍｏｎ+　レモン一絞り",
	}

	expecteds := [] string{
		"0123456789",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcdefghijklmnopqrstuvwxyz",
		"!\"#$%&'()*+,-./:;<>?@[¥]^_`{|}",
		"＝。、・「」",
		"ハンカク",
		"o-o",
		"majikaー",
		"わい",
		"スーパー",
		"!#",
		"ゼンカクスペース",
		"おお",
		"おお",
		"おお",
		"検索エンジン自作入門を買いました!!!",
		"アルゴリズムC",
		"PRML副読本",
		"Coding the Matrix",
		"南アルプスの天然水Sparking Lemonレモン一絞り",
		"南アルプスの天然水-Sparking*Lemon+レモン一絞り",
	}

	for i := 0; i < len(inputs); i++ {
		in := _func.NeoLogdNormalize(inputs[i])
		ex := expecteds[i]
		fmt.Println(in, ":", ex)
		assert.Equal(t, ex, in)
	}
}

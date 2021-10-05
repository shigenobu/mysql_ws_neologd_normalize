/*
 neologd normalize.
 */
package _func

import (
	"fmt"
	"golang.org/x/text/unicode/norm"
	"regexp"
	"strings"
)

var reZenNumberAlphabet = regexp.MustCompile(`([０-９Ａ-Ｚａ-ｚ｡-ﾟ]+)`)

var reZenMark = regexp.MustCompile(`([！”＃＄％＆’（）＊＋，－．／：；＜＞？＠［￥］＾＿｀｛｜｝〜]+)`)

var blocks = "\u4E00-\u9FFF\u3040-\u309F\u30A0-\u30FF\u3000-\u303F\uFF00-\uFFEF"

var basicLatin = "\u0000-\u007F"

var reBlocksBlocks = regexp.MustCompile(fmt.Sprintf("([%s]) ([%s])", blocks, blocks))

var reBlocksBasicLatin = regexp.MustCompile(fmt.Sprintf("([%s]) ([%s])", blocks, basicLatin))

var reBasicLatinBlocks = regexp.MustCompile(fmt.Sprintf("([%s]) ([%s])", basicLatin, blocks))

var reContLine1 = regexp.MustCompile(`[˗֊‐‑‒–⁃⁻₋−]+`)

var reContLine2 = regexp.MustCompile(`[﹣－ｰ—―─━ー]+`)

var reContLine3 = regexp.MustCompile(`[~∼∾〜〰～]`)

var reSpace = regexp.MustCompile(`[ 　]+`)

var hanList = strings.Split("!\"#$%&'()*+,-./:;<=>?@[¥]^_`{|}~｡､･｢｣", "")

var zenList = strings.Split("！”＃＄％＆’（）＊＋，－．／：；＜＝＞？＠［￥］＾＿｀｛｜｝〜。、・「」", "")

var reZen1 = regexp.MustCompile("[’]")
var reZen2 = regexp.MustCompile("[”]")

// 文字間のスペースを取り除く
func removeSpaceBetween(re* regexp.Regexp, s string) string {
	for re.MatchString(s) {
		s = re.ReplaceAllString(s, "$1$2")
	}
	return s
}

// 余分なスペースを取り除く
func removeExtraSpaces(s string) string {
	s = reSpace.ReplaceAllString(s, " ")

	s = removeSpaceBetween(reBlocksBlocks, s)
	s = removeSpaceBetween(reBlocksBasicLatin, s)
	s = removeSpaceBetween(reBasicLatinBlocks, s)

	return s
}

// UnicodeのNFKCで標準化する
func normNkfc(c string, re* regexp.Regexp) string {
	if re.MatchString(c) {
		c = norm.NFKC.String(c)
	}
	return c
}

// Unicodeの標準化を行う
func unicodeNormalize(re* regexp.Regexp, s string) string {
	matches := re.FindAllStringSubmatch(s, -1)
	for _, m := range matches {
		if len(m) > 0 {
			r := normNkfc(m[0], re)
			s = strings.Replace(s, m[0], r, -1)
		}
	}
	s = strings.Replace(s, "－", "-", -1)
	return s
}

// neologdのアルゴリズムで標準化する
// https://github.com/neologd/mecab-ipadic-neologd/wiki/Regexp.ja
func NeoLogdNormalize(sentence string) string {
	// 両端の半角スペースを取り除く
	sentence = strings.Trim(sentence, " 　")

	// タブ、改行などは削除
	sentence = strings.Replace(sentence, "\r", "", -1)
	sentence = strings.Replace(sentence, "\n", "", -1)
	sentence = strings.Replace(sentence, "\t", "", -1)
	sentence = strings.Replace(sentence, "\v", "", -1)

	// unicode正規化をかける
	sentence = unicodeNormalize(reZenNumberAlphabet, sentence)

	// 連続文字置換
	sentence = reContLine1.ReplaceAllString(sentence, "-")
	sentence = reContLine2.ReplaceAllString(sentence, "ー")
	sentence = reContLine3.ReplaceAllString(sentence, "")

	// 記号変換
	for i := 0; i < len(hanList); i++ {
		sentence = strings.Replace(sentence, string(hanList[i]), string(zenList[i]), -1)
	}

	// 空白除去
	sentence = removeExtraSpaces(sentence)

	// 再度正規化
	sentence = unicodeNormalize(reZenMark, sentence)
	sentence = reZen1.ReplaceAllString(sentence, "'")
	sentence = reZen2.ReplaceAllString(sentence, "\"")

	return sentence
}

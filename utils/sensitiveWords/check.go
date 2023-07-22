package sensitiveWord

import (
	"io/ioutil"
	"os"
	"project/consts"
	"project/types"
	"strings"

	"go.uber.org/zap"
)

// SensitiveMap 使用前缀树实现敏感词过滤
type SensitiveMap struct {
	sensitiveNode map[string]interface{}
	isEnd         bool
}

type Target struct {
	Indexes []int
	Len     int
}

var s *SensitiveMap

// getMap 将自己的词库放入/static/dictionary下，放入下列切片中！！！！
func getMap() *SensitiveMap {
	if s == nil {
		var Sen []string
		Sen = append(Sen, consts.OtherSen)
		s = InitDictionary(s, Sen)
	}

	return s
}

// CheckSensitiveWord 判断是否有敏感词
func CheckSensitiveWord(content string) []*types.SensitiveWord {
	var res []*types.SensitiveWord
	sensitiveMap := getMap()
	target := sensitiveMap.FindAllSensitive(content)
	for k, v := range target {
		t := &types.SensitiveWord{
			Word:    k,
			Indexes: v.Indexes,
			Length:  v.Len,
		}
		res = append(res, t)
	}

	return res
}

// FindAllSensitive 查找所有的敏感词
func (s *SensitiveMap) FindAllSensitive(text string) map[string]*Target {
	content := []rune(text)
	contentLength := len(content)
	result := false
	ta := make(map[string]*Target)

	for index := range content {
		sMapTmp := s
		target := ""
		in := index
		result = false

		for {
			wo := string(content[in])
			target += wo
			if _, ok := sMapTmp.sensitiveNode[wo]; ok {
				if sMapTmp.sensitiveNode[wo].(*SensitiveMap).isEnd {
					result = true
					break
				}
				if in == contentLength-1 {
					break
				}
				sMapTmp = sMapTmp.sensitiveNode[wo].(*SensitiveMap)
				in++
			} else {
				break
			}
		}
		if result {
			if _, targetInTa := ta[target]; targetInTa {
				ta[target].Indexes = append(ta[target].Indexes, index)
			} else {
				ta[target] = &Target{
					Indexes: []int{index},
					Len:     len([]rune(target)),
				}
			}
		}
	}

	return ta
}

// InitDictionary 初始化字典，构造前缀树
func InitDictionary(s *SensitiveMap, dictionary []string) *SensitiveMap {
	// 初始化字典树
	s = initSensitiveMap()
	var dictionaryContent []string
	for i := 0; i < len(dictionary); i++ {
		dictionaryContentTmp := ReadDictionary(dictionary[i])
		// TODO:将所有词拿到
		dictionaryContent = append(dictionaryContent, dictionaryContentTmp...)
	}
	for _, words := range dictionaryContent {
		sMapTmp := s
		// 将每一个词转换为一个rune数组，不光英文、中文
		w := []rune(words)
		wordsLen := len(w)
		for i := 0; i < wordsLen; i++ {
			t := string(w[i])
			isEnd := false
			if i == (wordsLen - 1) {
				isEnd = true
			}
			func(tx string) {
				if _, ok := sMapTmp.sensitiveNode[tx]; !ok {
					sMapTemp := new(SensitiveMap)
					sMapTemp.sensitiveNode = make(map[string]interface{})
					sMapTemp.isEnd = isEnd
					sMapTmp.sensitiveNode[tx] = sMapTemp
				}
				sMapTmp = sMapTmp.sensitiveNode[tx].(*SensitiveMap)
				sMapTmp.isEnd = isEnd
			}(t)
		}
	}

	return s
}

// initSensitiveMap 初始化map
func initSensitiveMap() *SensitiveMap {
	return &SensitiveMap{
		sensitiveNode: make(map[string]interface{}),
		isEnd:         false,
	}
}

// ReadDictionary 将词库读取出来
func ReadDictionary(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		zap.L().Error("read dictionary file failed:", zap.Error(err))
		return nil
	}

	defer file.Close()

	str, err := ioutil.ReadAll(file)
	dictionary := strings.Fields(string(str))

	return dictionary
}

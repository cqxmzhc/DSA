package trie

type Vertext struct {
	//标记一个word的结尾
	end bool

	//子节点
	sons map[string]*Vertext
}

//New 生成新的trie树
func New() *Vertext {
	return &Vertext{
		sons: make(map[string]*Vertext),
	}
}

//AddWord 向树中添加word
func AddWord(v *Vertext, word string) {
	if len(word) > 0 {
		//记录当前位置
		current := v

		next := v.sons
		for _, w := range word {
			v, ok := next[string(w)]
			if !ok {
				//新建节点
				current = &Vertext{
					sons: make(map[string]*Vertext),
				}
				next[string(w)] = current

				next = current.sons
			} else {
				current = v
				next = v.sons
			}
		}

		//标记一个单词的结尾
		current.end = true
	}
}

//DeleteWord 删除树中的某个word
func DeleteWord(v *Vertext, word string) {
	current := v
	next := v.sons
	lastIdx := len(word) - 1

	//在树中找到指定word并定位到最后一个字符，将word结束位标志为false
	for i, w := range word {
		if v, ok := next[string(w)]; ok {
			current = v
			next = v.sons

			if i == lastIdx {
				current.end = false
			}
		}
	}
}

//FindPrefix  找到以prefix为前缀的所有word
func FindPrefix(v *Vertext, prefix string) []string {
	res := []string{}

	current := v
	_ = current
	next := v.sons
	lastIdx := len(prefix) - 1

	for i, w := range prefix {
		if v, ok := next[string(w)]; ok {
			current = v
			next = v.sons

			if i == lastIdx {
				//遍历所有子节点
				res = append(res, getWordsStartWithKey(v, prefix)...)

				return res
			}
		}
	}

	return nil
}

//getWordsStartWithKey 递归获取指定节点下的所有words(结尾字符end标记为true)
func getWordsStartWithKey(v *Vertext, key string) []string {
	if v == nil {
		return nil
	}

	res := []string{}
	if v.end == true {
		res = append(res, key)
	}
	//遍历
	for k, v := range v.sons {
		res = append(res, getWordsStartWithKey(v, key+k)...)
	}

	return res
}

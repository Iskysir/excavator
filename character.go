package excavator

type CharacterFunc func(character *Character) error

//Character 字符
type Character struct {
	//URL            string //汉字地址
	Character      string //字符
	Pinyin         string //拼音
	Radical        string //部首
	RadicalStrokes string //部首笔画
	TotalStrokes   string //总笔画
	KangxiStrokes  string //康熙笔画数
	Phonetic       string //注音
	Folk
	Structure
	Explain
	Rhyme
	Index
}

//Folk民俗参考
type Folk struct {
	CommonlyCharacters   string //是否为常用字
	NameScience          string //姓名学
	FiveElementCharacter string //汉字五行
	GodBadMoral          string //吉凶寓意
}

//Structure 字形结构
type Structure struct {
	DecompositionSearch string //首尾分解查字
	StrokeNumber        string //笔顺编号
	StrokeReadWrite     string //笔顺读写
}

//Explain 康熙字典解释
type Explain struct {
	Intro  string //简介
	Detail string //详情
}

//Rhyme 音韵参考
type Rhyme struct {
	GuangYun  string //广　韵
	Mandarin  string //国　语
	Cantonese string //粤　语
}

//Index 索引参考
type Index struct {
	AncientWrite      string //古文字诂林
	HometownTrain     string //故训彙纂
	Explain           string //说文解字
	KangxiDictionary  string //康熙字典
	ChineseDictionary string //汉语字典
	Cihai             string //辞　海  
}

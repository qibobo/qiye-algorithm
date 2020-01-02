package others_test

import (
	"fmt"
	"github.com/qibobo/qiye-algorithm/others"
	"testing"
)

var strs1 []string = []string{
	"ren sheng zi shi you qing chi ci hen bu guan feng yu yue",
}

// 1.人生自是有情痴，此恨不关风与月。____欧阳修《玉楼春·尊前拟把归期说》 2.月落乌啼霜满天，江枫渔火对愁眠。——张继《枫桥夜泊》 3.眠沙卧水自成群，曲岸残阳极浦云。——李商隐《题鹅》 4.云想衣裳花想容， 春风拂槛露华浓。——李白《清平调》 5.浓艳初开小药栏，人人惆怅出长安。——张祜《杭州开元寺牡丹》 6.安得广厦千万间，大庇天下寒士俱欢颜！——杜甫《茅屋为秋风所破歌》 7.颜子将才应四科，料量时辈更谁过。——李涉《送颜觉赴举》 8.过尽千帆皆不是，斜晖脉脉水悠悠，肠断白苹洲。——温庭筠《梦江南·千万恨》 9.洲长春色遍，汉广夕阳迟。——刘长卿《赠别卢司直之闽中》 10.迟迟不见怜弓箭，惆怅秋鸿敢近飞。——刘商《行营病中》

// 11.飞鸟去不穷，连山复秋色。——王维《辋川集·华子冈》 12.色空荣落处，香醉往来人。——刘长卿《题灵祐上人法华院木兰花》 13.人事有代谢，往来成古今。——孟浩然《与诸子登岘山》 14.今日相如身在此，不知客右坐何人。——白居易《雪中寄令狐相公，兼呈梦得》 15.人生若只如初见， 何事秋风悲画扇。——纳兰性德《木兰词·拟古决绝词柬友》 16.扇风淅沥簟流离，万里南风滞所思。——李商隐《到秋》 17.思悠悠，恨悠悠，恨到归时方始休。——白居易《长相思·汴水流》 18.休垂绝徼千行泪，共泛清湘一叶舟。——韩愈《湘中酬张十一功曹》 19.舟行十里画屏上，身在西山红雨中。——陆游《连日至梅僊坞及花泾观桃花抵暮乃归》 20.中庭地白树栖鸦，冷露无声湿桂花。——王建《十五夜望月寄杜郎中》

// 21.花里可怜池上景，几重墙壁贮春风——张籍《九华观看花》 22.风急天高猿啸哀，渚清沙白鸟飞回。——杜甫《登高》 23.回首看花花欲尽，可怜寥落送春心。——高骈《池上送春》 24.心随长风去，吹散万里云。——李白《赠何七判官昌浩》 25.云光岚彩四面合，柔柔垂柳十馀家。——杜牧《商山麻涧》 26.家住水东西，浣纱明月下。——王维《白石滩》 27.下窥指高鸟，俯听闻惊风。——岑参《与高适薛据登慈恩寺浮图》 28.风絮飘残已化萍，泥莲刚倩藕丝萦。——纳兰性德《山花子·风絮飘残已化萍》 29.萦回谢女题诗笔，点缀陶公漉酒巾。——刘禹锡《柳絮》 30.巾栉不可见，枕席空馀香。——徐彦伯《相和歌辞·班婕妤》
var strs2 []string = []string{
	"ren sheng zi shi you qing chi ci hen bu guan feng yu yue",
	"yue luo wu di shuang man tian jian feng yu huo dui chou mian",
	"mian wo shui zi cheng qun qu an can yang ji pu yun",
	"yun xiang yi shang hua xiang rong chun feng fu jian lu hua nong",
	"nong yan chu kai xiao yao lan ren ren chou chang chu chang an",
	"an de guang sha qian wan jian da bi tian xia han shi ju huan yan",
	"yan zi jiang cai ying si ke liao liang shi bei geng shei guo",
	"guo jin qian fan jie bu shi xie hui mo mo shui you you chang duan bai ping zhou",
	"zhou chang chun se bian han guang xi yang chi",
	"",
}
var strs3 []string = []string{
	"guo jin qian fan jie bu shi xie hui mo mo shui you you chang duan bai ping zhou",
	"nong yan chu kai xiao yao lan ren ren chou chang chu chang an",

	"ren sheng zi shi you qing chi ci hen bu guan feng yu yue",
	"yue luo wu di shuang man tian jian feng yu huo dui chou mian",
	"an de guang sha qian wan jian da bi tian xia han shi ju huan yan",

	"mian wo shui zi cheng qun qu an can yang ji pu yun",

	"zhou chang chun se bian han guang xi yang chi",

	"",
	"yan zi jiang cai ying si ke liao liang shi bei geng shei guo",

}

func TestMatch(t *testing.T) {
	happy := [][2]string{
		{"A B", "B V"},
		{"A", "A B"},
	}
	bad := [][2]string{
		{"", ""},
		{"A", "A"},
		{"A", ""},
		{"", "B"},
		{"A", "B"},
	}
	for _, item := range happy {
		if !others.Match(item[0], item[1]) {
			t.Error(fmt.Sprintf("%s should match %s", item[0], item[1]))
		}
	}
	for _, item := range bad {
		if others.Match(item[0], item[1]) {
			t.Error(fmt.Sprintf("%s should not match %s", item[0], item[1]))
		}
	}

}

func TestFindLongestPathInGraph(t *testing.T) {
	length := others.FindLongestPathInGraph(strs1)
	if length != 1 {
		t.Error(fmt.Printf("%d is not 1\n", length))
	}

	length = others.FindLongestPathInGraph(strs2)
	if length != 9 {
		t.Error(fmt.Printf("%d is not 9\n", length))
	}

	length = others.FindLongestPathInGraph(strs3)
	if length != 5 {
		t.Error(fmt.Printf("%d is not 3\n", length))
	}
}

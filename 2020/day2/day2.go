package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	input := "4-5 r: rrrjr\n9-10 x: pxcbpxxwkqjttx\n8-13 b: rjbbbbvgrbrfjx\n3-5 d: dtddsdddddsddddddwd\n3-11 q: qbqsqqzqqxkmbqx\n3-4 v: vgvhcvxlbfcwg\n1-7 t: rtctmtt\n8-11 s: ksssswsssstssssss\n2-6 v: vxvvvlvv\n16-18 c: ptgprcccvqvbfcnckc\n18-20 w: wgtnwgnwvzwqhqtwvwrw\n14-18 k: kkkkkkxkkkkkknkkkl\n3-6 f: pwfgqlbbrfvwf\n6-10 h: pqlfbhcnglgvhdgddn\n2-13 k: bkfdvztsdkkkwm\n8-12 d: dddddddddddd\n1-10 v: lvvvvvvvvvvnvvvvvv\n2-3 d: dfdz\n3-10 w: vwwwwwnwwwwwxww\n2-7 c: ccwkpggxd\n4-5 g: kgqqg\n1-14 q: bcqqqqqqdqqqqbqqcq\n1-10 j: jjjjjjjfjjjjhjj\n1-4 x: bkjghxrsb\n2-15 v: vvvpvvvpvvrvvvvvvv\n1-6 q: pqlqqn\n2-6 k: skdkqk\n14-16 g: sggbggggrgbmgtgg\n8-9 f: ffffnkffk\n1-2 x: tjxxl\n6-8 t: tzdtvttkhctk\n1-4 z: zzzkzzzzzxzzz\n1-2 h: hthhjhhhqq\n2-4 d: ddpkzvxbjnxhw\n1-4 k: kkhq\n8-9 s: sssssssps\n5-11 b: bbbbbbbbbrbbbbbb\n2-4 g: crrj\n4-5 j: djjjqzmxttjg\n1-4 b: bbbbb\n3-4 z: zrqx\n2-12 l: vvggvdqqcjplqzkczgj\n15-16 l: sgclwcllcvltvthl\n2-3 c: tbclrfclrhxvtlw\n8-11 n: nhnnnwlqnnznnnnnn\n7-15 c: jlcxswgncffcpnbdrhlx\n6-14 k: kkkkkgkkkkkkkkk\n8-11 r: rrrrrrrsrrrkrrr\n4-6 m: fmhmtmm\n10-12 p: ppcpjppppppg\n2-5 m: zbtmmmmlmmmhvblmmnx\n4-5 q: hqwhr\n2-5 r: rzrrwrr\n9-10 k: kkkkkkkkchk\n3-5 p: ppptkpn\n2-3 f: ffxfzwd\n6-14 p: ppppppppppppppppp\n1-7 f: hffffffkffff\n2-13 h: hhhhhhhhhmmhvhhhjhhh\n8-10 z: zzzzkzzbdzz\n8-16 w: wwwwwwwwwwwwvwwwwww\n3-10 n: qqpnnnlnwznctsbjnn\n2-4 m: rmfw\n2-5 z: jzfsqt\n14-15 n: nnnnfnfndpnnnnn\n7-9 d: drvdpddpddd\n8-12 j: kjjjjjjqjjjjjj\n3-11 h: hxhhhhllhthwhjhxhhhz\n3-5 w: ndwww\n5-10 r: qrrrrrrrrrrr\n9-10 k: kkkdlkkkkwknkg\n5-7 t: lstztlh\n1-2 j: jljk\n5-19 q: zkrqqwthdjgzqzpjwkq\n2-3 r: drrrpswzwfmrfcrss\n8-12 r: qrrrdrxfzdqb\n17-18 x: cxbdcxxxxxxkxxdxxq\n12-15 d: dvxdddddndzldlmc\n6-14 m: mmmmmsmmmmmmmmm\n2-4 s: lssf\n1-4 p: tpppp\n4-10 w: wwwwwwwwwpwww\n3-14 p: fshdffznzrppwpp\n1-8 l: hllllmlf\n1-18 m: mmmmmmmmmmmmmmmmtmm\n5-7 p: pfpqpppp\n7-15 q: qzkqvfxqqrvqqqcnxccc\n3-4 g: jlgggtk\n2-3 f: dfffffffffff\n2-3 j: jjjmjs\n8-10 t: tthtmxttftftvttntqkh\n1-13 x: xxxxxxxsxxxxxx\n2-3 d: fjdd\n13-17 b: nnxgrpkqgtnhjsnbcqhb\n12-18 x: lnqlnzqxxlxclzpwxx\n7-10 l: gxxlqllqmllglmls\n2-15 j: djfqjtlswxsdrvwcwkl\n7-9 n: hdblrjntz\n2-4 f: ffzff\n2-3 b: bvbb\n10-14 n: tglngsjnzlptknxcrbdn\n15-16 m: mmmksfmsbzmmjnqmm\n3-6 g: kgzmgk\n9-17 g: gcgrgxhgggmgrrdqtdg\n3-5 h: whcfh\n4-5 g: kpdgw\n2-10 d: dhdddddddld\n9-12 l: stsjflnqblqtl\n2-3 t: tttt\n18-19 t: tttctttttttttttttqj\n5-8 x: xxxdxxxxxmxhpxdpx\n5-8 h: hhhphhhkhtthhph\n1-7 l: wlllllnl\n4-7 x: xxxlxxsx\n14-15 r: rrrrdlvrrrrrdjvrrr\n11-18 c: hbhtbtczsbcqvrkxwc\n14-17 d: ndqzdrcftdkdddddv\n2-3 l: llgf\n7-8 z: hvqcstzrzz\n3-13 g: ggtggggggggggggggg\n11-12 r: dgnzdmsrqgrccf\n8-9 z: zqzzzzzzbzz\n7-9 h: hllhhhwhhzhrc\n7-9 g: ghgkgwrvg\n5-6 d: ddddknd\n4-5 s: qqlss\n1-2 q: jcqq\n9-10 v: vvvvvvvvxwvvvv\n4-9 p: zqjpvpttf\n16-20 t: tpttthqttcpttttftttk\n8-13 p: pppppppmppppppp\n1-5 p: hpppzp\n2-4 t: ktttt\n2-4 f: jfltfr\n2-6 t: gxtmnt\n2-8 v: vvjqpvvxvvfvzvv\n1-4 l: dqlsll\n1-3 b: zhqgqrbk\n7-9 n: nncfnnnnwnnn\n4-12 w: bqswnwqwlkhwbcc\n1-2 b: kbbbb\n4-11 x: xjxkxxxnxxxxxxxxxx\n2-10 b: bbbwnbknghbbbbbj\n9-11 h: hhhmghhshtqh\n8-9 t: tdctmtfpt\n8-19 b: bbbbbbbsbngbzhbpbbb\n11-12 k: dxkkhkknrkrkkkz\n8-10 b: fcbwbgbkbb\n1-5 h: qfdhn\n6-9 w: wwwvkpwwbwwwww\n1-10 c: ccccccccccccccccc\n13-14 k: kxzkfkkkkkkfvckkk\n2-4 p: wrrrpb\n7-9 s: spsssssvssssxzsqzssh\n4-5 f: fffgf\n13-15 n: clhngxdnnsnqlptnz\n10-11 f: fffffffffhlff\n9-13 k: kkkkkkkktkkkkk\n15-16 p: pxppppppgprslpln\n4-6 s: mscssms\n2-4 b: zbjjjz\n13-17 n: cthpnbmzwxbnnxngns\n14-18 z: zzzzzzzzzzzzzdzzzvz\n16-18 h: hhhhhhhhhhhhhhhvhhh\n1-6 s: xssfswlsssscdnshss\n7-10 v: cwrzfzvppzgcvkxs\n1-2 n: xnjnnrn\n8-9 b: rnbbbskbrhbbbfbz\n10-20 x: jljkxspllxqhjvknfxqx\n6-7 q: qxskpqq\n5-8 q: mwhqqqdfq\n11-15 l: lplldlllltlljll\n3-8 j: grsdjjzj\n2-8 p: pppppppjp\n3-10 m: xpfmmmvmnmrtjmfbkpb\n10-11 k: cfkbkndxkxgk\n11-12 k: kkkkkkkkkkkb\n10-11 t: ttttttttttnt\n10-11 w: wwwwvgswwwwwwwwww\n3-4 c: drkc\n8-11 f: ffgfdffjffnflnqfsff\n5-6 n: nnnnwtq\n10-11 s: dsswkmshsbsslwgsvss\n8-11 z: zzzzzzzlzzsz\n8-15 s: sssssssssnsssssssss\n9-18 d: ghzcrfrhflqdsvrphd\n11-16 q: fqqqkpfhxqqmgxcj\n4-6 b: qcbzcv\n2-5 q: qqlpn\n9-10 c: ccccbqgxcc\n11-14 v: nzvjhxvvnvbvvvvv\n1-4 q: xqgq\n5-12 t: ftdtmwmqkhthj\n5-6 w: bqlwwr\n11-12 t: vlfgktvfttrt\n4-7 n: nfnvtnjnnbnbnnn\n9-15 w: wwwwwwwwwwwwwwjwww\n17-18 k: kkkkkpkkkkkkkkkkkckk\n2-4 x: nxnx\n2-4 s: lbjm\n4-6 j: jjjqjzj\n3-4 v: lvwv\n5-8 x: xhxxxsgxnjfdb\n4-6 m: wmmlcx\n6-10 t: tttttdttttt\n15-18 d: dddddsdddddddddddldd\n14-16 r: crwxgrrlrrgsrrqc\n4-8 b: bbccgprbbqmvmjcbbfn\n8-10 d: dddddddddldd\n3-5 k: jkkzk\n19-20 q: bnqnchhwjjrnzzqcqqvq\n2-3 g: ghfggxgbg\n11-16 r: pgqjshvmjrzrrrxsr\n4-6 j: qjjjnrjzjjrjj\n6-10 n: qrwwvnngtnsgmn\n7-8 m: mmmmmmhhmm\n5-6 c: clfcnx\n1-2 c: zcccc\n9-10 q: lpqfbkdqqq\n4-9 t: wrjtzlrptlkvsmt\n13-15 l: klsgqllzlvlrbkl\n4-5 j: jjjljj\n6-11 p: ppppppppppbrppppp\n11-12 q: cqqqqqqqqqtq\n8-9 z: zzzxmzszz\n12-18 z: zzzzzzzzwzzzzzzzfzzz\n12-13 h: nhhhhshhptzjs\n3-5 c: ccmghck\n6-9 z: zzjzzjzvzzz\n12-15 d: dddddddddddkddddd\n7-11 c: qwcccmccccc\n10-12 t: ttttdvwrwrwgt\n2-3 l: ljbll\n8-9 x: xxxxdxxmx\n1-4 b: wpdgxfdsngsbbbpl\n16-19 f: fzwrdbblftrblfjfplw\n6-7 p: ppppppg\n5-7 l: mllllglllfllllvlxll\n4-6 b: bbnlbbbbbbbbbbbbb\n2-4 d: ldlb\n7-8 h: khpmrbhh\n3-5 l: llllllllllllllllllll\n6-8 f: ffqffffx\n4-8 n: vxmflsfn\n1-6 k: zkktzfzkdkkzkk\n9-11 j: jjjjjjjjjjkjjjjjjj\n4-7 b: bbkvbbnbbp\n1-13 v: fvvtmxvlvvxkvvs\n1-4 h: hrhf\n1-4 s: ksgs\n3-6 z: zzzzzzzzzzz\n13-16 j: djjjjjjjjjjjgjjrjj\n13-16 j: jjjjjjjjjjjhljjjjj\n4-6 k: hkhkkkvktpbkk\n2-3 p: pkqp\n2-13 j: jjjkjjjjjjjjdfjjj\n9-13 z: wwwzpkcczbmzntkjlf\n3-7 n: nntnnxgnnnnp\n2-3 t: ttpt\n1-7 h: qlthhgppdhfxhkhjhhd\n10-11 z: zmzzzzzjtxzzzzz\n3-6 p: ppppppp\n3-7 k: kmzvfhkkhrkk\n16-17 k: kpkkjkkklknkkkvkw\n2-5 z: zbzdzz\n1-9 z: zzzzzzzzzzzzzzzzzzz\n7-10 t: ttttttnttrt\n13-15 b: bbbbbbbbbbbbbbbb\n2-3 h: hcnhlxh\n3-5 h: lqdhh\n4-5 p: twptp\n1-4 x: xjxq\n4-6 c: ccmcgpvrfkwxpcjvcq\n5-6 h: chhhhcch\n6-13 m: fmmgmhmmjmgnbmmf\n5-8 j: ljzjljvjbjjjj\n7-10 j: jjjjjjjjjfj\n1-4 f: tffmkfnfw\n7-8 n: fbhgkmnndrbq\n1-14 m: mmmmmmmqmxmmmmm\n2-3 v: vzvtpmp\n5-10 q: qqqqqpqqfj\n4-9 q: qqqqqhqqqqqb\n8-11 b: bbbbbbbhbbbb\n1-3 r: rrsrr\n6-8 f: zktkmfnkffdffwfxj\n3-9 c: pkckkkxjc\n1-7 h: hhhhhhh\n1-16 s: kssssssssssssssnss\n8-16 j: gxhqtbnjcnrtjcdm\n9-13 d: vxdspdnmkfdbwlzd\n2-14 f: ffffffffvffffffffff\n18-19 x: xxfkppxxlxxxxxxtxqx\n5-9 t: txmtptttpqbtnjtts\n12-14 r: rrcrrrrrrrrtrxr\n9-10 d: ddddbdbddd\n1-3 z: mwzrz\n5-12 q: qqqqqqqqqqqlqqq\n3-15 s: rphsmgngvmqtdhnlk\n5-6 n: nnnnmtnt\n4-5 z: gzxzc\n1-6 r: rrcwrmcrx\n4-5 n: dbfnrn\n15-16 f: ffffffffgfkfffff\n3-9 w: twwxvxspkwtj\n17-18 l: lllllllllllllllllgl\n1-3 q: mtwqqq\n10-11 k: kkkkkkkkkkkkkkk\n12-18 f: rfvkgfffkkxfdfhgfvb\n4-8 q: cnnvqvvrllwmw\n17-18 q: qxqqlqlqtqqqqqqrmq\n8-9 s: sjssssszsjsss\n4-5 w: cwwxw\n5-6 q: qqqqqr\n3-4 d: cddrng\n6-10 c: bzcgscknzc\n16-18 l: ttvsvtfkmbfcglslpljk\n3-7 c: xlwtncn\n5-14 q: cpkvqljhwqqxjqsngm\n7-8 b: bbbgjbbcjbbt\n11-14 m: mmmmmmmmmmmmmrm\n3-4 v: kjgvv\n14-16 w: vpwmfdkfkwlrnldbhmq\n2-11 w: gkxjlmgwkpw\n16-17 k: kkkxkkkktknkmkzkk\n3-7 c: cfcpczsw\n16-18 l: cllrclltlllllmlbldlq\n7-13 q: brqtqqqbqnqbf\n13-14 q: vqtqqqqqqqqqrrq\n10-11 x: xxxxdxxxxxwxxqxxxxbx\n6-8 p: ppnnpspcpwwdppfzk\n5-6 s: sssqsr\n2-16 j: jvbhqwqzrrhhxdhjm\n2-13 x: prgsbszbcsvvxxs\n9-11 h: bhwghbhhhvqql\n16-18 q: xjmqwmrlvkpjxdtzbq\n9-10 t: wtwtctrtpnf\n11-12 j: jjfsjqjrjjjhp\n13-14 x: jflkltqvqdzqxxh\n3-4 p: xppp\n6-8 d: mwpndflpfpk\n7-12 h: pqfhhhhhnglgfl\n5-6 b: lnrvhbfdbbdncbqrrbj\n1-4 m: cdmfmmmmglgmmcmmzmmm\n4-6 h: hhhhhj\n3-5 s: pssssc\n12-16 x: svqltrtdlgxnhcmx\n7-11 j: jjjsjjjjjjjjjjjjjj\n14-16 w: xwxdgzfwstfjwckw\n11-14 b: cbhbgbjbvbbcbc\n4-6 l: bpgllnxlj\n4-5 x: xxrvx\n11-12 d: dmdddcjdddzdk\n1-16 l: fllllllllslllllllll\n3-9 q: qqqqqqqsqqqqlqfq\n14-15 d: hnmwvtprhxpdcgd\n5-7 t: zccgcttxttft\n3-6 d: ddvdddf\n11-12 r: rrgxjjnsfprrgk\n8-10 h: hhhhhhjchhfhshd\n2-5 b: bbbbvbbb\n18-19 m: msxttmbmqmmwdrgmmss\n3-6 b: vbhxtw\n7-8 r: tjrrmrrn\n4-11 g: wgdgwtwllgg\n4-5 j: jjjjkj\n15-16 r: rrrrrrrrrqrrrrdr\n15-16 q: qqqqqqqqqqqqqqqp\n2-8 r: kzswfbvcm\n4-8 b: bbbnbbbbbb\n11-13 m: mmmmmmmmmmmmbmm\n3-4 s: wnlss\n11-13 q: rlzbslqcqnqchqmv\n2-6 x: bxqwxzq\n6-7 k: dkjgdjk\n6-9 h: hhhqhhgrwp\n5-6 n: pnnldn\n3-4 z: zzzz\n4-5 d: sdnqz\n4-8 b: kkccxbfdxbljbn\n1-6 v: vbvvvzvvvv\n3-5 j: wwmwj\n2-13 q: sqtvrbspzkflqj\n2-3 z: xzhd\n4-7 p: pcpmpwpg\n4-6 f: fffgfwf\n13-15 k: zkkkkrhmlnlckkcvkk\n7-9 q: qqqqqqkqqqlqq\n14-16 g: gggggwgsqggggkgzghg\n4-10 v: vvvdvvvvvvv\n18-19 c: ccbfkgnrcwxccwtlczc\n1-8 j: jfjjjjjjqfjjljj\n12-14 w: wtwhwxwqwnwmwq\n4-5 l: ngwwzwxxrsbxmlnc\n3-4 n: nnqn\n4-13 d: kddfddddmtddd\n4-5 m: mmmpjmmm\n15-16 r: rrrrrrrrrrrrrrzrr\n17-19 k: krkxkzkkskcmdkckwkkf\n3-11 t: xgthvllfhjtb\n1-6 p: fppppfpp\n10-11 d: djdddlcdwlmd\n11-12 h: hhhmhshhhchdhr\n12-16 k: nwkvtkskkkwzkcbnkkc\n16-19 m: mvmmmmmmmmmmqjmtmml\n3-4 v: qvvwvvv\n1-4 w: npwp\n6-8 f: fffffffqfff\n2-3 w: zwmwwwn\n1-3 t: ttstr\n1-4 d: dddmddd\n2-3 m: mkmm\n1-7 f: fffqffffffffffff\n1-2 c: qcccccptcc\n7-8 g: ggghkgggbg\n13-15 b: dlbpbxdbrbbbbdcv\n3-4 p: pppcpvpq\n13-17 q: qqqqqqsqqqqqqqqqzqvq\n1-7 q: qqqqqqqq\n10-15 l: lmlsfkgpllldnlxgh\n4-5 t: cbptdgtq\n5-12 x: xxxxcxxxxxxxxxxxxx\n10-12 c: ccchcfckcctcccd\n4-16 h: cgqhppwhgqbqgthw\n18-20 h: cqsdtbjlhzwlrvztnlth\n1-10 v: vvvvvvqvzvvlvvvvvv\n11-12 l: qzlxrgrpglsldwrx\n2-4 t: ttttmt\n16-17 z: zzztzzzzzzzlzzzzt\n16-19 q: plqzzmhddkrwqtkqxmq\n14-18 l: lklllllzdlllpllqlll\n8-14 x: gxkcgwxxxxxkcx\n3-6 n: nnnnnvn\n1-4 p: cppp\n2-3 s: xsdvnd\n6-11 m: xghqpmmjmsm\n1-9 j: jhnjjjtpx\n15-17 z: pqrzgxlznjzzzmhcj\n1-11 z: zzzzzzzzzztzzz\n7-10 k: nkkkkhkkdgkcn\n8-14 g: xdxglrzzbdjrmqpdqrzn\n11-12 q: whfqqqsqfqqc\n5-12 s: rfltwhrsnsksqmbv\n8-9 s: sssssssstbssss\n2-3 h: hfhhhhh\n2-4 g: tjtg\n13-15 z: gzzzxzzzzzzzzzlz\n4-11 c: ccccbcccxccccccc\n13-14 k: krxkkkjkwkkkpk\n2-18 c: cccccccccccccccccccc\n4-12 r: rrrrrrrrrrrmrr\n7-8 s: sstsssfssbssssp\n8-11 s: ssssssswssfs\n1-4 j: jghnjdsgkwjqmmj\n3-6 m: mmvmmm\n10-17 f: sfffffjffffgfwpffv\n2-4 p: pcvp\n18-19 f: wcffffflnfffnffpklf\n6-12 x: ctkckxgxghjq\n1-8 h: fhhhhhhhhhh\n14-17 l: llllllllfllgldllhll\n3-4 z: zfmwpt\n7-8 h: whhxhhhhhhh\n6-11 j: shzqfjgwtwjmgwktjstw\n7-10 n: cnngnzddmnpp\n2-6 f: tffhffz\n1-2 k: lkkzpq\n4-11 d: vkfbjgsmdld\n1-6 v: xlxvtvwvvvxv\n11-14 d: ffddjdffcspxzk\n3-13 z: tfqwwhwqjzdlz\n13-17 p: hgpjfpqrtvspgfppsh\n11-12 l: pbvlvkrlllxxfdll\n12-14 f: flffpffffffcfxpfnfff\n5-7 p: mvbpppk\n10-11 s: sqbsxsrtscsxtssssgs\n9-12 q: qqqqqqqqqqqf\n6-7 t: mzwjtzt\n14-15 v: xlvxsvctjgqwjvgm\n5-6 t: tctttt\n7-14 g: nfxgkvwzgcthpjvnl\n1-4 g: gggg\n11-14 t: blftntttcxhttxtwtl\n2-7 k: jkklsjw\n12-17 t: jmtjdtttttttfhtjnsp\n14-18 b: qbtsbbtwplkklbjsfb\n12-14 x: xxxxcjxgxxxxdhxxxx\n4-5 s: ssssss\n4-7 z: zzzzzzzz\n8-13 w: swdwhlqwwwpsr\n4-6 b: bdtbbd\n12-15 h: hhfxkhhhhhhqhhhhhhh\n14-15 n: nnnnnnnnnnnnnnjnnnnn\n4-5 c: ccccpcccc\n13-16 l: llllmljtddqpvwlhl\n6-7 p: pppppfp\n7-9 g: gwggghvvw\n7-8 v: vfzdvvvvvvv\n4-5 m: mxcvm\n1-5 r: rrsmm\n4-6 d: nddvddrkdxdnfld\n6-7 q: qqqqqrt\n9-15 r: rrrrrrrrrrrrrrrrr\n8-9 g: zlgvbgnkg\n4-5 c: ctgcc\n1-12 s: cmsssssssssl\n4-6 x: xxxxxxxx\n2-4 t: tthpx\n10-13 f: kfggfnfftjwcf\n11-12 l: llllllllllrl\n4-10 t: ttttttttft\n2-3 x: xvnxj\n4-17 j: jjjjnjtjjjxljznjjj\n6-10 z: zzzzzzzzzlz\n5-14 j: vdwjbtntgzcchjhfl\n12-15 l: lllllhxlflpplllcl\n5-7 s: fqsssmswsgqlsvsssvf\n5-8 r: bhrhbqzrfr\n2-5 x: lxxxxxbwxdmx\n9-10 d: zdxjjdjdhdd\n16-18 l: lqllrtllbllvlllblll\n5-9 n: nnnnnnnnznnnnnn\n4-5 g: ggghzg\n6-14 m: ldmpkmtzslmdbljvsxtm\n7-8 p: ppppprpmsp\n8-13 p: qbkqpkpwmdgppnbp\n1-3 r: rrrrrr\n6-12 q: dqqqthphqhpqkqqmnmfl\n5-8 l: llllzzll\n9-10 m: ljmbmmmdjwdgmm\n4-5 s: qxsls\n17-18 k: nkfwsdvldkktdvsmkk\n7-8 l: llllllvlsl\n3-5 t: dlpctqfkwtbrvxfdxwr\n13-15 r: rrrvrrrrgrrrrrm\n4-16 l: llnlllllllllllllblll\n7-9 w: kwwwwwgwv\n7-8 z: hvzhphwm\n2-3 g: qgghtgk\n3-7 s: xsfqthm\n12-16 b: bbbbbbbbbbbbbbbzbbbb\n3-4 s: hbmzwncbvsgk\n14-18 s: ssssssssssssssssssss\n5-9 d: dtnbnxjlqh\n5-8 x: xxxxpxxz\n12-19 r: jwrrrrrrrrjrjrzgzrrm\n15-16 r: rrrrrrrrrrwrrrplr\n6-8 d: jddddlzj\n3-15 v: kfvrbtsxlmsqvzgbdk\n7-12 n: nfnnnmhnnnncnnnnnrnw\n14-17 c: zswgzkxrfvcvtcvjcjmf\n11-13 v: vvvvvvvvvvvvsv\n9-11 n: tnlnqnnntnn\n2-3 x: rxrlrfwm\n7-9 c: cccccccctcccccccc\n15-16 v: vvzvvwfvvvvvvvvrvvvv\n2-4 r: rrrmrr\n10-16 d: gldpwpdlxdhbtjhm\n5-6 z: zqzzzqzz\n3-4 n: nntnq\n2-3 z: gzdzbs\n6-20 j: gwjhxjjlzqdjmjvjcmjj\n5-6 q: qqxqsq\n6-7 g: gqgggpbkhjbg\n7-8 t: ttttcttj\n12-19 v: vvvvvvvvvvvnvvvvvvsv\n2-5 k: xkkwf\n3-7 g: ghqhtng\n3-8 v: vvvvvvvthl\n5-7 h: qjhxhrdwjz\n7-10 g: kmkxqcgfgmxw\n1-8 h: ghfhhhhhqhhhhh\n4-11 g: gqggbtgzgqgzgt\n6-7 t: tpxxtjt\n9-12 q: qqqqqsqqqtqqqnj\n3-7 m: znghsmmfczwvd\n4-11 t: stgjwttqvtstlsfcsssz\n9-10 f: fffffffffffff\n14-15 l: mlwlvnlpcxldflt\n1-6 g: gjgggwgg\n1-3 g: gggg\n2-4 s: jszs\n3-5 n: nnmnnn\n11-12 c: cccccjjcccdq\n11-17 h: hhhhhhhhhhshhhhhhh\n7-15 h: hpfdvhhhhnkhchb\n7-10 g: gggggggggg\n4-5 x: xjdxx\n5-13 p: hptpppgpmqsxpxnj\n3-4 b: bbbb\n5-6 s: psssgmqtcxs\n2-8 w: rwwtfxsh\n4-5 v: wjjvvtbhz\n10-12 j: qjjjjjjjjjrjjj\n4-6 h: hxhhvgwhrz\n9-12 x: xxxxxxjxxxzxxxxx\n7-14 w: wpzbwwvrkgdwwwm\n15-18 x: pbpnrpjkcmgxqtxxmxxx\n15-16 q: pzzgjwwqpgvcwvqj\n3-4 d: gdddpd\n6-16 b: bjsrnqsjhbzjsbvbpbq\n2-4 h: whsqxphkjwfq\n18-19 x: xxxxxxxxxxxxxxxxxxx\n3-4 n: nnnv\n5-7 n: nnnnxnnth\n2-7 k: gkkkpfn\n13-14 m: pwmmxznzmdsznm\n3-4 p: tzpg\n6-7 h: hzhhhhchh\n2-4 l: lllll\n5-8 g: gggggggxgggg\n1-4 q: qqqqq\n5-9 j: wjqjnjjjjjjjjjjjjjj\n5-10 g: gtgggkgggprxkzxgw\n1-4 p: rppp\n4-5 z: jzzzzzzzzz\n4-5 j: pjhjk\n3-10 j: ljpxdgdmlq\n1-2 z: tzjzz\n4-8 k: kkkpkkksk\n8-16 h: hhhhhhhhhhhhhhhnh\n5-8 n: nnxzhgtnnnvq\n9-10 d: hhdlcrwcdlktjxrg\n15-16 h: nhhlmfkhlhhhwthlhgqx\n3-4 p: pppppp\n7-8 h: mfkwxmhh\n4-5 l: lllllb\n1-4 s: ssls\n6-10 c: qtwkdpfwcg\n4-5 m: mvcmm\n12-13 n: nnnnfnnnnnnnn\n8-9 b: bbbbbbbcbpbbbsb\n5-6 f: ffffttf\n4-11 g: gggggtgggggqgg\n4-8 h: bmhhrjnxhkbhdhqjwswh\n12-14 x: xxxxxwxmxxxxwbxxxxp\n2-7 r: crrrrrrrrrrrpr\n7-9 m: zmmpzmmtg\n4-13 l: llldlllllllllll\n4-5 d: dddsdd\n3-4 s: bsnb\n3-4 h: chhh\n5-10 t: gswkzvktjbbrltsbdlx\n4-6 k: kkkrtkqk\n2-4 m: cswpvtrzfvf\n15-18 r: rrrrrrrrrrrrrrvrrjrr\n1-6 l: lllmlt\n19-20 w: nwnpwqsrvwwwwcwwldqw\n8-10 k: mkkkkkvtkl\n16-20 x: dgxxxxxxrxxnxxxhxxxx\n1-5 s: ssprbsb\n7-9 q: brqbqmqrb\n7-12 w: wmqwcwwwjwvwwtm\n10-14 n: nnnnnnnnnnnnnxn\n4-12 t: tkjdhjnmhqrtdbn\n2-7 w: fpllpbplqspwgqbz\n1-9 b: btltdnbkbm\n3-4 c: cccqcg\n11-12 k: xkkkkknlkmtk\n4-5 v: fvvhvjlwvwwwq\n9-11 p: pppjppprqtp\n8-9 b: bbmbhbtbbb\n4-5 t: ttttt\n12-13 d: ddzdddkkdddplcspl\n4-5 r: zlxrpdpfr\n6-9 d: vddgddthh\n6-7 j: mdqqjjjmmj\n4-5 p: ppppxd\n10-14 x: xxxxnxxxxxxxlx\n3-6 k: kkqkkbkkkkkkkkkkk\n2-6 q: tgbqsx\n4-7 z: zwnzzzfzfzz\n1-7 g: cggggggg\n2-3 x: xcxjmqr\n3-6 p: pppppm\n7-8 h: txhfcrhhwr\n7-8 s: ssblssns\n9-12 q: qqqbqqqqqqqjcq\n16-19 z: zzzzzzzzzzzzzzzzzzzz\n6-7 n: nnbgnhn\n7-9 t: tthmbtttpw\n2-6 z: zzzfzqz\n4-8 r: pvplrnqfjbbsfc\n1-3 v: wvvvv\n3-10 s: btzqhmbbkszgvqsqjzf\n9-13 g: xhjpmqkcgbmpgmgz\n13-18 c: kwjrlmdvcwjwlsbbbf\n9-11 d: ddfdhqddkdmmjx\n7-11 l: dlllllblclp\n10-11 x: xxwpktxscxxxxqz\n1-12 k: kkkkkkkkkkkkkkkkk\n4-5 h: hhhfhhhhh\n4-10 l: lmlwkllrklll\n7-8 d: ddddfdnd\n10-19 d: ddpdddwddwvdbddddddl\n5-6 g: bgsldggjk\n4-5 b: bgbnb\n2-5 x: xvxxxbxxtxx\n4-6 g: pggkjg\n3-8 j: jjkjjjjtjjjjjjjjjjj\n5-10 d: jcwvxxdhvk\n12-19 w: wwqwwwwbxcwvwwnwgdvw\n1-2 v: dtfd\n6-12 p: tpnshppqppgp\n5-11 w: kjwzpfwhxhtv\n7-8 h: mhhjkhhhv\n3-4 t: lttqf\n5-6 x: wkxxfj\n4-7 j: qnjcjjs\n6-10 x: kblfnxxzzx\n4-5 w: twwwxtwwpw\n4-8 p: pflpllppcwcq\n1-13 w: wdnbjrhtmxwkb\n10-11 x: xxxxxgxxxqcdx\n10-17 x: lgxwbzxsxxvxbmvxk\n3-4 r: rjkr\n8-11 x: xxxxxxxcxxxxx\n1-3 l: wlll\n3-6 d: sjmwjdtr\n6-10 b: bbbbbwbbws\n16-19 j: jjjjjjcjljwjjjjnjjjj\n4-10 q: qgxsjxzvgqzvqvx\n4-15 p: lppjppppppppppppp\n11-12 d: ddddddddbdtddd\n4-5 h: rpqch\n3-4 n: nnds\n1-4 h: gjxh\n14-20 s: cxscvksfqspjwvsnsnjk\n8-13 p: ppfppppdppppxpppppp\n3-7 w: wzwwwggwww\n3-8 p: xfdzhzpz\n5-7 j: jbjtwgjxj\n6-10 b: bdpfbfvlbbdbczbfbbzp\n1-8 w: jwlzdwnw\n15-17 l: llllllllllllllvlkl\n13-16 s: sssssssssssssssqss\n8-9 m: bqvdbrrfmm\n3-9 b: sbfbbjbbn\n5-6 z: zzzzzzzz\n18-19 d: dddddddddddddddddjd\n16-17 n: rnspnwwtnnnnnnnnln\n3-4 b: kmbw\n10-14 d: ddddddddqdchdfddd\n5-8 t: htttdstkm\n1-6 g: gggggdg\n4-6 h: qnjhhcb\n6-7 t: tttttfttt\n2-3 v: vpvvrvvv\n5-10 d: dddddddddldd\n15-17 p: dzpmmlcppwpbzwbxpj\n2-3 r: rwlhr\n7-8 j: jjjjsjgjjj\n4-9 m: mmnlmlmmmv\n7-9 r: vrrnrrrrrrrr\n6-7 m: mmmmmmqmm\n7-8 j: jnjjjszj\n11-12 t: tttxttsttttt\n7-8 t: dsjjwctn\n6-16 z: zzzzzzzzzzzzwzzpzz\n11-13 m: mmmmmmmmmmmmmmmmmm\n4-7 z: kzpzzzzmwcqmmf\n3-8 s: sbgcnjssbv\n6-7 k: kkkkkknkk\n2-11 w: wwlvtxglnfw\n3-5 k: gvmlg\n8-9 r: crrrrrrrvgrrfrr\n1-4 b: rbfbbbbbbgffmbnb\n4-6 t: rmwmjthbtlh\n3-10 c: ccmcccxcdc\n9-11 d: ldldddnddjd\n12-17 s: smsfnssssssdssssms\n3-6 m: qxmmmfdj\n9-10 v: vvvcvvvtvvzvvvvvv\n10-15 v: wdlmskhsvvqddxvgllj\n1-9 k: kkkkkkkkkckk\n9-14 h: hhxfsdnthnrhhxh\n7-10 p: pwcppppljphkppzjp\n7-8 n: snqcpwnnnnnplcc\n9-11 h: hhgrhsrhhhz\n1-4 v: zcvvvvsmzjjz\n1-6 x: txxxxqx\n4-13 j: jjjbjjjjjjjjpjj\n1-5 c: ccwvfccjpckcckbc\n2-4 t: btttbp\n4-5 d: cwwfd\n9-18 p: ppphppphfjspsbgpnp\n4-12 r: ksvfpjgmrsvr\n2-3 q: lhnq\n1-5 c: czcxc\n4-6 m: mmzmbzrq\n4-7 m: mmmmmmbm\n14-17 v: mvhvthfvvvvvvvvvvv\n5-12 d: cdwwdnbvhddqhwxbnghz\n8-16 z: mjzzzzzczzzzdzpz\n5-10 j: jjjjjjjjjnj\n15-16 q: qqqqqqqqqqqqqqlq\n4-7 g: hglngrgpkzmw\n4-17 c: xlscnhcbmccrxhsqc\n6-8 l: slqljlglr\n12-13 q: qqqqqqqqqqqdq\n4-9 q: qrcnpnqmdrq\n1-7 l: llllllt\n5-6 c: ccccpc\n1-2 c: clcc\n4-13 g: ggglggggggggwgn\n7-11 g: ggggxgbggggg\n5-10 b: bbbbbqbbbbbb\n9-12 d: ddddhdddddxddd\n1-3 d: jchdvcsdfzkt\n8-9 w: qqwvwpwwrswq\n1-2 p: mrnwxjtndwkzbgc\n3-6 h: hhcbhrhh\n3-14 m: frmxqtgnxdtphp\n3-8 l: llmllllll\n6-7 b: bdjgcbb\n3-11 x: mslfthzgqcscpctqcxd\n2-11 n: ncntpnnlnnnqnnn\n15-18 p: pppppppppdppgppppfhp\n4-5 z: wrfzzz\n6-7 h: tqgnhvhjhddzs\n2-4 m: mmvgw\n13-14 n: tpvfpwxlnqnbnvzk\n13-14 s: stnsssssvsssmscssss\n5-8 z: zzhkzdzdzz\n4-7 x: bxjxzgxsgmrmlln\n1-12 m: mmmmmmmmjbmbzmmbmmmm\n7-15 m: mmjmmmgmmmmmmmmmmmmf\n10-15 d: llrcldwnckfpwfv\n9-12 g: mfggggggvnggg\n2-10 p: grbctspghpj\n16-18 g: qhhgjwwgwggskcgghntq\n6-10 r: lrtdjrxqrrrnxrhc\n4-8 f: fffmfffkfff\n9-11 j: tbsxfwbwjjc\n11-14 s: sssssssssxsspss\n4-5 t: jthdjtttt\n1-7 j: jkjlhjh\n17-19 z: zzzvzzzzzzzzzzzzzzzg\n8-18 n: tgglgtfnnthsxpqvfj\n2-11 h: qhhlpkmbxqdrhhhkhrm\n9-14 m: xvmjjxmwmmmkmq\n14-15 c: cccccsccbccffjc\n9-10 j: xjzxpxjjfj\n8-10 x: pmfxfxxsxxq\n7-14 t: ttttttsqttttwttt\n4-16 d: dddwddddddddddddd\n8-10 f: fkcbqlffgqzvdzfhrn\n1-9 h: hhhhhhhhhhh\n5-12 v: lvvvvvvvvvvjvvt\n5-15 p: gqvpvpbxsgrgpvpjpvl\n14-15 d: qbnbtdrwkswfbtbdrx\n7-19 f: ffffffffffffffffffff\n7-9 r: jjbrrsqcrqmpxr\n4-5 h: mpthjh\n12-13 v: vvrvvvpvqvvvv\n11-17 c: ccccccccccccccccgc\n3-4 t: brttb\n5-9 r: rrrrxrrrrr\n1-4 f: fwfsfft\n15-16 z: znfzdqzzzhzwzgsr\n2-3 v: vwfvvvvvvvvvvvvvvvvv\n3-4 n: nnnzn\n2-11 k: hgptsdsqxlmpjqmt\n1-15 g: ggggcgggggggggtggggg\n14-15 z: xzzqzwzzjwrzkzz\n5-9 s: snsgscrssscs\n4-5 l: slbtm\n1-4 s: jmsjpss\n6-9 h: hhzhhhhhhhhhhhh\n3-12 b: xgbpbwqhzwzl\n10-11 g: gsgjggkgggg\n4-8 m: jrmlgmmm\n2-4 q: gqqd\n1-4 w: zqkwww\n6-7 f: bctdsvvs\n7-14 w: wwwwwwcwwwgwwk\n9-15 w: fjwgwpwwwstnrwgs\n4-9 v: qvvklvshvwvvrvvpvc\n1-6 w: wwswlw\n2-11 w: xmmqwjrvxglxnzmv\n13-15 x: hnxnvxxrxzvnrvxmxxx\n5-12 f: gfwtfxnmhfmjxlsllpx\n5-15 f: gmsrffblxvppctrt\n1-5 h: rhbhh\n1-2 c: cckccccccc\n12-13 c: cntcqmcltrcnc\n6-8 f: bbrflcrffqfj\n2-3 v: vvvkmzz\n1-2 j: jjknrj\n3-5 p: sppptqjp\n3-4 t: gttt\n1-3 h: hhzhjthh\n6-7 k: kkldvkg\n2-10 s: sssvsssssqswslljs\n1-10 w: swwwwwwwwwwwwwww\n10-17 j: jhjmbzjjjglbgndzhdm\n8-9 l: lllllllgdl\n7-10 g: wdsjsglvltf\n3-5 g: ggvggg\n6-10 j: jjnwrjnkjx\n2-14 p: wkdwppdpsnzvppb\n5-14 h: shjfjqhrwphgxhjqdm\n16-17 t: ttwhttgtttttnjtrqjt\n6-7 d: ldcdjdz\n13-16 j: wxjtlcpjjhxsfjzjzjj\n4-6 s: sqgpsx\n3-6 z: ztjtzz\n4-10 c: bjgcmlwfxc\n1-13 c: hcccccccccccfc\n4-7 b: grtbjrxfg\n3-4 z: mzxx\n9-10 t: tmgdbttltf\n4-12 n: qnvmnmnnnnnngnnn\n9-10 b: prllqxdsbbcw\n3-6 t: ttmtttbt\n7-8 c: kccpnccbc\n3-8 s: sssssssgss\n9-10 z: zzzzzzzzzzzzzzzz\n11-14 g: ggggggggggggghg\n4-7 m: vgjmmffhspmkn\n3-4 n: nnlnn\n4-5 t: ttttttt\n5-6 h: xrczhbqqhvhgq\n8-17 n: vntdgfcqbrnbdwngsnnq\n3-4 f: gflfptww\n9-16 k: kkfkkkkkckkkkkkkkk\n15-17 r: rrrrrrrrrrrrrrwrrr\n2-7 n: bnkpljdtdscpsczjlf\n9-11 k: vksvzkkkkkszkwb\n2-4 x: hxpxxsfr\n8-19 c: mlwncdlckrpnnwfjshgf\n7-10 j: jjjqjjljjjjj\n2-3 n: nlxnnnf\n2-4 k: kskkkkkk\n5-6 r: dmsrrr\n13-14 x: nwcxqzxpxxrxxx\n4-6 b: bbbcbbb\n3-6 l: wmpzrl\n1-7 w: wcdwwwwwkwtw\n3-8 w: dhwlrbrwbkzqhxfwpfxs\n3-20 n: vnpzwmdbvrhdrlvqhmsn\n1-4 z: zzzzjzzz\n6-13 s: pfnzxdsssxpps\n4-12 q: qhqtrkwmfgvn\n4-10 t: mwgrtptprhtbzcfbvcc\n12-17 n: nnnnnnknwgnnknfgn\n2-3 k: jtnkk\n9-10 f: ffffhffffffff\n1-7 k: pkkkkkvpkkkkkk\n5-7 n: nnnnnnnn\n5-9 t: tcdttttttxj\n9-10 x: xxxxxxfxwxxxx\n12-17 s: pbsxshpbdssshfmxcbs\n4-10 n: nfnpnnnnnp\n1-7 d: vdddddbd\n7-8 r: rrrrrrrq\n10-12 v: vhvvvvvvvvzq\n6-10 n: xcnnbpnwnx\n4-5 f: ffflf\n5-9 b: jxpbbcjgrbqfbbmbbbf\n10-11 b: wchbzbbbbgbdb\n4-7 s: ssrnsssss\n5-12 p: vppgpdnkrptfv\n8-19 j: jvjjjjjtjjjjjjjjjjjj\n10-14 l: llllllllllklll\n11-13 h: hqhvhhchbjprhzhl\n1-10 h: sjfhhhhhhljrhhhh\n8-9 w: wwwwwwwww\n9-10 t: tgtttttrtt\n5-11 x: lgqsgxxpzprx\n3-4 n: qvngbnrvt\n2-9 x: drqckjkxkbs\n8-14 w: wjbswmpwwwhwwq\n3-14 k: tktxjdzrzrdnlk\n2-3 l: hmllslzlx\n10-11 s: sssssssssmss\n3-4 m: mmmnm\n4-5 b: jbbbq\n4-12 q: hqbwbbtqtqdvjjj\n2-10 g: gdcddxlrkgmzw\n1-13 z: zsgfzqpbbzjndvz\n1-10 b: nbbbbbbbbqbdb\n4-7 j: jmrbqsj\n9-15 r: prrkbzsgjlrrtlrrsp\n8-12 h: hhfhhkhhphfq\n4-17 p: hvgrcjhzmsmmpznlphb\n14-17 v: lvvvvvvvvvvvvvvvbvvp\n9-13 r: bhwrptnrrcpzxtrhvd\n5-7 b: trbqnthjrfwdg\n12-14 j: jjjjjjjjjjjjbsjj\n2-4 s: jswmwvtw\n1-7 c: mccctnkcccccccm\n7-8 m: zjmmmmgd\n4-6 x: xhqxddbfrngbxfzb\n6-9 x: xdtqxbdwx\n10-14 r: qskrrvvzclrrvztrdkrh\n4-13 c: wccwcccbccccrccc\n11-14 l: wlqzllcljjtqglbhl\n17-19 b: bbbbbbbbbbbbbbbbbbq\n3-5 x: xxqxm\n"
	inputSplit := strings.Split(input, "\n")
	part1(inputSplit)
	part2(inputSplit)
	fmt.Println("[*] Done...")
}

func part1(input []string) {
	valid := 0
	for _, rule := range input {
		if rule == "" {
			break
		}
		lower, upper, letter, pass := ruleSplit(rule)
		count := strings.Count(pass, letter)
		if count <= upper && count >= lower {
			valid++
		}
	}
	fmt.Printf("[*] Valid Passwords: %d\n", valid)
}

func part2(input []string) {
	valid := 0
	for _, rule := range input {
		if rule == "" {
			break
		}
		pos1, pos2, letter, pass := ruleSplit(rule)
		pos1Found := false
		pos2Found := false

		if string(pass[pos1-1]) == letter {
			pos1Found = true
		}
		if string(pass[pos2-1]) == letter {
			pos2Found = true
		}
		if pos1Found && pos2Found {

		} else if pos1Found || pos2Found {
			valid++
		}
	}

	fmt.Printf("[*] Valid Passwords: %d\n", valid)
}

func ruleSplit(rule string) (int, int, string, string) {
	ruleSplit := strings.Split(rule, " ")
	num := ruleSplit[0]
	numSplit := strings.Split(num, "-")
	lower, _ := strconv.Atoi(numSplit[0])
	upper, _ := strconv.Atoi(numSplit[1])
	letter := string(ruleSplit[1][0])
	pass := ruleSplit[2]
	return lower, upper, letter, pass
}

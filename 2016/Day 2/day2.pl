use strict;
use warnings;
my $input = "DUURRDRRURUUUDLRUDDLLLURULRRLDULDRDUULULLUUUDRDUDDURRULDRDDDUDDURLDLLDDRRURRUUUDDRUDDLLDDDURLRDDDULRDUDDRDRLRDUULDLDRDLUDDDLRDRLDLUUUDLRDLRUUUDDLUURRLLLUUUUDDLDRRDRDRLDRLUUDUDLDRUDDUDLLUUURUUDLULRDRULURURDLDLLDLLDUDLDRDULLDUDDURRDDLLRLLLLDLDRLDDUULRDRURUDRRRDDDUULRULDDLRLLLLRLLLLRLURRRLRLRDLULRRLDRULDRRLRURDDLDDRLRDLDRLULLRRUDUURRULLLRLRLRRUDLRDDLLRRUDUDUURRRDRDLDRUDLDRDLUUULDLRLLDRULRULLRLRDRRLRLULLRURUULRLLRRRDRLULUDDUUULDULDUDDDUDLRLLRDRDLUDLRLRRDDDURUUUDULDLDDLDRDDDLURLDRLDURUDRURDDDDDDULLDLDLU
LURLRUURDDLDDDLDDLULRLUUUDRDUUDDUDLDLDDLLUDURDRDRULULLRLDDUDRRDRUDLRLDDDURDUURLUURRLLDRURDRLDURUDLRLLDDLLRDRRLURLRRUULLLDRLULURULRRDLLLDLDLRDRRURUUUDUDRUULDLUDLURLRDRRLDRUDRUDURLDLDDRUULDURDUURLLUDRUUUUUURRLRULUDRDUDRLLDUDUDUULURUURURULLUUURDRLDDRLUURDLRULDRRRRLRULRDLURRUULURDRRLDLRUURUDRRRDRURRLDDURLUDLDRRLDRLLLLRDUDLULUDRLLLDULUDUULLULLRLURURURDRRDRUURDULRDDLRULLLLLLDLLURLRLLRDLLRLUDLRUDDRLLLDDUDRLDLRLDUDU
RRDDLDLRRUULRDLLURLRURDLUURLLLUUDDULLDRURDUDRLRDRDDUUUULDLUDDLRDULDDRDDDDDLRRDDDRUULDLUDUDRRLUUDDRUDLUUDUDLUDURDURDLLLLDUUUUURUUURDURUUUUDDURULLDDLDLDLULUDRULULULLLDRLRRLLDLURULRDLULRLDRRLDDLULDDRDDRURLDLUULULRDRDRDRRLLLURLLDUUUDRRUUURDLLLRUUDDDULRDRRUUDDUUUDLRRURUDDLUDDDUDLRUDRRDLLLURRRURDRLLULDUULLURRULDLURRUURURRLRDULRLULUDUULRRULLLDDDDURLRRRDUDULLRRDURUURUUULUDLDULLUURDRDRRDURDLUDLULRULRLLURULDRUURRRRDUDULLLLLRRLRUDDUDLLURLRDDLLDLLLDDUDDDDRDURRL
LLRURUDUULRURRUDURRDLUUUDDDDURUUDLLDLRULRUUDUURRLRRUDLLUDLDURURRDDLLRUDDUDLDUUDDLUUULUUURRURDDLUDDLULRRRUURLDLURDULULRULRLDUDLLLLDLLLLRLDLRLDLUULLDDLDRRRURDDRRDURUURLRLRDUDLLURRLDUULDRURDRRURDDDDUUUDDRDLLDDUDURDLUUDRLRDUDLLDDDDDRRDRDUULDDLLDLRUDULLRRLLDUDRRLRURRRRLRDUDDRRDDUUUDLULLRRRDDRUUUDUUURUULUDURUDLDRDRLDLRLLRLRDRDRULRURLDDULRURLRLDUURLDDLUDRLRUDDURLUDLLULDLDDULDUDDDUDRLRDRUUURDUULLDULUUULLLDLRULDULUDLRRURDLULUDUDLDDRDRUUULDLRURLRUURDLULUDLULLRD
UURUDRRDDLRRRLULLDDDRRLDUDLRRULUUDULLDUDURRDLDRRRDLRDUUUDRDRRLLDULRLUDUUULRULULRUDURDRDDLDRULULULLDURULDRUDDDURLLDUDUUUULRUULURDDDUUUURDLDUUURUDDLDRDLLUDDDDULRDLRUDRLRUDDURDLDRLLLLRLULRDDUDLLDRURDDUDRRLRRDLDDUDRRLDLUURLRLLRRRDRLRLLLLLLURULUURRDDRRLRLRUURDLULRUUDRRRLRLRULLLLUDRULLRDDRDDLDLDRRRURLURDDURRLUDDULRRDULRURRRURLUURDDDUDLDUURRRLUDUULULURLRDDRULDLRLLUULRLLRLUUURUUDUURULRRRUULUULRULDDURLDRRULLRDURRDDDLLUDLDRRRRUULDDD";

my @keypad = (
  [1,2,3],
  [4,5,6],
  [7,8,9]
);
my $x = 1;
my $y = 1;
my $code = '';
foreach my $line (split /\n/, $input) {
  $x = 1;
  $y = 1;
  foreach my $dir (split //, $line) {
    if($dir eq 'D'){
      if($x != 2){
        $x+=1;
      }
    }elsif($dir eq 'U'){
      if($x != 0){
        $x-=1;
      }
    }elsif($dir eq 'R'){
      if($y != 2){
        $y+=1;
      }
    }elsif($dir eq 'L'){
      if($y != 0){
        $y-=1;
      }
    }
  }

  $code = $code.$keypad[$x][$y];
  print "END OF LINE ";
  print "$code - $x,$y\n";
}

print $code;
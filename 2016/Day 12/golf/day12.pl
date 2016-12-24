use strict;
use warnings;
use Data::Dumper qw(Dumper);
my $l = "cpy 1 a
cpy 1 b
cpy 26 d
jnz c 2
jnz 1 5
cpy 7 c
inc d
dec c
jnz c -2
cpy a c
inc a
dec b
jnz b -2
cpy c b
dec d
jnz d -6
cpy 18 c
cpy 11 d
inc a
dec d
jnz d -2
dec c
jnz c -5";

#$input = "cpy 41 a
#inc a
#inc a
#dec a
#jnz a 4
#dec a
#inc a
#inc a
#inc b
#jnz b 2
#jnz 1 54
#dec b
#jnz 1 -2";

my@cA=();push @cA,$_ for split/\n/,$l;my%rV=(c=>0);my$i=0;
$cA[$i]=~/^cpy (.+)\s(.)/ ?do{my $reg = $2;$1=~/([abcd])/ ?do{my$sR=$1;$rV{$reg}=$rV{$sR};}:do{$rV{$reg}=$1};$i++;}:$cA[$i]=~/^inc\s([abcd])/ ?do{$rV{$1}=$rV{$1}+1;$i++;}:$cA[$i]=~/^dec\s([abcd])/ ?do{$rV{$1}=$rV{$1}-1;$i++;}:$cA[$i]=~/^jnz\s(\d+|[abcd])\s(.+)/ ?do{my$in=$2;my$v=$1;$v=~s/([abcd])/$v=$rV{$v}/e;$v!=0?do{$i+=$in}:$i++;}:0 while $i<@cA;
print $rV{'a'};

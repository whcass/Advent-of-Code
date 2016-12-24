use strict;
use warnings;
use Data::Dumper qw(Dumper);
my $input = "cpy 1 a
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

my @commandArray = ();

foreach my $com (split /\n/,$input) {
  push @commandArray,$com;
}

my %registerValues = (
  a=>0,
  b=>0,
  c=>1,
  d=>0
);
my $i = 0;
while($i<@commandArray){
  if($commandArray[$i]=~/^cpy (.+)\s(.)/){
    my $register = $2;
    if($1=~/([abcd])/){
      my $sendReg = $1;
      $registerValues{$register} = $registerValues{$sendReg};
    }else{
      $registerValues{$register} = $1;
    }
  }elsif($commandArray[$i]=~/^inc\s([abcd])/){
    $registerValues{$1} = $registerValues{$1}+1;
  }elsif($commandArray[$i]=~/^dec\s([abcd])/){
    $registerValues{$1} = $registerValues{$1}-1;
  }elsif($commandArray[$i]=~/^jnz\s(\d+|[abcd])\s(.+)/){
    my $increment = $2;
    my $value = $1;

    if($1=~/([abcd])/){
      $value = $registerValues{$1};
    }
    if($value != 0){
      $i+=$increment;
      next;
    }
  }
  $i++;
}

print Dumper \%registerValues;

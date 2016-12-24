use strict;
use warnings;
use Digest::MD5 qw(md5 md5_hex md5_base64);
my $i = 0;

my $string = 'abc';
my $count = 0;

while($count<64){
  my $data = "$string$i";
  my $digest = md5_hex($data);

  $digest = md5_hex $digest for 1 .. 2016;



  if($digest=~/([0-9a-f])\1{2}/){
    #print "$1\n";
    #print "$digest - $i - MATCH\n";
    #print "Checking the next 1000 indexes...\n";
    my $j = 1;
    my $char = $1;
    while($j<1000){
      my $index = $i+$j;
      my $data = "$string$index";
      my $digest2 = md5_hex($data);
      $digest2 = md5_hex $digest2 for 1 .. 2016;
      #print "$data - $digest - $index - $i\n";
      if($digest2=~/$char{5}/g){
        $count++;
        #print "$char\n";
        #print "$digest\n";
        #print "$digest2\n";
        print "$count - Key found at index $i($j).\n";
        last;
      }
      $j++
    }
  }
  $i++;
}

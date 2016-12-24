use strict;
use warnings;
use Digest::MD5 qw(md5 md5_hex md5_base64);
my $i = 0;
my $string = 'uqwqemis';
my $count = 0;
my $password = '________';

while($count<9){
  my $data = "$string$i";
  my $digest = md5_hex($data);
  print "$data-$digest-$password-$count\n";
  if($digest =~ /^0{5}[0-9]/){
    my $pos = substr($digest,5,1);
    if($pos<=7){
      my $char = substr($digest,6,1);
      substr($password,$pos,1) = $char;

      $count++;
    }
  }
  $i++;
}

print $password;

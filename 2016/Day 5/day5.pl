use strict;
use warnings;
use Digest::MD5 qw(md5 md5_hex md5_base64);
my $i = 0;
my $string = 'uqwqemis';
my $count = 0;
my $password = '';
while($count<8){
  my $data = "$string$i";
  my $digest = md5_hex($data);
  print "$data-$digest-$password-$count\n";
  if($digest =~ /^0{5}/){
    $password = $password.substr($digest,5,1);
    $count++;
  }
  $i++;
}

print $password;

use strict;
use warnings;
use Digest::MD5 qw(md5 md5_hex md5_base64);
my $string = 'bgvyzdsv';
my $i = 0;
while($i<10000000){
	$i++;
	my $data = "$string$i";
	my $digest = md5_hex($data);
	if($digest =~ /^0{6,}/){
		print "$digest - $data - $i\n";
		last;
	}
}
print $i;
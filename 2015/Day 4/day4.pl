use strict;
use warnings;
use Digest::MD5 qw(md5 md5_hex md5_base64);

my $string = 'bgvyzdsv';
my $i = 0;

while($i<1000000){
	$i++;
	my $data = "$string$i";
	my $digest = md5_hex($data);
	if($digest =~ /^0{5,}/){
		print "$digest - $data - $i\n";
		last;
	}
	print "$digest - $data - $i\n";
}

print $i;
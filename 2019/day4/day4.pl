use strict;
use warnings;

my $min = 372304;
my $max = 847060;
my $inc = 1;
my $val = $min;
my $count = 0;
for($val;$val<=$max;$val+=$inc){
	my @digit = (split //,$val);
	my $pass = 1;
	for(my $i = 0;$i<@digit;$i++)
	{
		if($i!=@digit){
			if($digit[$i] > $digit[$i+1])
			{
				$pass = 0;
				last;
			}
		}
	}

	if($pass == 1){
		print "$val\n";
	}
}

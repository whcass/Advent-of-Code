#!/usr/bin/perl
use strict;
use warnings;

my $file = "day2.txt";
open my $data, $file or die "Could not open file: $!";
my @code = split /,/,<$data>;
#print "@code[0]\n";
my @codecopy = @code;
for(my $noun = 0; $noun < 99;$noun++){
	for(my $verb = 0;$verb < 99;$verb++){
		$code[1] = $noun;
		$code[2] = $verb;
		my $ptr = 0;
		while (1){
			my $inst = $code[$ptr];
			if($inst == 99)
			{
				last;
			}elsif($inst == 1)
			{
				my $first = $code[$code[$ptr+1]];
				my $second = $code[$code[$ptr+2]];
				my $out = $code[$ptr+3];
				#print "$first : $second : $out\n";
				$code[$out] = $first+$second;
			}elsif($inst == 2)
			{
				my $first = $code[$code[$ptr+1]];
				my $second = $code[$code[$ptr+2]];
				my $out = $code[$ptr+3];
				#print "OUT:$out\n";
				#print "$first : $second : $out\n";
				$code[$out] = $first*$second;
			}
			$ptr+=4;
		}
		if($code[0] == 19690720)
		{
			my $res = 100 * $noun + $verb;
			print "[*] $code[0] : $noun : $verb : $res\n";
			exit;;
		}
		@code = @codecopy;
	}
}

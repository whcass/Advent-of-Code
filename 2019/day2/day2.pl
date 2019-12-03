#!/usr/bin/perl
use strict;
use warnings;

my $file = "day2.txt";
open my $data, $file or die "Could not open file: $!";
my @code = split /,/,<$data>;
$code[1] = 12;
$code[2] = 2;
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
		$code[$out] = $first+$second;
	}elsif($inst == 2)
	{
		my $first = $code[$code[$ptr+1]];
		my $second = $code[$code[$ptr+2]];
		my $out = $code[$ptr+3];
		$code[$out] = $first*$second;
		}
	$ptr+=4;
}

print "$code[0]\n";
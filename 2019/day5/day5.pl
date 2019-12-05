#!/usr/bin/perl
use strict;
use warnings;

sub GetParameterModes {
	my $code = $_[0];
	my @digit = split//,$code;
	my $len = @digit;

	if ($len == 1)
	{
		return (0,0,0,$digit[0]);
	}
	elsif($len == 3)
	{
		return (0,0,$digit[0],$digit[2]);
	}
	elsif($len == 4)
	{
		return (0,$digit[0],$digit[1],$digit[3]);
	}
	elsif($len == 5)
	{
		return ($digit[0],$digit[1],$digit[2],$digit[4]);
	}
}



my $file = "day5.txt";
open my $data, $file or die "Could not open file: $!";
# my $test = "1101,100,-1,4,0";
my @code = split /,/,<$data>;
# $code[1] = 12;
# $code[2] = 2;
my $ptr = 0;
my $input = 1;
my  $ptrinc;
# print "second:$code[1]\n";
while (1){
	my $opcode = $code[$ptr];
	if($opcode == 99)
	{
		last;
	}
	my @modes = GetParameterModes($opcode);
	# print "modes:@modes\n";
	my $inst = $modes[3];
	my $p1 = $modes[2];
	my $p2 = $modes[1];
	my $p3 = $modes[0];
	# print "inst:$inst\n";
	# print "p1:$p1\n";
	# print "p2:$p2\n";
	
	if($inst == 1)
	{
		my $first = $p1==0 ? $code[$code[$ptr+1]] : $code[$ptr+1];
		# print "first:$first\n";
		my $second = $p2==0 ? $code[$code[$ptr+2]] : $code[$ptr+2];
		# print "second:$second\n";
		my $out = $code[$ptr+3];
		$code[$out] = $first+$second;
		$ptrinc = 4;
	}
	elsif($inst == 2)
	{
		my $first = $p1==0 ? $code[$code[$ptr+1]] : $code[$ptr+1];
		# print "first:$first - $code[$code[$ptr+1]]\n";
		my $second = $p2==0 ? $code[$code[$ptr+2]] : $code[$ptr+2];
		my $out = $code[$ptr+3];
		$code[$out] = $first*$second;
		$ptrinc = 4;
	}
	elsif($inst == 3)
	{
		my $out = $code[$ptr+1];
		$code[$out] = $input;
		$ptrinc = 2;
	}
	elsif($inst == 4)
	{
		my $out = $code[$ptr+1];
		print "$code[$out]\n";
		$ptrinc = 2;
	}
	$ptr+=$ptrinc;
	# print "ptr:$code[$ptr]\n";
	}

# print "$code[0]\n";

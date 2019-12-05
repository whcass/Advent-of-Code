#!/usr/bin/perl
use strict;
use warnings;

sub GetParameterModes {
	my $code = $_[0];
	my @digit = split//,$code;
	my $len = @digit;
	print "len:$len\n";
	if ($len == 1)
	{
		return (0,0,0,$digit[1]);
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
my $test = "1002,4,3,4,33";
my @code = split /,/,$test;
$code[1] = 12;
$code[2] = 2;
my $ptr = 0;
my $input = 1;
my  $ptrinc;
#while (1){
	my $opcode = $code[$ptr];
	my @modes = GetParameterModes($opcode);
		
	my $inst = $modes[3];
	my $p1 = $modes[2];
	my $p2 = $modes[1];
	my $p3 = $modes[0];
	print "inst:$inst\n";
	if($inst == 99)
	{
		last;
	}
	elsif($inst == 1)
	{
		my $first = $code[$code[$ptr+1]];
		my $second = $code[$code[$ptr+2]];
		my $out = $code[$ptr+3];
		$code[$out] = $first+$second;
		$ptrinc = 4;
	}
	elsif($inst == 2)
	{
		my $first = $code[$code[$ptr+1]];
		my $second = $code[$code[$ptr+2]];
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
		print "$code[$out]";
		$ptrinc = 2;
	}
	$ptr+=$ptrinc;
	#}

print "$code[0]\n";

use strict;
use warnings;
use Data::Dumper;
my $min = 372304;
my $max = 847060;
my $inc = 1;
my $val = $min;
my $count = 0;

sub ascending {
	my $d = $_[0];
	my @digit = split//,$d;
	my $last = $digit[0];
	foreach (@digit){
		if($_<$last){
			return 0;
		}
		$last = $_;
	}
	return 1;
}

sub doubles {
	my $d = $_[0];
	my @digit = split//,$d;
	my $last = $digit[0];
	foreach(@digit){
		if($_ == $last){
			return 1;
		}
		$last = $_;
	}

	return 0;
}

sub doublesPart2 {
	my $d = $_[0];
	my @digit = split//,$d;
	my $last = $digit[0];
	my $doubles = 0;
	my %hash;
	$hash{$_}++ for @digit;
	foreach(keys %hash){
		if($hash{$_} == 2){
			return 1;
		}
	}
	return 0;

}
for($val;$val<=$max;$val+=$inc){
	my $pass = 1;
	if (ascending($val) == 0){
		$pass = 0;
	}

	if (doubles($val) == 0){
		$pass = 0;
	}

	if(doublesPart2($val) == 0){
		$pass = 0;
	}

	if($pass == 1){
		$count++;
	}
}

print $count;

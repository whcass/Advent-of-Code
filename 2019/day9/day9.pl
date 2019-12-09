use strict;
use warnings;
use File::Basename qw(dirname);
use Cwd qw(abs_path);
use lib dirname( dirname abs_path $0) . '/lib';
use AdventOfCode::ComputerObj qw(new run);

my $file = "day9.txt";
open my $data, $file or die "Could not open file: $!";
my @input            = split /,/, <$data>;

my $comp =AdventOfCode::ComputerObj->new( \@input );

my @out = $comp->run(1);
pop @out;
print "@out";
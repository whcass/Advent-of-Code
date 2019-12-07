#!/usr/bin/perl
use strict;
use warnings;
use Data::Dumper;
use File::Basename qw(dirname);
use Cwd qw(abs_path);
use lib dirname( dirname abs_path $0) . '/lib';
use Algorithm::Combinatorics qw(permutations);
use AdventOfCode::ComputerObj qw(new run);

my $file = "part2test.txt";
open my $data, $file or die "Could not open file: $!";
my @input            = split /,/, <$data>;
my @inputCopy        = @input;
my @sequence         = ( 5, 6, 7, 8, 9 );
my @all_permutations = permutations( \@sequence );

# print Dumper(@all_permutations);
my $highest = 0;

# my @highestPermute;

my @out = $computer->run(1);
print Dumper(@out);
foreach my $permute (@all_permutations) {
    my @permute = @{$permute};

    # Computer A
    my $A = AdventOfCode::ComputerObj->new( \@input );
    @input = @inputCopy;
    my $B = AdventOfCode::ComputerObj->new( \@input );
    @input = @inputCopy;
    my $C = AdventOfCode::ComputerObj->new( \@input );
    @input = @inputCopy;
    my $D = AdventOfCode::ComputerObj->new( \@input );
    @input = @inputCopy;
    my $E = AdventOfCode::ComputerObj->new( \@input );
    @input = @inputCopy;

    while(1){
        my @aout = $A->run
    }
}

print "[*] highest: $highest\n";


#!/usr/bin/perl
use strict;
use warnings;
use Data::Dumper;
use File::Basename qw(dirname);
use Cwd qw(abs_path);
use lib dirname(dirname abs_path $0) . '/lib';
use Algorithm::Combinatorics qw(permutations);
use AdventOfCode::Computer qw(Computer);


my $file = "day7.txt";
open my $data, $file or die "Could not open file: $!";
my @input = split /,/,<$data>;
my @inputCopy = @input;
my @sequence = (0, 1, 2, 3, 4);
my @all_permutations = permutations(\@sequence);
# print Dumper(@all_permutations);
my $highest = 0;
# my @highestPermute;
foreach my $permute (@all_permutations){
    my @permute = @{$permute};
    # Computer A
    my @Ainput = ($permute[0],0);
    # print Dumper(@Ainput);
    my @output = Computer( \@input, \@Ainput );
    # Computer B
    # print "[*] A: $output[0]\n";
    my @Binput = ($permute[1],$output[0]); 
    @input = @inputCopy;
    @output = Computer( \@input, \@Binput );
    # print "[*] B: $output[0]\n";
    # Computer C
    my @Cinput = ($permute[2],$output[0]);
    @input = @inputCopy;
    @output = Computer( \@input, \@Cinput );
    # print "[*] C: $output[0]\n";
    # Computer D
    my @Dinput = ($permute[3],$output[0]);
    @input = @inputCopy;
    @output = Computer( \@input, \@Dinput );
    # print "[*] D: $output[0]\n";
    # Computer E
    my @Einput = ($permute[4],$output[0]);
    @input = @inputCopy;
    @output = Computer( \@input, \@Einput );
    # print "[*] E: $output[0]\n";
    # print "[*] $output[0]\n";
    # last;
    if($output[0] > $highest){
        $highest = $output[0];
    }
}

print "[*] highest: $highest\n";


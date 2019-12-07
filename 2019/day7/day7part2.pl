#!/usr/bin/perl
use strict;
use warnings;
use Data::Dumper;
use File::Basename qw(dirname);
use Cwd qw(abs_path);
use lib dirname( dirname abs_path $0) . '/lib';
use Algorithm::Combinatorics qw(permutations);
use AdventOfCode::ComputerObj qw(new run);

my $file = "day7.txt";
open my $data, $file or die "Could not open file: $!";
my @input            = split /,/, <$data>;
my @inputCopy        = @input;
my @sequence         = ( 5, 6, 7, 8, 9 );
my @all_permutations = permutations( \@sequence );

my $highest = 0;

my @A = @input;
my @B = @input;
my @C = @input;
my @D = @input;
my @E = @input;

foreach my $permute (@all_permutations) {
    my @permute = @{$permute};

    my $A    = AdventOfCode::ComputerObj->new( \@A, $permute[0] );
    my $B    = AdventOfCode::ComputerObj->new( \@B, $permute[1] );
    my $C    = AdventOfCode::ComputerObj->new( \@C, $permute[2] );
    my $D    = AdventOfCode::ComputerObj->new( \@D, $permute[3] );
    my $E    = AdventOfCode::ComputerObj->new( \@E, $permute[4] );
    my $total = 0;
    my $steps = 0;
    my (@aout,@bout,@cout,@dout,@eout) = [];
    print "[*] permute @permute - ";
    while (1) {
        if ($steps==0) {
            @aout = $A->run( 0 );
            $steps++;
        }
        else {
            @aout = $A->run( $eout[-1] );
        }
        @bout = $B->run( $aout[-1] );
        @cout = $C->run( $bout[-1] );
        @dout = $D->run( $cout[-1] );
        @eout = $E->run( $dout[-1] );
        if ( $eout[-1] == 99 ) {
            last;
        }
    }

    print "$eout[-2]\n";
    my $out = $eout[-2];
    if ( $out > $highest ) {
        $highest = $out;
    }
}

print "[*] highest: $highest\n";


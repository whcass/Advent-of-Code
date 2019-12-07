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

# my @out = $computer->run(1);
# print Dumper(@out);
foreach my $permute (@all_permutations) {
    my @permute = @{$permute};

    # Computer A
    my $A    = AdventOfCode::ComputerObj->new( \@input, $permute[0] );
    @input = @inputCopy;
    my $B    = AdventOfCode::ComputerObj->new( \@input, $permute[1] );
    @input = @inputCopy;
    my $C    = AdventOfCode::ComputerObj->new( \@input, $permute[2] );
    @input = @inputCopy;
    my $D    = AdventOfCode::ComputerObj->new( \@input, $permute[3] );
    @input = @inputCopy;
    my $E    = AdventOfCode::ComputerObj->new( \@input, $permute[4] );
    @input = @inputCopy;
    my $total = 0;
    my $steps = 0;
    my (@aout,@bout,@cout,@dout,@eout) = [];
    while (1) {
        if ($steps > 100){
            exit;
        }
        if ($steps) {
            print "[*] A Running...\n";
            print "[*] First time A\n";
            @aout = $A->run( 0 );
            print "[*] A Waiting...\n";
            $steps++;
        }
        else {
            print "[*] A Running...\n";
            @aout = $A->run( $eout[-1] );
            print "[*] A Waiting...\n";
        }
        
        print "[*] B Running...\n";
        @bout = $B->run( $aout[-1] );
        print "[*] B Waiting...\n";
        print "[*] C Running...\n";
        @cout = $C->run( $bout[-1] );
        print "[*] C Waiting...\n";
        print "[*] D Running...\n";
        @dout = $D->run( $cout[-1] );
        print "[*] D Waiting...\n";
        print "[*] E Running...\n";
        @eout = $E->run( $dout[-1] );
        print "[*] E Waiting...\n";
        print "[*] @eout\n";
        # exit;
        # last;
        if ( $eout[-1] == 99 ) {
            last;
        }
        else {
            $total += $eout[-1];
        }
        $steps++;
    }

    # last;

    if ( $total > $highest ) {
        $highest = $total;
    }
}

print "[*] highest: $highest\n";


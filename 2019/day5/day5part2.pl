#!/usr/bin/perl
use strict;
use warnings;

sub GetParameterModes {
    my $code  = $_[0];
    my @digit = split //, $code;
    my $len   = @digit;

    if ( $len == 1 ) {
        return ( 0, 0, 0, $digit[0] );
    }
    elsif ( $len == 3 ) {
        return ( 0, 0, $digit[0], $digit[2] );
    }
    elsif ( $len == 4 ) {
        return ( 0, $digit[0], $digit[1], $digit[3] );
    }
    elsif ( $len == 5 ) {
        return ( $digit[0], $digit[1], $digit[2], $digit[4] );
    }
}

sub Computer {
    my @code  = @{ $_[0] };
    my $input = $_[1];

    # print "input:$input\n";
    # print "code:@code\n";
    # $code[1] = 12;
    # $code[2] = 2;
    my $ptr = 0;
    my $step = 4;
    my $steps = 0;

    # my $input = 7;
    my $ptrinc;

    # print "second:$code[1]\n";
    while (1) {
        my $opcode = $code[$ptr];
        if ( $opcode == 99  ) {
            last;
        }
        my @modes = GetParameterModes($opcode);
        # print "================ step: $steps\n";
        # print "modes:@modes\n";
        my $inst = $modes[3];
        my $p1   = $modes[2];
        my $p2   = $modes[1];
        my $p3   = $modes[0];

        # print "inst:$inst\n";
        # print "p1:$p1\n";
        # print "p2:$p2\n";

        if ( $inst == 1 ) {

            # Add two parameters together
            my $first =
              $p1 == 0 ? $code[ $code[ $ptr + 1 ] ] : $code[ $ptr + 1 ];

            # print "first:$first\n";
            my $second =
              $p2 == 0 ? $code[ $code[ $ptr + 2 ] ] : $code[ $ptr + 2 ];

            # print "second:$second\n";
            my $out = $code[ $ptr + 3 ];
            $code[$out] = $first + $second;
            $ptrinc = 4;
        }
        elsif ( $inst == 2 ) {

            # Multiply two parameters
            my $first =
              $p1 == 0 ? $code[ $code[ $ptr + 1 ] ] : $code[ $ptr + 1 ];

            # print "first:$first - $code[$code[$ptr+1]]\n";
            my $second =
              $p2 == 0 ? $code[ $code[ $ptr + 2 ] ] : $code[ $ptr + 2 ];
            my $out = $code[ $ptr + 3 ];
            $code[$out] = $first * $second;
            $ptrinc = 4;
        }
        elsif ( $inst == 3 ) {

            # Read an input
            my $out = $code[ $ptr + 1 ];
            $code[$out] = $input;
            $ptrinc = 2;
        }
        elsif ( $inst == 4 ) {

            # Ouput at position
            my $out = $p1 == 0 ? $code[$code[$ptr+1]] : $code[ $ptr + 1 ];
            # print "@code\n";
            print "[*] $out\n";
            $ptrinc = 2;
        }
        elsif ( $inst == 5 ) {

            # Jump if True
            my $first =
              $p1 == 0 ? $code[ $code[ $ptr + 1 ] ] : $code[ $ptr + 1 ];
            my $second =
              $p2 == 0 ? $code[ $code[ $ptr + 2 ] ] : $code[ $ptr + 2 ];
            # $ptr    = $first != 0 ? $second : $ptr;
            if ($first != 0){
              $ptr = $second;
              $ptrinc = 0;
            }else{
              $ptrinc = 3;
            }
            # print "@code\n";
        }
        elsif ( $inst == 6 ) {

            # Jump if False
            my $first = $p1 == 0 ? $code[ $code[ $ptr + 1 ] ] : $code[ $ptr + 1 ];
            # print "$p1:$p2\n";
            my $second = $p2 == 0 ? $code[ $code[ $ptr + 2 ] ] : $code[ $ptr + 2 ];
            # print "$second: $code[ $code[ $ptr + 2 ] ]\n";
            # $ptr    = $first == 0 ? $second : $ptr;
            if ($first == 0){
              $ptr = $second;
              $ptrinc = 0;
            }else{
              $ptrinc = 3;
            }
            # print "2nd - $first : ptr - $ptr\n";
            
        }
        elsif ( $inst == 7 ) {

            # Less than
            my $first =
              $p1 == 0 ? $code[ $code[ $ptr + 1 ] ] : $code[ $ptr + 1 ];
            my $second =
              $p2 == 0 ? $code[ $code[ $ptr + 2 ] ] : $code[ $ptr + 2 ];
            my $out = $code[ $ptr + 3 ];
            $code[$out] = $first < $second ? 1 : 0;
            $ptrinc = 4;
        }
        elsif ( $inst == 8 ) {

            # Equals
            my $first =
              $p1 == 0 ? $code[ $code[ $ptr + 1 ] ] : $code[ $ptr + 1 ];
            my $second =
              $p2 == 0 ? $code[ $code[ $ptr + 2 ] ] : $code[ $ptr + 2 ];
            my $out = $code[ $ptr + 3 ];
            $code[$out] = $first == $second ? 1 : 0;
            $ptrinc = 4;
            
        }
        $ptr += $ptrinc;
        $steps++;
        # print "@code\n";
        # print "ptr:$code[$ptr] - $ptr\n";
    }

    
}

sub GenTestInput {
    return split /,/, <@_>;
}

my $file = "day5.txt";
open my $data, $file or die "Could not open file: $!";
my $test = "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99";
# my $test = "3,3,1105,-1,9,1101,0,0,12,4,12,99,1";
# print "test : 1 - out: ";
# my @input = split /,/,<$data>;
my @input = GenTestInput($test);
Computer( \@input, 9 );

# @input = GenTestInput($file);
# Computer(\@input,5);

# $equals8test = "3,3,1108,-1,8,3,4,3,99";
# print "test : 1 - out: ";
# @input = GenTestInput($equals8test);
# Computer( \@input, 8 );

# my @code = split /,/,<$data>;

# print "@code\n";

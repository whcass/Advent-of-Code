#!/usr/bin/perl
package AdventOfCode::ComputerObj;

use strict;
use warnings;
use Data::Dumper;
use Exporter qw(import);

our @EXPORT_OK = qw (new run);

sub new {
  my $class = shift;
  my @code = @{$_[0]};
  # print "@code";
  my $self = bless{
    code => [@code],
    ptr =>0,
    mem=>[],
  }, $class;
  return $self;
}

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

sub run {
  my $self = shift;
  my $input = $_[0];
  $self->Computer($input);
  return @{$self->@{mem}};
}

sub Computer {
    my $self = shift;
    my @code  = @{$self->@{code}};
    my $input = $_[0];
    # print "input:$input";
    my $ptr = $self->{ptr};
    my $step = 4;
    my $steps = 0;
    my $ptrinc;
    my @output;
    my $inputptr = 0;
    while (1) {
        my $opcode = $code[$ptr];
        if ( $opcode == 99  ) {
            push @{$self->@{mem}}, 99;
            last;
        }
        my @modes = GetParameterModes($opcode);
        # Debug help
        # print "================ step: $steps\n";
        # print "modes:@modes\n";
        my $inst = $modes[3];
        my $p1   = $modes[2];
        my $p2   = $modes[1];
        my $p3   = $modes[0];

        if ( $inst == 1 ) {

            # Add two parameters together
            my $first =
              $p1 == 0 ? $code[ $code[ $ptr + 1 ] ] : $code[ $ptr + 1 ];
            my $second =
              $p2 == 0 ? $code[ $code[ $ptr + 2 ] ] : $code[ $ptr + 2 ];
            my $out = $code[ $ptr + 3 ];
            $code[$out] = $first + $second;
            $ptrinc = 4;
        }
        elsif ( $inst == 2 ) {

            # Multiply two parameters
            my $first =
              $p1 == 0 ? $code[ $code[ $ptr + 1 ] ] : $code[ $ptr + 1 ];
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
            $inputptr++;
            $ptrinc = 2;
        }
        elsif ( $inst == 4 ) {

            # Ouput at position
            my $out = $p1 == 0 ? $code[$code[$ptr+1]] : $code[ $ptr + 1 ];
            # print "[*] $out\n";
            push @{$self->@{mem}}, $out;
            $ptr += 2;
            $self->{ptr} = $ptr;
            return;
            # $ptrinc = 2;
        }
        elsif ( $inst == 5 ) {

            # Jump if True
            my $first =
              $p1 == 0 ? $code[ $code[ $ptr + 1 ] ] : $code[ $ptr + 1 ];
            my $second =
              $p2 == 0 ? $code[ $code[ $ptr + 2 ] ] : $code[ $ptr + 2 ];
            if ($first != 0){
              $ptr = $second;
              $ptrinc = 0;
            }else{
              $ptrinc = 3;
            }
        }
        elsif ( $inst == 6 ) {

            # Jump if False
            my $first = $p1 == 0 ? $code[ $code[ $ptr + 1 ] ] : $code[ $ptr + 1 ];
            # print "first:$code[$code[$ptr+1]]\n";
            # print "@code\n";
            my $second = $p2 == 0 ? $code[ $code[ $ptr + 2 ] ] : $code[ $ptr + 2 ];
            if ($first == 0){
              $ptr = $second;
              $ptrinc = 0;
            }else{
              $ptrinc = 3;
            }
            
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
    }

    return @output;
}


1;
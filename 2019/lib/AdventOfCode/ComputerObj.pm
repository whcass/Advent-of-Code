#!/usr/bin/perl
package AdventOfCode::ComputerObj;

use strict;
use warnings;
use Data::Dumper;
use Exporter qw(import);

our @EXPORT_OK = qw (new run);

sub new {
    my $class = shift;
    my @code  = @{ $_[0] };
    my @array = (0) x 5000;
    @code = ( @code, @array );

    my $self = bless {
        code => [@code],
        ptr  => 0,
        mem  => [],
        base => 0,
    }, $class;
    return $self;
}

sub GetParamFromMode {
    my $mode = $_[0];
    my $ptr  = $_[1];
    my $base = $_[2];
    my @code = @{ $_[3] };

    if ( $mode == 0 ) {
        return $code[ $code[$ptr] ];
    }
    elsif ( $mode == 1 ) {
        return $code[$ptr];
    }
    elsif ( $mode == 2 ) {
        return $code[ $code[$ptr] + $base ];
    }
}

sub GetOutParamFromMode {
    my $mode = $_[0];
    my $ptr  = $_[1];
    my $base = $_[2];
    my @code = @{ $_[3] };

    if ( $mode == 2 ) {
        return $code[$ptr] + $base;
    }
    else {
        return $code[$ptr];
    }
}

sub CalcInstruction {
    my $base   = $_[0];
    my $ptr    = $_[1];
    my @modes  = @{ $_[2] };
    my @code   = @{ $_[3] };
    my $first  = GetParamFromMode( $modes[2], $ptr + 1, $base, \@code );
    my $second = GetParamFromMode( $modes[1], $ptr + 2, $base, \@code );
    my $out    = GetOutParamFromMode( $modes[0], $ptr + 3, $base, \@code );

    return ( $first, $second, $out );
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
    my $self  = shift;
    my $input = $_[0];
    $self->Computer($input);
    return @{ $self->@{mem} };
}

sub Computer {
    my $self  = shift;
    my @code  = @{ $self->@{code} };
    my $input = $_[0];

    my $ptr = $self->{ptr};

    my $step  = 4;
    my $steps = 0;
    my $ptrinc;
    my @output;
    my $inputptr = 0;
    while (1) {
        my $opcode = $code[$ptr];
        if ( $opcode == 99 ) {
            push @{ $self->@{mem} }, 99;
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

            my ( $first, $second, $out ) =
              CalcInstruction( $self->{base}, $ptr, \@modes, \@code );

            $code[$out] = $first + $second;
            $ptr += 4;
        }
        elsif ( $inst == 2 ) {

            # Multiply two parameters
            my ( $first, $second, $out ) =
              CalcInstruction( $self->{base}, $ptr, \@modes, \@code );

            $code[$out] = $first * $second;
            $ptr += 4;
        }
        elsif ( $inst == 3 ) {

            # Read an input
            my $out =
              GetOutParamFromMode( $p1, $ptr + 1, $self->{base}, \@code );

            $code[$out] = $input;
            $inputptr++;
            $ptr += 2;
        }
        elsif ( $inst == 4 ) {

            # Ouput at position
            # Use the original method for this, because reasons
            my $out = GetParamFromMode( $p1, $ptr + 1, $self->{base}, \@code );
            push @{ $self->@{mem} }, $out;
            $ptr += 2;
        }
        elsif ( $inst == 5 ) {

            # Jump if True
            my ( $first, $second ) =
              CalcInstruction( $self->{base}, $ptr, \@modes, \@code );
            $ptr = $first!=0?$second:($ptr+=3);
        }
        elsif ( $inst == 6 ) {

            # Jump if False
            my ( $first, $second ) =
              CalcInstruction( $self->{base}, $ptr, \@modes, \@code );
              $ptr = $first==0?$second:($ptr+=3);
        }
        elsif ( $inst == 7 ) {

            # Less than
            my ( $first, $second, $out ) =
              CalcInstruction( $self->{base}, $ptr, \@modes, \@code );

            $code[$out] = $first < $second ? 1 : 0;
            $ptr += 4;
        }
        elsif ( $inst == 8 ) {

            # Equals
            my ( $first, $second, $out ) =
              CalcInstruction( $self->{base}, $ptr, \@modes, \@code );

            $code[$out] = $first == $second ? 1 : 0;
            $ptr += 4;
        }
        elsif ( $inst == 9 ) {
            my $first =
              GetParamFromMode( $p1, $ptr + 1, $self->{base}, \@code );
              
            $self->{base} += $first;
            $ptr += 2;
        }

        # $steps++;
    }

    return @output;
}

1;

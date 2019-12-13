package AdventOfCode::Planet;
use strict;
use warnings;
use Exporter qw(import);

our @EXPORT_OK = qw(new update printOut getEnergy);

sub new {
    my $class = shift;
    my $x     = $_[0];
    my $y     = $_[1];
    my $z     = $_[2];
    my @pos;
    $pos[0] = $x;
    $pos[1] = $y;
    $pos[2] = $z;
    my @vel;
    $vel[0] = $vel[1] = $vel[2] = 0;

    # print $pos{'x'};
    my $self = bless {
        pos => [@pos],
        vel => [@vel],
    }, $class;

    return $self;
}

sub getEnergy {
    my $self = shift;
    my @pos = @{$self->@{pos}};
    my @vel = @{$self->@{vel}};

    return (abs($pos[0]) + abs($pos[1]) + abs($pos[2])) * (abs($vel[0]) + abs($vel[1]) + abs($vel[2]));
}

sub printOut {
    my $self = shift;
    my @pos  = @{ $self->@{pos} };
    my @vel  = @{ $self->@{vel} };
    print
"pos=<x=$pos[0], y=$pos[1], z=$pos[2]>, vel=<x=$vel[0], y=$vel[1], z=$vel[2]>\n";
}

sub update {
    my $self         = shift;
    my @otherPlanets = @{ $_[0] };
    my ( $lhsX, $lhsY, $lhsZ ) = @{ $self->@{pos} };
    for (@otherPlanets) {
        my $obj = $_;
        if ( $self == $obj ) {
            next;
        }
        my ( $rhsX, $rhsY, $rhsZ ) = @{ $obj->@{pos} };

        # my @objvel = @{$obj->@{vel}};
        if ( $lhsX > $rhsX ) {
            @{ $self->@{vel} }[0]--;
            @{ $obj->@{vel} }[0]++;
        }
        elsif($lhsX == $rhsX){}
        else {
            @{ $self->@{vel} }[0]++;
            @{ $obj->@{vel} }[0]--;
        }

        if ( $lhsY > $rhsY ) {
            @{ $self->@{vel} }[1]--;
            @{ $obj->@{vel} }[1]++;
        }
        elsif($lhsY == $rhsY){}
        else {
            @{ $self->@{vel} }[1]++;
            @{ $obj->@{vel} }[1]--;
        }

        if ( $lhsZ > $rhsZ ) {
            @{ $self->@{vel} }[2]--;
            @{ $obj->@{vel} }[2]++;
        }
        elsif($lhsZ == $rhsZ){}
        else {
            @{ $self->@{vel} }[2]++;
            @{ $obj->@{vel} }[2]--;
        }

    }
    my @vel = @{ $self->@{vel} };
    @{ $self->@{pos} }[0] += $vel[0];
    @{ $self->@{pos} }[1] += $vel[1];
    @{ $self->@{pos} }[2] += $vel[2];
}

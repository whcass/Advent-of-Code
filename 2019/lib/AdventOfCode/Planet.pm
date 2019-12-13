package AdventOfCode::Planet;
use strict;
use warnings;
use Exporter qw(import);

our @EXPORT_OK = qw(new);
sub new {
    my $class = shift;
    my $x = $_[0];
    my $y = $_[1];
    my $z = $_[2];
    my %pos;
    $pos{'x'} = $x;
    $pos{'y'} = $y;
    $pos{'z'} = $z;
    my %vel;
    $vel{'x'} = $vel{'y'} = $vel{'z'} = 0;
    my $self = bless {
        pos=>%pos,
        vel=>%vel,
    }, $class;

    return $self;
}
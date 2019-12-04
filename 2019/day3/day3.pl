use strict;
use warnings;
use Data::Dumper;
use Array::Utils qw(:all);

sub BuildInstructions {
    my @wire = @_;
    my @res;

    for ( my $i = 0 ; $i < @wire ; $i++ ) {
        my $inst = $wire[$i];
        my $dir  = substr $inst, 0, 1;
        my $num  = substr $inst, 1;
        push( @res, $dir ) for 1 .. $num;
    }

    return @res;
}

sub ShortestDistance {
    my @list = @_;
    my $shortest = 999999999999999999;
    foreach my $point (@list){
        my @coord = split /,/,$point;
        my $x = $coord[0];
        my $y = $coord[1];
        my $dist = abs($x)+abs($y);
        $shortest = $dist if $dist < $shortest;
    }

    print $shortest;
}

sub findPath {
    my @inst = @_;
    my @res;
    my $x = 0;
    my $y = 0;
    for ( my $i = 0 ; $i < @inst ; $i++ ) {
        my $dir = $inst[$i];
        if($dir eq 'R'){
            $x++;
            push(@res,"$x,$y");
        }elsif($dir eq 'L'){
            $x--;
            push(@res,"$x,$y");
        }elsif($dir eq 'U'){
            $y++;
            push(@res,"$x,$y");
        }elsif($dir eq 'D'){
            $y--;
            push(@res,"$x,$y");
        }
    }
    return @res;
    # print("$res[0]\n");
}

my $filea = "day3a.txt";
my $fileb = "day3b.txt";
open my $wirea, $filea or die "Could not open file: $!";
open my $wireb, $fileb or die "Could not open file: $!";

my @one = split /,/, <$wirea>;
my @two = split /,/, <$wireb>;

@one = findPath( BuildInstructions(@one) );
@two = findPath( BuildInstructions(@two) );
my @intersect = intersect(@one,@two);
ShortestDistance(@intersect);


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

sub manhattanDistance {
    my %list = @_;
    my $shortest = 999999999999999999;
    foreach my $point (keys %list){
        my @coord = split /,/,$point;
        
        my $x = $coord[0];
        my $y = $coord[1];
        my $dist = abs($x)+abs($y);
        $shortest = $dist if $dist < $shortest;
    }

    print "$shortest\n";
}

sub signalDistance {
    my %list = @_;
    my $shortest = 9999999999;
    foreach (keys %list){
        $shortest = $list{$_} if $list{$_} < $shortest;
        # print "$list{$_}\n";
    }
    print "$shortest\n";
}

sub findPath {
    my @inst = @_;
    my %res;
    my $x = 0;
    my $y = 0;
    for ( my $i = 0 ; $i < @inst ; $i++ ) {
        my $steps = $i + 1;
        my $dir = $inst[$i];
        if($dir eq 'R'){
            $x++;
            $res{"$x,$y"} = $steps;
        }elsif($dir eq 'L'){
            $x--;
            $res{"$x,$y"} = $steps;
        }elsif($dir eq 'U'){
            $y++;
            $res{"$x,$y"} = $steps;
        }elsif($dir eq 'D'){
            $y--;
            $res{"$x,$y"} = $steps;
        }
    }
    return %res;
    # print("$res[0]\n");
}

my $filea = "day3a.txt";
my $fileb = "day3b.txt";
open my $wirea, $filea or die "Could not open file: $!";
open my $wireb, $fileb or die "Could not open file: $!";

my @one = split /,/, <$wirea>;
my @two = split /,/, <$wireb>;

my %oneHash = findPath( BuildInstructions(@one) );
my %twoHash = findPath( BuildInstructions(@two) );
my %intersect;
foreach(keys %oneHash){
    if (exists $twoHash{$_}){
        $intersect{$_} = $twoHash{$_}+$oneHash{$_};
    }
}
# print Dumper(%intersect);
manhattanDistance(%intersect);
signalDistance(%intersect);

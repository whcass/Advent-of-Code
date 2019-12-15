use strict;
use warnings;
use AdventOfCode::Planet;
use List::Util 'first';
use Math::Utils 'lcm';
# <x=8, y=0, z=8>
# <x=0, y=-5, z=-10>
# <x=16, y=10, z=-5>
# <x=19, y=-10, z=-7>
my $A = AdventOfCode::Planet->new(8,0,8);
my $B = AdventOfCode::Planet->new(0,-5,-10);
my $C = AdventOfCode::Planet->new(16,10,-5);
my $D = AdventOfCode::Planet->new(19,-10,-7);
# <x=-1, y=0, z=2>
# <x=2, y=-10, z=-7>
# <x=4, y=-8, z=8>
# <x=3, y=5, z=-1>
# my $A = AdventOfCode::Planet->new( -1, 0,   2 );
# my $B = AdventOfCode::Planet->new( 2,  -10, -7 );
# my $C = AdventOfCode::Planet->new( 4,  -8,  8 );
# my $D = AdventOfCode::Planet->new( 3,  5,   -1 );
my @otherPlanets = ( $A, $B, $C, $D );
# print "After 0 steps:\n";
# for (@otherPlanets) {
    # $_->printOut();
# }
# print "\n";
for my $index ( 1 .. 1000 ) {
    my @temp = @otherPlanets;
    # print "After $index steps:\n";
    shift @temp;
    $A->update( \@temp );
    shift @temp;
    $B->update( \@temp );
    shift @temp;
    $C->update( \@temp );
    shift @temp;
    $D->update( \@temp );
    # for (@otherPlanets) {
        # $_->printOut();
    # }
    # print "\n";
}
my $total = 0;
for my $planet (@otherPlanets){
    $total += $planet->getEnergy();
}

print "[*] Part1 - Total Energy: $total\n";

$A = AdventOfCode::Planet->new(8,0,8);
$B = AdventOfCode::Planet->new(0,-5,-10);
$C = AdventOfCode::Planet->new(16,10,-5);
$D = AdventOfCode::Planet->new(19,-10,-7);
# $A = AdventOfCode::Planet->new( -1, 0,   2 );
# $B = AdventOfCode::Planet->new( 2,  -10, -7 );
# $C = AdventOfCode::Planet->new( 4,  -8,  8 );
# $D = AdventOfCode::Planet->new( 3,  5,   -1 );
@otherPlanets = ( $A, $B, $C, $D );

my @states;
my $steps = 0;
my @factors;
while(1){
    my @temp = @otherPlanets;
    # print "After $index steps:\n";
    shift @temp;
    $A->updateX( \@temp );
    shift @temp;
    $B->updateX( \@temp );
    shift @temp;
    $C->updateX( \@temp );
    shift @temp;
    $D->updateX( \@temp );

    my $checksum = $A->getCheckSumX() . $B->getCheckSumX() . $C->getCheckSumX() . $D->getCheckSumX();
    
    my @match = grep{/$checksum/} @states;
    
    if (@match > 0){
        push @factors, $steps;
        last;
    }else{
        push @states, $checksum;
    }
    
    # last;
    $steps++;
}

# 18 28 44
# my $factor = @factors;
print "[*] Part2 - $factors[0]\n";

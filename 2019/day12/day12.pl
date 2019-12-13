use strict;
use warnings;
use AdventOfCode::Planet;

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
print "After 0 steps:\n";
for (@otherPlanets) {
    $_->printOut();
}
print "\n";
for my $index ( 1 .. 1000 ) {
    my @temp = @otherPlanets;
    print "After $index steps:\n";
    shift @temp;
    $A->update( \@temp );
    shift @temp;
    $B->update( \@temp );
    shift @temp;
    $C->update( \@temp );
    shift @temp;
    $D->update( \@temp );
    for (@otherPlanets) {
        $_->printOut();
    }
    print "\n";
}
my $total = 0;
for my $planet (@otherPlanets){
    $total += $planet->getEnergy();
}

print "[*] Total Energy: $total\n";

use strict;
use warnings;
no warnings 'recursion';
use Data::Dumper;

# use Data::TreeDumper;
use Tree::Simple;

my $input = "COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L";
my $file = "day6.txt";
open my $data, $file or die "Could not open file: $!";
my %bodies;
my $bodyCount = 0;
foreach my $line (<$data>) {

    # chomp $line;
    my ( $body, $orbiter ) = $line =~ /(.*)\)(.*)/;

    # print "$orbiter orbits $body\n";
    if ( !$bodies{$body} ) {
        # print "[*] body $body does not exist, making a node\n";
        $bodies{$body} = Tree::Simple->new( $body, Tree::Simple->ROOT );
        $bodies{$body}->addChild( Tree::Simple->new($orbiter) );
        
    }
    else {
        # print "[*] adding $orbiter to $body\n";
        $bodies{$body}->addChild( Tree::Simple->new($orbiter) );
    }
    $bodyCount++;
    # my @children = @{$tree{$parent}{'children'}};
    # $tree{$parent}{'children'} = push @children,$child;
}

my $root  = $bodies{'COM'};
my $total = 0;

# print Dumper(%bodies);
# print Dumper($root->getAllChildren);
foreach my $child ( $root->getAllChildren ) {
    count( $child->getNodeValue, 0 );
}

sub count {
    my $node  = $bodies{ $_[0] };
    my $depth = $_[1];
    $depth++;
    
    my $bodyName = $node->getNodeValue;

    # print "[+] $bodyName - D:1 - I:$depth = $depth\n";
    # print "body: $name\n";
    foreach my $child ( $node->getAllChildren ) {
        my $name = $child->getNodeValue;
        print "[*] $name:$depth\n";
        $total += $depth;
        # print "[*] $bodyName has orbiting:$name\n\t$depth\n";
        if ( $bodies{$name} ) {
            count( $name, $depth );
        }
    }
}
print "[*] indirects: $total\n";
print "[*] bodyCount: $bodyCount\n";
my $count = $total + $bodyCount;
print "[*] total: $count\n";

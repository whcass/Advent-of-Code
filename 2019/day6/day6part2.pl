use strict;
use warnings;
no warnings 'recursion';
use Data::Dumper;
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

    my ( $body, $orbiter ) = $line =~ /(.*)\)(.*)/;

    if ( !$bodies{$body} ) {
        $bodies{$body} = Tree::Simple->new( $body, Tree::Simple->ROOT );
        $bodies{$body}->addChild( Tree::Simple->new($orbiter) );
    }
    else {
        $bodies{$body}->addChild( Tree::Simple->new($orbiter) );
    }
    $bodyCount++;
}

my $root  = $bodies{'COM'};
my $total = 0;

my @pathToYou = findPath($root,"YOU");
my @pathToSanta = findPath($root,'SAN');
my %repeats;
my @diff;
for (@pathToYou,@pathToSanta){ $repeats{$_}++}
for (keys %repeats){
    push @diff, $_ unless $repeats{$_} > 1;
}
my $steps = @diff;
$steps -= 2;
print "[*] $steps\n";

sub findPath
{
    my $who = $_[1];
    my $root = $_[0];

    my ($ignore,@path) = findWho($root->getNodeValue,$who);
    @path = reverse(@path);
    return @path;
}


sub findWho
{
    my $node = $bodies{$_[0]};
    my $who = $_[1];
    my @path;
    foreach my $child ($node->getAllChildren){
        if($bodies{$child->getNodeValue}){
            my($record,@path) = findWho($child->getNodeValue,$who);
            if($record){
                push @path, $child->getNodeValue;
                return 1, @path;
            }
        }else{
            if($child->getNodeValue eq $who){
                return (1,$child->getNodeValue);
            }
        }
    }
}

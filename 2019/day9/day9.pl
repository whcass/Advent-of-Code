use strict;
use warnings;
use File::Basename qw(dirname);
use Cwd qw(abs_path);
use lib dirname( dirname abs_path $0) . '/lib';
use AdventOfCode::ComputerObj qw(new run);

my $file = "day9.txt";
open my $data, $file or die "Could not open file: $!";
my @input = split /,/, <$data>;

my $comp = AdventOfCode::ComputerObj->new( \@input );

my $start = time;
my @out = $comp->run(1);
my $end = time;
my $elapsed = $end-$start;
print "[*] Finished, time elapsed:$elapsed"."s\n";
pop @out;

# print "@out";
if ( $out[0] == 2870072642 ) {
    print "[*] Pass\n";
}
else {
    print "[*] FAIL : got:$out[0] - want:2870072642\n";
}
my $compB = AdventOfCode::ComputerObj->new(\@input);
$start = time;
@out = $compB->run(2);
$end = time;
$elapsed = $end-$start;
print "[*] Finished, time elapsed:$elapsed"."s";
pop @out;
if ( $out[0] == 58534 ) {
    print "[*] Pass\n";
}
else {
    print "[*] FAIL : got:$out[0] - want:58534\n";
}
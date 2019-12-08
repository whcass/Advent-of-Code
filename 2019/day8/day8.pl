use strict;
use warnings;

my $file = "day8.txt";
open my $data, $file or die "Could not open file: $!";
my @pixels   = split //, <$data>;
my $width    = 25;
my $height   = 6;
my $least    = 99999999999999999;
my $sum      = 0;
my $perlayer = $width * $height;
my @layers;
push @layers, [splice @pixels,0,$perlayer] while @pixels;
my $layercount = @layers;
for my $ilayer (0..$layercount-2){
    my %layerHash;
    my @layer = @{$layers[$ilayer]};
    $layerHash{$_}++ for (@layer);
    print "@{[%layerHash]}\n";
    if($layerHash{0} < $least){
        print "[*] badgers\n";
        $least = $layerHash{0};
        $sum = $layerHash{1} * $layerHash{2};
    }
}

print "[*] sum:$sum\n";
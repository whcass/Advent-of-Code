use strict;
use warnings;

my $file = "day8.txt";
open my $data, $file or die "Could not open file: $!";
close $file;
my @pixels   = split //, <$data>;
my $width    = 25;
my $height   = 6;
my $perlayer = $width * $height;
my @layers;

push @layers, [splice @pixels,0,$perlayer] while @pixels;
my @image;
my $pixelLength = @pixels;

@layers = reverse(@layers);
for my $layerIndex (0..@layers-1){
    my @layer = @{$layers[$layerIndex]};
    for my $i( 0..$perlayer-1){

        if ($layer[$i]!=2){
            @image[$i] = $layer[$i];
        }
    }
}

@image = map{$_ =~ s/0/ /g; $_ } @image;
@image = map{$_ =~ s/1/#/g; $_ } @image;

for my $pix (0..$perlayer-1){
    if ($pix % $width == 0){
        print "\n";
    }
    print "$image[$pix] ";
}
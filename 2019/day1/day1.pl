#!/usr/bin/perl
use strict;
use warnings;
my $file = "day1.txt";
open my $data, $file, or die "Could not open $file: $!";
my $output = 0;
while(my $line = <$data>){
    my $value = int($line/3)-2;
    $output += $value;
    my $val = $value;
    while($val > 0){
        $val = int($val/3)-2;
        if($val > 0)
        {
            $output+=$val;
        }
    }
}
print "[*] 4941547\n";
print "[*] $output\n";
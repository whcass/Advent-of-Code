#!/usr/bin/perl
use strict;
use warnings;
my $file = "day1.txt";
open my $data, $file, or die "Could not open $file: $!";
my $output = 0;
while(my $line = <$data>){
    my @matches = $line =~ /([\+|-])(\d+)/;
    $output += $line;
    
}

print "[*] $output";
#!/usr/bin/perl
use strict;
use warnings;
use List::Util 1.33 'any';
my $file = "day1.txt";
open my $data, $file, or die "Could not open $file: $!";
my $output = 0;
my @frequencies;
my $currentFrequency = 0;
while(my $line = <$data>){
    my @matches = $line =~ /([\+|-])(\d+)/;
    $currentFrequency += $line;
    print "[*] $currentFrequency\n";
    my $matchFound = any{$currentFrequency} @frequencies;
    
    if($matchFound){
        $output = $currentFrequency;
        last;
    }
    
    push(@frequencies, $currentFrequency);


}

print "[*] $output\n";
#!/usr/bin/perl
use strict;
use warnings;
my $file = "day4.txt";
open my $data, $file, or die "Could not open $file: $!";
while(my $line = <$data>){
    my $matches = () = $line =~ /(\w+).+\1/i;
    
    print "[*] $line - $matches\n" if $matches==0;
}
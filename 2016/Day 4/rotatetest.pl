use strict;
use warnings;

use Data::Dumper qw(Dumper);

my $letter = 'a';
my $num = 27;
my %alphaNums;
my %numAlphas;
my $index = 0;
foreach('a'..'z'){
        $alphaNums{$_} = $index;
        $numAlphas{$index} = $_;
        $index++;
}



print Dumper \%numAlphas;
my $letterEncrypt = ($alphaNums{$letter} + $num) % 26;
print "$numAlphas{$letterEncrypt} - $letter";

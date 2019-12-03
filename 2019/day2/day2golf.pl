open my $d, "day2.txt" or die;
my @c = split /,/, <$d>;
$c[1] = 12;
$c[2] = 2;
my $p = 0;
while () {
    my $i = $c[$p];
    $i =~ s/99/last/e;
    $i =~ s/1/$c[$c[$p+3]]=$c[$c[$p+1]]+$c[$c[$p+2]]/e;
    $i =~ s/2/$c[$c[$p+3]]=$c[$c[$p+1]]*$c[$c[$p+2]]/e;
    $p += 4;
}
print "$c[0]\n";

my $file = "day2.txt";
open my $data, $file or die "Could not open file: $!";
my @c = split /,/, <$data>;

#print "@c[0]\n";
my @ccopy = @c;
for ( my $noun = 0 ; $noun < 99 ; $noun++ ) {
    for ( my $verb = 0 ; $verb < 99 ; $verb++ ) {
        $c[1] = $noun;
        $c[2] = $verb;
        my $p = 0;
        while (1) {
            my $i = $c[$p];
            $i =~ s/99/last/e;
            $i =~ s/1/$c[$c[$p+3]]=$c[$c[$p+1]]+$c[$c[$p+2]]/e;
            $i =~ s/2/$c[$c[$p+3]]=$c[$c[$p+1]]*$c[$c[$p+2]]/e;
            $p += 4;
        }
        if ( $c[0] == 19690720 ) {
            my $res = 100 * $noun + $verb;
            print "[*] $c[0] : $noun : $verb : $res\n";
            exit;
        }
        @c = @ccopy;
    }
}

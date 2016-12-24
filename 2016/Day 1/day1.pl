use strict;
use warnings;

my $string = "R1, R1, R3, R1, R1, L2, R5, L2, R5, R1, R4, L2, R3, L3, R4, L5, R4, R4, R1, L5, L4, R5, R3, L1, R4, R3, L2, L1, R3, L4, R3, L2, R5, R190, R3, R5, L5, L1, R54, L3, L4, L1, R4, R1, R3, L1, L1, R2, L2, R2, R5, L3, R4, R76, L3, R4, R191, R5, R5, L5, L4, L5, L3, R1, R3, R2, L2, L2, L4, L5, L4, R5, R4, R4, R2, R3, R4, L3, L2, R5, R3, L2, L1, R2, L3, R2, L1, L1, R1, L3, R5, L5, L1, L2, R5, R3, L3, R3, R5, R2, R5, R5, L5, L5, R2, L3, L5, L2, L1, R2, R2, L2, R2, L3, L2, R3, L5, R4, L4, L5, R3, L4, R1, R3, R2, R4, L2, L3, R2, L5, R5, R4, L2, R4, L1, L3, L1, L3, R1, R2, R1, L5, R5, R3, L3, L3, L2, R4, R2, L5, L1, L1, L5, L4, L1, L1, R1";
my $orientation = 'up';
my $x = 0;
my $y = 0;
foreach my $x (split /,/, $string) {
  my @arr = split /(R|L)/, $x;
  my $rotate = @arr[1];

  my $distance = @arr[3];
  print $rotate,$distance;

  if($rotate eq 'R'){
    if($orientation eq'up'){
      $orientation = 'right';
      $x+=$distance;
    }elsif($orientation eq 'right'){
      $orientation = 'down';
    }elsif($orientation eq 'down'){
      $orientation = 'left';
      $x-=$distance;
    }elsif($orientation = 'left'){
      $orientation = 'up';
    }



  }elsif($rotate eq'L'){
    if($orientation eq'up'){
      $orientation = 'left';
    }elsif($orientation eq 'right'){
      $orientation = 'up';
    }elsif($orientation eq 'down'){
      $orientation = 'right';
    }elsif($orientation eq   'left'){
      $orientation = 'down';
    }
  }
}

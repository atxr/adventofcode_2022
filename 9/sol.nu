let input = (
  open input3 
  | lines
  | each {|it|
    split column " "
  }
)

mut visited = [[0 0]]
mut rope = []
let I = 10
for i in 1..($I + 1) {
  $rope = ($rope | append [[0 0]])
}

for line in $input {
  let n = ((($line).0).column2 | into int)
  for k in 1..($n) {
    mut d = (($line).0).column1
    for i in 0..($I - 1) {
      let H = ($rope | get $i)
      let T = ($rope | get ($i + 1))

      mut new_H = $H
      if "R" in $d {
        $new_H = ($new_H | update 0 (($new_H | get 0) + 1))
      } 
      if "L" in $d {
        $new_H = ($new_H | update 0 (($new_H | get 0) - 1))
      } 
      if "U" in $d {
        $new_H = ($new_H | update 1 (($new_H | get 1) + 1))
      }
      if "D" in $d {
        $new_H = ($new_H | update 1 (($new_H | get 1) - 1))
      }

      mut dx = ($new_H | get 0) - ($T | get 0)
      mut dy = ($new_H | get 1) - ($T | get 1)
      if $i == 0 && ([$dx $dy] | math abs | math max) >= 2 {
        $d = ""
        $dx = ($H | get 0) - ($T | get 0)
        $dy = ($H | get 1) - ($T | get 1)
        if $dx > 0 {
          $d = ($d | append "R")
        }
        if $dx < 0 {
          $d = ($d | append "L")
        }
        if $dy > 0 {
          $d = ($d | append "U")
        }
        if $dy < 0 {
          $d = ($d | append "D")
        }
        print "log" | append $d
      } else if ([$dx $dy] | math abs | math max) < 2 {
        $d = ""
      }

      $rope = ($rope | update $i $new_H)
    }

    let Tlast = ($rope | get ($I - 1))
    if ($Tlast not-in $visited) {
      $visited = ($visited | append [($Tlast)])
    }

    print $line
    print $k
    for el in $rope {
      print $el
    }
    print "\n"

  }


}

$visited | length

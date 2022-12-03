import strutils
import sequtils
import sugar

proc find_match[T](a: seq[T],b: seq[T]): seq[T] =
    var matches: seq[T]
    for letter in @a:
        let match = collect:
            for x in @b:
                if x == letter: x
        matches.add(match)
            
    return deduplicate(matches)

proc get_score(c: char): int = 
    if c.isLowerAscii(): int(c) - int('a') + 1
    else: int(c) - int('A') + 27


let 
    f = readFile("input")
    lines = f.split('\n')[0 .. ^2]

var score1 = 0
for rucksack in lines:
    let 
        n = rucksack.len
        comp1 = rucksack[0 ..< n div 2]
        comp2 = rucksack[n div 2 .. ^1]
    assert comp1 & comp2 == rucksack

    let matches = find_match(@comp1, @comp2)
    assert matches.len == 1
    score1 += get_score(matches[0])


var score2 = 0
for group in lines.distribute(lines.len div 3):
    echo group
    let
        a = group[0].toSeq
        b = group[1].toSeq
        c = group[2].toSeq
        matches = find_match(find_match(a, b), find_match(b, c))

    assert matches.len == 1
    score2 += get_score(matches[0])

echo "Score1: " & score1.intToStr
echo "Score2: " & score2.intToStr
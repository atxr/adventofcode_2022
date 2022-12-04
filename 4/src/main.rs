use std::str::FromStr;
use std::fs;

struct Pair {
    s1: u8,
    e1: u8,
    s2: u8,
    e2: u8
}

impl FromStr for Pair {
    type Err = std::num::ParseIntError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let tmp: Vec<Vec<u8>> = s.split(',').map(|r| r.split('-').map(|x| x.parse().unwrap()).collect()).collect();
        let s1: u8 = tmp[0][0];
        let e1: u8 = tmp[0][1];
        let s2: u8 = tmp[1][0];
        let e2: u8 = tmp[1][1];

        Ok(Pair {s1, e1, s2, e2})
    }
}

fn main() {
    let input = fs::read_to_string("./input")
        .expect("Should have been able to read the file");
    let lines: Vec<&str> = input.split('\n').filter(|l| l.len() != 0).collect();

    let mut score1: u32 = 0;
    let mut score2: u32 = 0;
    for line in lines {
        match Pair::from_str(line) {
            Ok(pair) => {
                if pair.s1 <= pair.s2 && pair.e2 <= pair.e1
                || pair.s2 <= pair.s1 && pair.e1 <= pair.e2 {
                    score1 += 1;
                    score2 += 1;
                }
                else if pair.s1 <= pair.s2 && pair.s2 <= pair.e1
                || pair.s1 <= pair.e2 && pair.e2 <= pair.e1 {
                    score2 += 1;
                }
            }
            Err(_) => {
                println!("Err when parsing {line}");
            }
        }
    }

    println!("Score 1: {score1}");
    println!("Score 2: {score2}");
}

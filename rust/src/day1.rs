use std::{collections::HashMap, iter::Iterator};

pub fn day1_part1(lines: impl Iterator<Item = String>) -> i64 {
    let (mut left, mut right) = parse_input(lines);
    left.sort();
    right.sort();
    let mut res = 0;
    for i in 0..left.len() {
        let dist = left[i] - right[i];
        res += dist.abs();
    }
    res
}

pub fn day1_part2(lines: impl Iterator<Item = String>) -> i64 {
    let (left, right) = parse_input(lines);
    let mut m = HashMap::new();
    for r in right.iter() {
        m.entry(r).and_modify(|e| *e += 1).or_insert(1);
    }
    let mut res = 0;
    for l in left.iter() {
        res += l * m.get(l).unwrap_or(&0);
    }
    res
}

fn parse_input(input: impl Iterator<Item = String>) -> (Vec<i64>, Vec<i64>) {
    let mut left = Vec::new();
    let mut right = Vec::new();
    for line in input {
        let mut split = line.split_whitespace();
        let (l, r) = (
            split.next().unwrap().parse::<i64>().unwrap(),
            split.next().unwrap().parse::<i64>().unwrap(),
        );
        left.push(l);
        right.push(r);
    }
    (left, right)
}

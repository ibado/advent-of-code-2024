use std::{
    fs::File,
    io::{BufRead, BufReader},
};

pub fn read_lines(day: u8) -> impl Iterator<Item = String> {
    assert!(day <= 25);
    let file_name = format!("../input/{day}.txt");
    let f = File::open(file_name).unwrap();
    BufReader::new(f).lines().map(|op| op.unwrap())
}

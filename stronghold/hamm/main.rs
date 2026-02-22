use std::fs;

fn main() {
    let content = fs::read_to_string("hamm.in").expect("Failed to read input file");
    let lines: Vec<&str> = content.lines().collect();

    if lines.len() != 2 {
        panic!("Expected exactly 2 lines in input file");
    }

    let (a, b) = (lines[0], lines[1]);
    let mut ret = 0;

    for (c1, c2) in a.chars().zip(b.chars()) {
        if c1 != c2 {
            ret += 1;
        }
    }
    println!("{}", ret);
}

use std::fs;

fn main() {
    let content = fs::read_to_string("subs.in").expect("Failed to read input file");
    let lines: Vec<&str> = content.lines().collect();

    if lines.len() != 2 {
        panic!("Expected exactly 2 lines in input file");
    }

    let s = lines[0];
    let q = lines[1];

    let mut ret = Vec::new();
    for i in 0..s.len() {
        if s.chars().nth(i) == q.chars().nth(0) {
            for j in 0..q.len() {
                if i + j >= s.len() || s.chars().nth(i + j) != q.chars().nth(j) {
                    break;
                }
                if j == q.len() - 1 {
                    ret.push(i + 1);
                }
            }
        }
    }

    for r in ret {
        print!("{} ", r);
    }
    println!();
}

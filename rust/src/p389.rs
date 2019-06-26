struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn find_the_difference(s: String, t: String) -> char {
        let mut i: u8 = 0;
        for c in s.chars() {
            i ^= c as u8
        }
        for c in t.chars() {
            i ^= c as u8
        }
        i as char
    }

    pub fn find_the_difference_with_map(s: String, t: String) -> char {
        let mut hmap: HashMap<char, i32> = HashMap::new();

        for c in s.chars() {
            if hmap.contains_key(&c) {
                hmap.insert(c, hmap.get(&c).unwrap() + 1);
            } else {
                hmap.insert(c, 1);
            }
        }

        for c in t.chars() {
            if !hmap.contains_key(&c) {
                return c;
            } else {
                let count = *hmap.get(&c).unwrap();
                if count == 0 {
                    return c;
                }
                hmap.insert(c, count - 1);
            }
        }
        return 'a';
    }
}

pub fn run() {
    let r = Solution::find_the_difference(String::from("aaaa"), String::from("aaaae"));
    println!("{}", r);
}

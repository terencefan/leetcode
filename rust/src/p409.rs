struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn longest_palindrome(s: String) -> i32 {
        let mut map: HashMap<char, i32> = HashMap::new();

        for c in s.chars() {
            if map.contains_key(&c) {
                map.insert(c, map.get(&c).unwrap() + 1);
            } else {
                map.insert(c, 1);
            }
        }

        let mut r = 0;
        let mut flag = false;
        for (_, count) in map {
            if count > 1 {
                r += count / 2 * 2;
            }
            let c = count % 2;
            if c == 1 {
                flag = true
            }
        }
        r + if flag {1} else {0}
    }
}

pub fn run() {
    println!("{}", Solution::longest_palindrome(String::from("abccccdd")))
}

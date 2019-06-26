struct Solution;

impl Solution {
    pub fn valid_word_abbreviation(word: String, abbr: String) -> bool {
        let mut num: i32 = 0;
        let mut i = 0;
        for c in abbr.chars() {
            match c {
                'a'...'z' => {
                    i += num as usize;
                    if i >= word.len() {
                        return false;
                    } else if word.chars().nth(i).unwrap() != c {
                        return false;
                    }
                    i += 1;
                    num = 0;
                },
                '0'...'9' => {
                    if num == 0 && c == '0' {
                        return false;
                    }
                    num *= 10;
                    num += (c as u8 - '0' as u8) as i32;
                },
                _ => break
            }
        }
        i + num as usize == word.len()
    }
}

pub fn run() {
    let r = Solution::valid_word_abbreviation(
        String::from("internationalization"),
        String::from("i5a11o1")
    );
    println!("{}", r);
}

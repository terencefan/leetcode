struct Solution;

impl Solution {
    pub fn valid_word_square(words: Vec<String>) -> bool {
        if words.len() == 0 || words[0].len() == 0 {
            return true;
        }
        let (M, N) = (words.len(), words[0].len());

        if M != N {
            return false;
        }

        let get = |x: usize, y: usize| {
            if x >= M {
                ' '
            } else if words[x].len() > M {
                '!'
            } else if y >= words[x].len() {
                ' '
            } else {
                words[x].chars().nth(y).unwrap()
            }
        };

        for i in 0..M {
            for j in i + 1..N {
                if get(i, j) != get(j, i) {
                    return false;
                }
            }
        }
        true
    }
}

pub fn run() {
    let v = vec![
        String::from("abc"),
        String::from("bde"),
        String::from("cefg"),
    ];
    println!("{}", Solution::valid_word_square(v));
}

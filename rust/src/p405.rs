struct Solution;

impl Solution {
    pub fn to_hex(num: i32) -> String {
        let mut n = num as u32;

        let mut p = 1;
        while p < 0x0fffffff && p * 16 <= n {
            p *= 16;
        }

        let mut r = String::from("");
        while p > 0 {
            match (n / p) as u8 {
                x @ 0...9 => r.push((x + '0' as u8) as char),
                x @ 10...15 => r.push((x + 'a' as u8 - 10) as char),
                _ => break
            }
            n %= p;
            p /= 16;
        }
        r
    }
}

pub fn run() {
    println!("{}", Solution::to_hex(26));
    println!("{}", Solution::to_hex(-1));
}

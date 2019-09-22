struct Solution;

impl Solution {
    pub fn arrange_coins(mut n: i32) -> i32 {
        let mut i = 1;
        let mut r = 0;

        while n - i >= 0 {
            n -= i;
            r += 1;
            i += 1;
        }

        r
    }
}

pub fn run() {
    println!("{}", Solution::arrange_coins(10));
}

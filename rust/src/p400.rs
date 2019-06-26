struct Solution;

fn pos(mut number: i32, p: i32) -> i32 {
    let mut x = 1;
    while x <= number {
        x *= 10;
    }

    println!("x={}, p={}, number={}", x, p, number);

    for i in 0..p {
        number -= (number / x) * x;
        x /= 10;
    }

    number / x
}

impl Solution {
    pub fn find_nth_digit(mut n: i32) -> i32 {
        let (mut len, mut next) = (1, 10);
        let mut i = 1;

        loop {
            if n <= len {
                return pos(i, n);
            }

            i += 1;
            n -= len;
            if i == next {
                next *= 10;
                len += 1;
            }
        }
    }
}

pub fn run() {
    let r = Solution::find_nth_digit(11);
    println!("{}", r);
}
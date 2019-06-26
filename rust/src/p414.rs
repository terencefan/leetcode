struct Solution;

impl Solution {
    pub fn third_max(nums: Vec<i32>) -> i32 {
        let (mut a, mut b, mut c) = (std::i64::MIN, std::i64::MIN, std::i64::MIN);

        for n32 in nums {
            let num = n32 as i64;
            if num > a {
                c = b;
                b = a;
                a = num;
            } else if num < a && num > b {
                c = b;
                b = num;
            } else if num < b && num > c {
                c = num;
            }
        }

        if c > std::i64::MIN {
            c as i32
        } else {
            a as i32
        }
    }
}

pub fn run() {
    let v = vec![5];
    println!("{}", Solution::third_max(v));
}
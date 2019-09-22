struct Solution;

fn quicksort(mut nums: Vec<i32>, left: usize, right: usize) {
}

impl Solution {
    pub fn min_moves(mut nums: Vec<i32>) -> i32 {
        let mut r = 0;

        quicksort(&nums, 0, nums.len() - 1);

        r
    }
}

pub fn run() {
    let v = vec![-2147483648,-1];
    let r = Solution::min_moves(v);
    println!("{}", r);
}

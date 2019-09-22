struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn number_of_boomerangs(points: Vec<Vec<i32>>) -> i32 {
        let n = points.len();

        let mut r = 0;
        for i in 0..n {
            let (x1, y1) = (points[i][0], points[i][1]);
            let mut m: HashMap<i32, i32> = HashMap::new();
            for j in 0..n {
                if i == j {
                    continue
                }
                let (x2, y2) = (points[j][0], points[j][1]);
                let (dx, dy) = (x2 - x1, y2 - y1);
                let distance = dx * dx + dy * dy;

                let v = m.entry(distance).or_insert(0);
                *v += 1
            }

            for v in m.values() {
                r += v * (v - 1);
            }
        }
        r
    }
}

pub fn run() {
    let v = vec![
        vec![0,0],
        vec![1,0],
        vec![2,0],
    ];
    let r = Solution::number_of_boomerangs(v);
    println!("{}", r);
}

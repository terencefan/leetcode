pub fn run() {
    // if condition can have return value
    let a = if 2 > 1 { 3 } else { 5 };

    // closed, or 0..10 means left-closed, right-open
    for number in 0..=10 {
        match number {
            // match is like switch but more powerful
            0 => println!("is zero"),
            m @ 1...3 => println!("between 1 and 3, {}", m),
            5 | 7 | 9 => println!("is 5 / 7 / 9"),
            n @ 8 => println!("equals to {}", n),
            _ => println!("others"),
        }
    }

    let mut v = vec![1, 2, 3];
    loop {
        match v.pop() {
            Some(x) => println!("{}", x),
            None => break,
        }
    }

    let mut v2 = vec![4, 5, 6];
    while let Some(x) = v2.pop() {
        println!("{}", x);
    }
}
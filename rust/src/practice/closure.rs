fn math<F: Fn() -> i32>(op: F) -> i32 {
    op()
}

fn build_times(i: i32) -> impl Fn(i32) -> i32 {
    move |j| j * i
}

pub fn run() {
    let a = 2;
    let b = 3;
    assert_eq!(math(|| a + b), 5);
    assert_eq!(math(|| a * b), 6);

    let times = build_times(3);
    println!("{}", times(5));
}
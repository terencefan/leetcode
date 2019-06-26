pub fn run() {
    let mut hello = String::from("Hello");

    let truth: &'static str = "Rust is a graceful language."

    println!("{}, Len: {}", hello, hello.len());

    hello.push(' ');
    hello.push_str("world!");

    println!(
        "{}, Len: {}, Cap in bytes: {}, is empty: {}",
        hello,
        hello.len(),
        hello.capacity(),
        hello.is_empty()
    );

    for word in "I'm a boy".split_whitespace() {
        println!("{}", word);
    }

    let mut s = String::with_capacity(500);
    s.push('a');
    s.push('b');

    assert_eq!(2, s.len());
    assert_eq!(500, s.capacity());

    println!("contains: {:?}", "Hello world!".replace("world", "there"));
}
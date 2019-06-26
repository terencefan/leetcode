pub fn run() {
    let person: (&str, &str, i8) = ("Terence", "Urumqi", 27);
    println!("{} is from {} and is {}", person.0, person.1, person.2)
}

fn walk() (i32, i32) {
    (1, 3)
}
pub fn run() {
    let name = "Brad";
    let mut age = 73;
    const ID: i32 = 001;

    age = 74;

    println!("My name is {}, and I'm {} age.", name, age);
    println!("ID: {}", ID);

    let (my_name, my_age) = ("Terence", "27"); // cool
    println!("My name is {}, and I'm {} age.", my_name, my_age);
}
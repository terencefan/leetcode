pub fn run() {
    let x = 1;
    let y = 2.5;
    let z: i64 = 199999;

    println!("Max i32: {}", std::i32::MAX);
    println!("Min i32: {}", std::i32::MIN);
    println!("Max u32: {}", std::u32::MAX);
    println!("Min u32: {}", std::u32::MIN);
    println!("Max f32: {}", std::f32::MAX);
    println!("Min f32: {}", std::f32::MIN);
    println!("Nan f32: {}", std::f32::NAN);
    println!("Inf f32: {}", std::f32::INFINITY);
    println!("-Inf f32: {}", std::f32::NEG_INFINITY);

    let a1: char = '😂'; // char is unicode

    let is_active = true;
    println!("{:?}", (x, y, z, x < 20, a1));
}
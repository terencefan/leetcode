use std::mem;

pub fn run() {
    let mut numbers: [i32; 5] = [1, 2, 3, 4, 5];
    numbers[2] = 7;

    let mut strs: [String; 20] = Default::default();
    for i in 0..20 {
        strs[i] = i.to_string();
    }

    println!("{:?}", numbers);
    println!("{:?}", strs);

    println!(
        "{}, Len: {}, Occupies: {}",
        numbers[0],
        numbers.len(),
        mem::size_of_val(&numbers)
    );

    let x = 1;

    // Get Slice
    let slice: &[i32] = &numbers[x..2];
    println!("Slice: {:?}", slice)
}
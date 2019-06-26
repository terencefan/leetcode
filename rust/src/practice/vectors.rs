use std::mem;

pub fn run() {
    let mut numbers: Vec<i32> = vec![1, 2, 3, 4, 5];

    numbers[2] = 7;

    numbers.push(555);
    numbers.pop();

    println!("{:?}", numbers);

    println!(
        "{}, Len: {}, Occupies: {}",
        numbers[0],
        numbers.len(),
        mem::size_of_val(&numbers)
    );

    let x = 1;

    // Get Slice
    let slice: &[i32] = &numbers[x..2];
    println!("Slice: {:?}", slice);

    for x in numbers.iter_mut() {
        *x *= 2
    }
    println!("Numbers: {:?}", numbers);
}
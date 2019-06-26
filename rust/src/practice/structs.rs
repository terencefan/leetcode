// Name-Field Struct
struct Color {
    red: u8,
    green: u8,
    blue: u8,
}

// Tuple-Like Struct
struct Color2(u8, u8, u8);

// Unit-Like Struct
// will be optimized to the same object in RELEASE mode.
struct Empty;

struct Person {
    first_name: String,
    last_name: String,
}

impl Person {
    fn new(first: &str, last: &str) -> Person {
        Person {
            first_name: first.to_string(),
            last_name: last.to_string(),
        }
    }

    fn full(&self) -> String {
        format!("{} {}", self.first_name, self.last_name)
    }

    fn set_last(&mut self, last: &str) {
        self.last_name = last.to_string()
    }

    fn to_tuple(self) -> (String, String) {
        (self.first_name, self.last_name)
    }
}

pub fn run() {
    let mut c = Color {
        red: 255,
        green: 0,
        blue: 0,
    };
    println!("Color: {} {} {}", c.red, c.green, c.blue);

    let mut c2 = Color2(255, 0, 0);
    println!("Color: {} {} {}", c2.0, c2.1, c2.2);

    let mut p = Person::new("Terence", "Fan");
    p.set_last("Tan");
    println!("{:?}", p.to_tuple());
}
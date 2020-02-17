use std::fs;
use std::io::Bytes;

fn read_file(filepath: &str) -> Result<dyn Bytes, dyn std::error::Error> {
    let f = fs::File::open(filepath);
    let data: Bytes = match f {
        Ok(file) => file.bytes(),
        Err(error) => return Err("error no hay error"),
    };

    Ok(data)
}

fn main() {}

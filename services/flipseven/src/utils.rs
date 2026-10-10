pub fn simple_function() {
    println!("This is a simple function.");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_simple_function() {
        simple_function();
    }
}

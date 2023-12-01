fn solve_part_1(_input: &str) -> String {
    let result = "Hello, World";
    return String::from(result);
}

fn solve_part_2(_input: &str) -> String {
    let result = "Goodbye, World";
    return String::from(result);
}


fn main() {
    println!("Part 1 Answer -> {}", solve_part_1("Actual Case"));
    println!("Part 2 Answer -> {}", solve_part_2("Actual Case"));
}




#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solve() {
        assert_eq!(solve_part_1("Test Case 1"), String::from("Hello, World"));
        assert_eq!(solve_part_2("Test Case 2"), String::from("Goodbye, World"));
    }
}
#![feature(yeet_expr)]

use std::fs::read_to_string;
//use std::str::FromStr;

fn solve_part_1(input: &str) -> Result<i32, String> {
    for _line in read_to_string(input).unwrap().lines() {
        
    }
    return Ok(0);
}

fn solve_part_2(input: &str) -> Result<i32, String> {
    for _line in read_to_string(input).unwrap().lines() {
        
    }
    return Ok(0);
}


fn main() {
    println!("Part 1 Answer -> {}", solve_part_1("input.txt").unwrap());
    println!("Part 2 Answer -> {}", solve_part_2("input.txt").unwrap());
}




#[cfg(test)]
mod tests {
    use super::*;

    const PART_1_EXAMPLE_SOLUTION : i32 = 0;
    const PART_2_EXAMPLE_SOLUTION : i32 = 0;

    const PART_1_ACTUAL_SOLUTION : i32 = 0;
    const PART_2_ACTUAL_SOLUTION : i32 = 0;

    // Test the solution for part 1 works for the example input
    #[test]
    fn test_solve_part_1_example() {
        assert_eq!(solve_part_1("test_input.txt"), Ok(PART_1_EXAMPLE_SOLUTION));
    }

    // Test the solution for part 2 works for the example input
    #[test]
    fn test_solve_part_2_example() {
        assert_eq!(solve_part_2("test_input.txt"), Ok(PART_2_EXAMPLE_SOLUTION));
    }

    // Test the solution for part 1 works for the actual input
    #[test]
    fn test_solve_part_1_actual() {
        assert_eq!(solve_part_1("input.txt"), Ok(PART_1_ACTUAL_SOLUTION));
    }

    // Test the solution for part 2 works for the actual input
    #[test]
    fn test_solve_part_2_actual() {
        assert_eq!(solve_part_2("input.txt"), Ok(PART_2_ACTUAL_SOLUTION));
    }
}
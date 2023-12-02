#![feature(yeet_expr)]

use std::fs::read_to_string;
use std::str::FromStr;
use std::cmp;

const RED_CUBES: i32 = 12;
const GREEN_CUBES: i32 = 13;
const BLUE_CUBES: i32 = 14;

fn solve_part_1(input: &str) -> Result<i32, String> {
    let mut result : i32 = 0;

    for line in read_to_string(input).unwrap().lines() {
        let mut max_reds : i32 = 0;
        let mut max_greens: i32 = 0;
        let mut max_blues: i32 = 0;

        
        let game_id : Vec<&str> = line.split(":").next().unwrap().split(" ").collect();
        let game_num: i32 = FromStr::from_str(game_id.last().unwrap()).unwrap();

        let color_sets = line.split(":").last().unwrap();

        for set in color_sets.split(";") {
            for color in set.split(","){
                let color_val: i32 = FromStr::from_str(color.split_whitespace().next().unwrap()).unwrap();
                let color_name = color.split_whitespace().last().unwrap();

                match color_name {
                    "red" => max_reds = cmp::max(max_reds, color_val),
                    "green" => max_greens = cmp::max(max_greens, color_val),
                    "blue" => max_blues = cmp::max(max_blues, color_val),
                    _=> do yeet "There was an issue"
                }
            }
        }

        if max_reds <= RED_CUBES && max_greens <= GREEN_CUBES && max_blues <= BLUE_CUBES {
            result += game_num;
        }

    }

    return Ok(result);
}

fn solve_part_2(input: &str) -> Result<i32, String>{
    let mut result : i32 = 0;

    for line in read_to_string(input).unwrap().lines() {
        let mut max_reds : i32 = 0;
        let mut max_greens: i32 = 0;
        let mut max_blues: i32 = 0;

        let color_sets = line.split(":").last().unwrap();

        for set in color_sets.split(";") {
            for color in set.split(","){
                let color_val: i32 = FromStr::from_str(color.split_whitespace().next().unwrap()).unwrap();
                let color_name = color.split_whitespace().last().unwrap();

                match color_name {
                    "red" => max_reds = cmp::max(max_reds, color_val),
                    "green" => max_greens = cmp::max(max_greens, color_val),
                    "blue" => max_blues = cmp::max(max_blues, color_val),
                    _=> do yeet "There was an issue"
                }
            }
        }

        let game_power: i32 = max_reds * max_greens * max_blues;
        result += game_power;

    }

    return Ok(result);
}


fn main() {
    println!("Part 1 Answer -> {}", solve_part_1("input.txt").unwrap());
    println!("Part 2 Answer -> {}", solve_part_2("input.txt").unwrap());
}




#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solve_part_1() {
        assert_eq!(solve_part_1("test_input.txt"), Ok(8));
    }

    #[test]
    fn test_solve_part_2() {
        assert_eq!(solve_part_2("test_input.txt"), Ok(2286));
    }
}
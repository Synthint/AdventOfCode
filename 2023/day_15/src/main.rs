#![feature(yeet_expr)]

use std::fmt;
use std::fs::read_to_string;
//use std::str::FromStr;

#[derive(Debug)]
struct Lens {
    label : String,
    power : i32,
}

impl fmt::Display for Lens {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{} -> {}", self.label,self.power)
    }
}

fn solve_part_1(input: &str) -> Result<i32, String> {
    let mut full_input : String = "".to_string();
    for _line in read_to_string(input).unwrap().lines() {
        full_input.push_str(_line)
    }

    let lenses = full_input.split(",");

    let mut total_sum: i32 = 0;
    for lens in lenses {
        total_sum += hash_string(lens.to_string());
    }
    return Ok(total_sum);
}

fn str_to_lens(input : &str) -> Result<Lens, String>{
    let parts = input.split("=").collect::<Vec<&str>>();
    let ret_lens = Lens{
        label: parts.get(0)
                    .expect("Invalid or Missing Lens Label")
                    .to_string(),
        power: parts.get(1)
                    .expect("Invalid or Missing Focal Length")
                    .parse()
                    .expect("Focal Length Not Int")
    };

    return Ok(ret_lens)
}


fn hash_string(input: String) -> i32 {
    let mut current_value: i32 = 0;
    for curr in input.chars() {
        let ascii = curr as u8 as i32;
        current_value += ascii;
        
        current_value *= 17;
        current_value = current_value % 256;
    }
    return current_value
}


fn solve_part_2(input: &str) -> Result<i32, String> {
    let mut full_input : String = "".to_string();
    for _line in read_to_string(input).unwrap().lines() {
        full_input.push_str(_line)
    }

    let mut boxes :  Vec<Vec<Lens>> = Vec::with_capacity(256);
    for _ in 0..256 {
        boxes.push(Vec::<Lens>::new())
    }
    //let mut boxes: Vec<Vec<Lens>> = vec![vec![]];
    let lenses = full_input.split(",");
    for lens in lenses {

        if lens.contains("=") {
            let curr : Lens = str_to_lens(lens).unwrap();
            let box_num: usize = hash_string(curr.label.clone()) as usize;
            let mut found = false;
            for curr_lens in boxes[box_num].iter_mut() {
                if curr_lens.label == curr.label {
                    found = true;
                    let new_power = str_to_lens(lens).unwrap().power;
                    curr_lens.power = new_power;
                    break
                }
            }
            if !found {
                boxes[box_num].push(str_to_lens(lens).unwrap());
            }
            
        }else{
            let mut rem_lens = lens.to_string();
            rem_lens.pop();
            let box_num = hash_string(rem_lens.clone()) as usize;
            let mut rem_ind: usize = 0;
            let mut do_remove = false;
            for (i,lens) in boxes[box_num].iter().enumerate() {
                if lens.label == rem_lens {
                    rem_ind = i;
                    do_remove = true;
                    break;
                }
            }
            if do_remove {
                boxes[box_num].remove(rem_ind);
            }


        }
    }
    
    let mut total_focusing_power = 0;
    for (box_num,lens_box) in boxes.iter().enumerate() {
        for (lens_num, lens) in lens_box.iter().enumerate() {
            let current_power = (1 + box_num as  i32) * (1+lens_num as i32) * (lens.power);
            total_focusing_power += current_power
        }
    }

    return Ok(total_focusing_power)
}


fn main() {
    println!("Part 1 Answer -> {}", solve_part_1("input.txt").unwrap());
    println!("Part 2 Answer -> {}", solve_part_2("input.txt").unwrap());
}




#[cfg(test)]
mod tests {
    use super::*;

    const PART_1_EXAMPLE_SOLUTION : i32 = 1320;
    const PART_2_EXAMPLE_SOLUTION : i32 = 0;
    
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
}
#![feature(yeet_expr)]

// use std::fs::read_to_string;
// use std::str::FromStr;




fn solve_part_1(times: Vec<i32>, dists: Vec<i32>) -> Result<i64, String> {
    let mut better_dists_mult: i64 = 1;
    for (ind, time) in times.iter().enumerate() {
        let mut better_count: i32 = 0;
        for hold in 0..time+1 {
            if (time-hold)*hold > dists[ind] {
                better_count += 1;
            }
        }
        better_dists_mult *= better_count as i64;
    }
    
    return Ok(better_dists_mult);
}

fn solve_part_2(time: i64, dist: i64) -> Result<i64, String> {
    let mut better_count: i64 = 0;
    for hold in 0..time+1 {
        if (time-hold)*hold > dist {
            better_count += 1;
        }
    }
    return Ok(better_count);
}




fn main() {
    let part_1_times = vec![40,82,91,66];
    let part_1_dists = vec![277,1338,1349,1063];
    println!("Part 1 Answer -> {}", solve_part_1(part_1_times,part_1_dists).unwrap());

    let part_2_time: i64 = 40829166;
    let part_2_dist: i64 = 277133813491063;
    println!("Part 2 Answer -> {}", solve_part_2(part_2_time,part_2_dist).unwrap());
}




#[cfg(test)]
mod tests {
    use super::*;

    const PART_1_EXAMPLE_SOLUTION : i64 = 288;
    const PART_2_EXAMPLE_SOLUTION : i64 = 71503;

    // Test the solution for part 1 works for the example input
    #[test]
    fn test_solve_part_1_example() {
        let test_1_times = vec![7, 15, 30];
        let test_1_dists = vec![9, 40, 200];
        assert_eq!(solve_part_1(test_1_times,test_1_dists), Ok(PART_1_EXAMPLE_SOLUTION));
    }

    // Test the solution for part 2 works for the example input
    #[test]
    fn test_solve_part_2_example() {
        let test_2_time: i64 = 71530;
        let test_2_dist: i64 = 940200;
        assert_eq!(solve_part_2(test_2_time,test_2_dist), Ok(PART_2_EXAMPLE_SOLUTION));
    }

}
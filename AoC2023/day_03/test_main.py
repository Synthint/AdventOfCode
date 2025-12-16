from main import solve_part_1, solve_part_2
import pytest

# Test the solution for part 1 works for the example input
@pytest.mark.parametrize("input, actual",
                          [("test_input.txt", 4361)])
def test_solve_part_1_example(input,actual):
    assert solve_part_1(input) == actual

# Test the solution for part 2 works for the example input
@pytest.mark.parametrize("input, actual",
                          [("test_input.txt", 467835)])
def test_solve_part_2_example(input,actual):
    assert solve_part_2(input) == actual

# Test the solution for part 1 works for the actual input
@pytest.mark.parametrize("input, actual",
                          [("input.txt", 525181)])
def test_solve_part_1_actual(input,actual):
    assert solve_part_1(input) == actual

# Test the solution for part 2 works for the actual input
@pytest.mark.parametrize("input, actual",
                          [("input.txt", 84289137)])
def test_solve_part_2_actual(input,actual):
    assert solve_part_2(input) == actual
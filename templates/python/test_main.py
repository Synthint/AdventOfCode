from main import solve_part_1, solve_part_2
import pytest

PART_1_EXAMPLE_SOLUTION = None
PART_2_EXAMPLE_SOLUTION = None

PART_1_ACTUAL_SOLUTION = None
PART_2_ACTUAL_SOLUTION = None


  # Test the solution for part 1 works for the example input
@pytest.mark.parametrize("input, actual",
                          [("test_input.txt", PART_1_EXAMPLE_SOLUTION)])
def test_solve_part_1_example(input,actual):
    assert solve_part_1(input) == actual

# Test the solution for part 2 works for the example input
@pytest.mark.parametrize("input, actual",
                          [("test_input.txt", PART_2_EXAMPLE_SOLUTION)])
def test_solve_part_2_example(input,actual):
    assert solve_part_2(input) == actual

# Test the solution for part 1 works for the actual input
@pytest.mark.parametrize("input, actual",
                          [("input.txt", PART_1_ACTUAL_SOLUTION)])
def test_solve_part_1_actual(input,actual):
    assert solve_part_1(input) == actual

# Test the solution for part 2 works for the actual input
@pytest.mark.parametrize("input, actual",
                          [("input.txt", PART_2_ACTUAL_SOLUTION)])
def test_solve_part_2_actual(input,actual):
    assert solve_part_2(input) == actual
from main import solve_part_1, solve_part_2
import pytest


@pytest.mark.parametrize("input_1, input_2, expected_1, expected_2",
                          [("test_input.txt", "test_input.txt",
                            4361, 467835)])
def test_problem(input_1, input_2, expected_1, expected_2):
    assert solve_part_1(input_1) == expected_1
    assert solve_part_2(input_2) == expected_2
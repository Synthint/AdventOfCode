import {solve_part_1,solve_part_2} from "./main"

const PART_1_EXAMPLE_SOLUTION = 142;
const PART_2_EXAMPLE_SOLUTION = 281;

describe("Solve Tests", () =>{
    test("Check literal value", () => {
        expect( solve_part_1("test_input_1.txt") ).toBe(PART_1_EXAMPLE_SOLUTION);
        expect( solve_part_2("test_input_2.txt") ).toBe(PART_2_EXAMPLE_SOLUTION);
    })
});


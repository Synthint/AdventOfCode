import {solve_part_1,solve_part_2} from "./main"

const PART_1_EXAMPLE_SOLUTION = 24000;
const PART_2_EXAMPLE_SOLUTION = 45000;

describe("Solve Tests", () =>{
    test("Check literal value", () => {
        expect( solve_part_1("test_input.txt") ).toBe(PART_1_EXAMPLE_SOLUTION);
        expect( solve_part_2("test_input.txt") ).toBe(PART_2_EXAMPLE_SOLUTION);
    })
});


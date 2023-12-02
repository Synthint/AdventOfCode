import {solve_part_1,solve_part_2} from "./main"


describe("Solve Tests", () =>{
    test("Check literal value", () => {
        expect( solve_part_1() ).toBe("Hello, World");
        expect( solve_part_2() ).toBe("Hello, World");
    })
});


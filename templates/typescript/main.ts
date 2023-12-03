import * as fs from 'fs';

export function solve_part_1(file : string) : number{
    const lines_arr = fs.readFileSync(file, 'utf-8').split('\n');
    for (let i = 0; i < lines_arr.length; i++) {
        const line = lines_arr[i];
        
    }
    return -1
}


export function solve_part_2(file : string) : number{
    const lines_arr = fs.readFileSync(file, 'utf-8').split('\n');
    for (let i = 0; i < lines_arr.length; i++) {
        const line = lines_arr[i];
        
    }
    return -1
}




console.log("Part 1 Answer ->",solve_part_1("input.txt"))
console.log("Part 2 Answer ->",solve_part_2("input.txt"))



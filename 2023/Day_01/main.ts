import * as fs from 'fs';

export function solve_part_1(file : string) : number{
    const lines_arr = fs.readFileSync(file, 'utf-8').split('\n');
    let sum: number = 0
    const regexp: RegExp = /^[a-zA-Z]*([0-9]).*([0-9])[a-zA-Z]*$|^[a-zA-Z]*([0-9])[a-zA-Z]*$/g
    for (let i = 0; i < lines_arr.length; i++) {
        const line = lines_arr[i];
        const matches = Array.from(line.matchAll(regexp))
        if (matches[0][3] === undefined){
            sum += Number(matches[0][1])*10 + Number(matches[0][2])
        }else{
            sum += Number(matches[0][3])*10 + Number(matches[0][3])
        }
    }
    return sum
}

function str_to_int(input: string): number{
    let convert = { "zero":0, "one":1, "two":2, "three":3,
                "four":4, "five":5, "six":6, "seven":7,
                "eight":8, "nine":9, "0":0, "1":1, "2":2,
                "3":3, "4":4, "5":5, "6":6, "7":7, "8":8, "9":9, }
    return convert[input]
}

export function solve_part_2(file : string) : number{
    const lines_arr = fs.readFileSync(file, 'utf-8').split('\n');
    let sum: number = 0
    const regexp: RegExp = /^[a-zA-Z]*?([0-9]|one|two|three|four|five|six|seven|eight|nine|zero).*([0-9]|one|two|three|four|five|six|seven|eight|nine|zero)[a-zA-Z]*?$|^[a-zA-Z]*([0-9]|one|two|three|four|five|six|seven|eight|nine|zero)[a-zA-Z]*?$/g
    for (let i = 0; i < lines_arr.length; i++) {
        const line = lines_arr[i];
        const matches = Array.from(line.matchAll(regexp))
        if (matches[0][3] === undefined){
            sum += str_to_int(matches[0][1])*10 + str_to_int(matches[0][2])
        }else{
            sum += str_to_int(matches[0][3])*10 + str_to_int(matches[0][3])
        }
    }
    return sum
}




console.log("Part 1 Answer ->",solve_part_1("input.txt"))
console.log("Part 2 Answer ->",solve_part_2("input.txt"))



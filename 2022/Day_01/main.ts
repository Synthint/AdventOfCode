import * as fs from 'fs';

export function solve_part_1(file : string) : number{
    const lines_arr = fs.readFileSync(file, 'utf-8').split('\n');
    var max_cals = 0;
    var curr_cals = 0;
    for (let i = 0; i < lines_arr.length; i++) {
        const line = lines_arr[i];
        if (line == ""){
            if(curr_cals>max_cals){
                max_cals = curr_cals;
            }
            curr_cals = 0;
        }else{
            curr_cals += parseInt(line);
        }
        
    }
    return max_cals
}

export function solve_part_2(file : string) : number{
    const lines_arr = fs.readFileSync(file, 'utf-8').split('\n');
    var max_cals = [0,0,0];
    var curr_cals = 0;
    for (let i = 0; i < lines_arr.length; i++) {
        const line = lines_arr[i];
        if (line == ""){
            for (let j = 0; j < max_cals.length; j++) {
                if(curr_cals>max_cals[j]){
                    max_cals[j] = curr_cals;
                    //sort max_cals smallest to largest
                    max_cals.sort((a,b) => a-b);
                    break;
                }
            }
            curr_cals = 0;
        }else{
            curr_cals += parseInt(line);
        }
        
    }
    return max_cals[0]+max_cals[1]+max_cals[2];
}




console.log("Part 1 Answer ->",solve_part_1("input.txt"))
console.log("Part 2 Answer ->",solve_part_2("input.txt"))



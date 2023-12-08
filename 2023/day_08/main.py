import re
from multiprocessing import Pool
from math import gcd


def solve_part_1(file):
    lines = []
    with open(file, "r") as input:
        for line in input:
            lines.append(line)
    
    instruct_str = lines[0].strip()
    map_dict = {}

    for line in range(2,len(lines)):
        curr = lines[line]
        key = curr.split("=")[0].strip()
        coords = curr.split("=")[1].strip()
        left = coords.split(",")[0].strip("(").strip()
        right = coords.split(",")[1].strip(")").strip()
        map_dict.update({key : (left,right)})

    steps = 0
    next_step = "AAA"
    while next_step != "ZZZ":
        if instruct_str[steps % len(instruct_str)] == "L":
            next_step = map_dict[next_step][0]
            steps += 1
        elif instruct_str[steps % len(instruct_str)] == "R":
            next_step = map_dict[next_step][1]
            steps += 1

    return steps



def find_end_Z_steps(start, map_dict, instruct_str):
    steps = 0
    next_step = start
    while next_step[2] != "Z":
        if instruct_str[steps % len(instruct_str)] == "L":
            next_step = map_dict[next_step][0]
            steps += 1
        elif instruct_str[steps % len(instruct_str)] == "R":
            next_step = map_dict[next_step][1]
            steps += 1
    return steps


def solve_part_2(file):
    lines = []
    with open(file, "r") as input:
        for line in input:
            lines.append(line)
    
    instruct_str = lines[0].strip()
    all_paths = []
    map_dict = {}

    for line in range(2,len(lines)):
        curr = lines[line]
        key = curr.split("=")[0].strip()
        if key[2] == "A":
            all_paths.append(key)
        coords = curr.split("=")[1].strip()
        left = coords.split(",")[0].strip("(").strip()
        right = coords.split(",")[1].strip(")").strip()
        map_dict.update({key : (left,right)})

    pool_inputs = []
    for item in all_paths:
        pool_inputs.append((item,map_dict,instruct_str))
    

    
    pool = Pool(len(pool_inputs))
    results = pool.starmap(find_end_Z_steps,pool_inputs)
        
    lcm = 1
    for i in results:
        lcm = lcm*i//gcd(lcm, i)
    return lcm

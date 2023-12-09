import re

def recur_predictions(nums):
    sub = find_sub_history(nums)
    if all(i == 0 for i in sub) :
       nums.append(nums[-1])
    else:
        sub = recur_predictions(sub)
        nums.append(nums[-1]+sub[-1])
    return nums
    

def find_sub_history(nums):
    sub = []
    for x in range(len(nums)-1):
        sub.append(nums[x+1] - nums[x])
    return sub

def solve_part_1(file):
    lines = []
    with open(file, "r") as input:
        for line in input:
            lines.append(line)
    new_tot = 0
    for line in lines:
        nums = [int(item.strip()) for item in line.split(" ")]
        new = recur_predictions(nums)
        new_tot += new[-1]
    return new_tot
    

def recur_more_history(nums):
    sub = find_sub_history(nums)
    if all(i == 0 for i in sub) :
       nums.insert(0, nums[0])
    else:
        sub = recur_more_history(sub)
        nums.insert(0 ,nums[0] - sub[0] )
    return nums


def solve_part_2(file):
    lines = []
    with open(file, "r") as input:
        for line in input:
            lines.append(line)
    new_tot = 0
    for line in lines:
        nums = [int(item.strip()) for item in line.split(" ")]
        new = recur_more_history(nums)
        new_tot += new[0]
    return new_tot

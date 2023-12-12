from functools import cache
# n-k+1 -> k consecutive elements from n
#????????????? 1,5,1,1
# n = 13
# 13 - 5 = 8 + 1 = 9

@cache 
def recur_count(record: str , nums):
    sum = 0
    if ( len(nums) == 0 or len(record) == 0 or
        ("#" not in record and "?" not in record) ):
        return sum
    
    
    if record[0] == "?":
        sum += recur_count(record.replace("?","#",1), nums)
        sum += recur_count(record.replace("?",".",1), nums)
    elif record[0] == ".":
        sum += recur_count(record[1:],nums)
    elif record[0] == "#":
        count = 0
        while count < len(record) and record[count] == "#":
            count += 1
        if count == nums[0]:
            nums = nums[1:]
            record = record[count:]

            if len(nums) == 0 and "#" not in record:
                sum += 1
            elif len(record) > 0 and record[0] == ".":
                sum += recur_count(record, nums)
            elif len(record) > 0 and record[0] == "?":
                sum += recur_count(record.replace("?",".",1), nums)
        elif count < nums[0] and count < len(record) and record[count] == "?":
            sum += recur_count(record.replace("?","#",1), nums)
    return sum

def solve_part_1(file):
    sum = 0
    with open(file, "r") as input:
        for line in input:
            record = line.split(" ")[0]
            nums = tuple([int(i) for i in line.split(" ")[1].split(",")])
            sum += recur_count(record , nums)
    return sum
        

def solve_part_2(file):
    sum = 0
    with open(file, "r") as input:
        for line in input:
            line = line.strip()
            record = ((line.split(" ")[0]+"?") * 5)[:-1]
            nums = tuple([int(i) for i in ((line.split(" ")[1]+",") * 5).split(",")[:-1]])
            sum += recur_count(record , nums)
    return sum


def solve_general(file, unfold_count):
    sum = 0
    with open(file, "r") as input:
        for line in input:
            line = line.strip()
            record = ((line.split(" ")[0]+"?") * unfold_count)[:-1]
            nums = tuple([int(i) for i in ((line.split(" ")[1]+",") * unfold_count).split(",")[:-1]])
            sum += recur_count(record , nums)
    return sum
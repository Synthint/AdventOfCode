from functools import cache

def solve_part_1(file):
    lines = []
    with open(file, "r") as input:
        lines = [item.strip() for item in input.readlines()]

    new_lines = roll_rocks(tuple(lines),0)
    sum = 0
    for ind, item in enumerate(new_lines):
        sum += (item.count("O") * (len(new_lines)-ind))
        # print(f"{item} -> {item.count('O')},{len(new_lines)-ind}")
    roll_rocks.cache_clear()
    return sum


@cache
def roll_rocks(platform,dir):
    lines_rot = rotate_left(platform)
    for n in range(dir):
        lines_rot = rotate_left(lines_rot)
    new_lines_rot = []
    for item in lines_rot:
        rock_li = list(item)
        new_str = ""
        curr_str = ""
        while len(rock_li) != 0:
            next = rock_li.pop(0)
            if next == "#":
                curr_str = ""+("O" * curr_str.count("O")) + ("." * curr_str.count("."))
                new_str += curr_str + "#"
                curr_str = ""
            else:
                curr_str += next
        curr_str = ""+("O" * curr_str.count("O")) + ("." * curr_str.count("."))
        new_str += curr_str
        new_lines_rot.append(new_str)

    new_lines = rotate_right(new_lines_rot)
    for n in range(dir):
        new_lines = rotate_right(new_lines)
    return tuple(new_lines)


def rotate_right(platform):
    rot_li = list(zip(*platform[::-1]))
    ret = []
    for row in rot_li:
        new_str = ""
        for col in row:
            new_str += col
        ret.append(new_str)
    return ret

def rotate_left(platform):
    rot_li = reversed(list(zip(*platform[::-1])))
    ret = []
    for row in rot_li:
        new_str = ""
        for col in row:
            new_str += col
        ret.append(new_str[::-1])
    return ret

def solve_part_2(file):
    lines = []
    with open(file, "r") as input:
        lines = [item.strip() for item in input.readlines()]

    new_lines = roll_rocks(roll_rocks(roll_rocks(roll_rocks(tuple(lines),0),3),2),1)
    for n in range(1,1000000000):
        if n % 10000 == 0:
            print(n)
        new_lines = roll_rocks(roll_rocks(roll_rocks(roll_rocks(tuple(new_lines),0),3),2),1)
    
    sum = 0
    for ind, item in enumerate(new_lines):
        sum += (item.count("O") * (len(new_lines)-ind))
        # print(f"{item} -> {item.count('O')},{len(new_lines)-ind}")
    print(roll_rocks.cache_info())
    roll_rocks.cache_clear()
    return sum
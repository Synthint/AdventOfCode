
import math


def solve_part_1(file):
    lines = []
    with open(file, "r") as input:
        for line in input:
            lines.append(line.strip())
    
    lines_expanded = []
    for line in lines:
        lines_expanded.append(line)
        if "#" not in line:
            lines_expanded.append(line)
    lines_rotated = [''.join(s) for s in zip(*lines_expanded)]
    
    lines_expanded = []
    for line in lines_rotated:
        lines_expanded.append(line)
        if "#" not in line:
            lines_expanded.append(line)

    lines_rotated = [''.join(s) for s in zip(*lines_expanded)]
    
    galaxy_coords = []
    for x in range(len(lines_rotated)):
        for y in range(len(lines_rotated[x])):
            if lines_rotated[x][y] == "#":
                galaxy_coords.append((x,y))
    sum = 0
    for ind in range(len(galaxy_coords)):
        for check in range(ind,len(galaxy_coords)):
            sum += calc_dist(galaxy_coords[ind],galaxy_coords[check])
    return sum

def calc_dist(coord1,coord2):
    return abs(coord2[0] - coord1[0]) + abs(coord2[1] - coord1[1])


 
def solve_part_2(file, empty_mult):
    lines = []
    with open(file, "r") as input:
        for line in input:
            lines.append(line.strip())

    galaxy_coords = []
    x_adjust = 0
    y_adjust = 0
    for x in range(len(lines)):
        y_adjust = 0
        for y in range(len(lines[x])):
            has_galaxy = False
            for w in range(len(lines)):
                if lines[w][y] == "#":
                    has_galaxy = True
            if not has_galaxy:
                y_adjust += empty_mult - 1
            if "#" not in lines[x]:
                x_adjust += empty_mult - 1
                break
            elif lines[x][y] == "#":
                galaxy_coords.append((x+x_adjust,y+y_adjust))
    sum = 0
    for ind in range(len(galaxy_coords)):
        for check in range(ind,len(galaxy_coords)):
            sum += calc_dist(galaxy_coords[ind],galaxy_coords[check])
    return sum
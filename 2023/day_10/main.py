import math
import re

"""


    | is a vertical pipe connecting north and south.
    - is a horizontal pipe connecting east and west.
    L is a 90-degree bend connecting north and east.
    J is a 90-degree bend connecting north and west.
    7 is a 90-degree bend connecting south and west.
    F is a 90-degree bend connecting south and east.
    . is ground; there is no pipe in this tile.
    S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.

    
    S

"""


pipe_guide = {
    "|":((-1,0),(1,0)),
    "-":((0,-1),(0,1)),
    "L":((-1,0),(0,1)),
    "J":((-1,0),(0,-1)),
    "7":((1,0),(0,-1)),
    "F":((1,0),(0,1)),
    ".":((-100,-100),(-100,-100)),
    "S":((-100,-100),(-100,-100))
    }


class Pipe:
    connections = ((0,0),(0,0))
    position = (-1,-1)
    shape = "."
    

    def __init__(self, shape, pos_y, pos_x) -> None:
        self.shape = shape
        connections = pipe_guide[shape]

        self.connections = ((connections[0][0] + pos_y,connections[0][1] + pos_x),
                       (connections[1][0] + pos_y,connections[1][1] + pos_x))

        self.position = (pos_y,pos_x)

    def in_to_out(self, in_coords):
        if in_coords == self.connections[0]:
            return self.connections[1]
        elif in_coords == self.connections[1]:
            return self.connections[0]
        else:
            return -1
        
    def __repr__(self):
        return f"{self.connections[0]} {self.shape} {self.connections[1]}" 


def solve_part_1(file):
    lines = []
    pos_x = 0
    pos_y = 0
    pipes = []
    with open(file, "r") as input:
        for line in input:
            lines.append(line.strip())
    for y in range(len(lines)):
        for x in range(len(lines[y])):
            pipes.append(Pipe(
                lines[y][x],
                y,
                x
            ))

    for pipe in pipes:
        if pipe.shape == "S":
            pos_y = pipe.position[0]
            pos_x = pipe.position[1]

    pipe_loop = []
    for pipe in pipes:
        next_coords = pipe.in_to_out((pos_y,pos_x))
        if next_coords != -1:
            pipe_loop.append(pipe)
            break
    
    curr_coords = (pos_y,pos_x)
    while pipe_loop[-1].shape != "S":
        next_coords = pipe_loop[-1].in_to_out(curr_coords)
        for pipe in pipes:
            if pipe.position == next_coords:
                curr_coords = pipe_loop[-1].position
                pipe_loop.append(pipe)
    
    return int(len(pipe_loop)/2)








def solve_part_2(file):
    lines = []
    pos_x = 0
    pos_y = 0
    pipes = []
    with open(file, "r") as input:
        for line in input:
            lines.append(line.strip())
    for y in range(len(lines)):
        for x in range(len(lines[y])):
            pipes.append(Pipe(
                lines[y][x],
                y,
                x
            ))

    for pipe in pipes:
        if pipe.shape == "S":
            pos_y = pipe.position[0]
            pos_x = pipe.position[1]

    pipe_loop = []
    for pipe in pipes:
        next_coords = pipe.in_to_out((pos_y,pos_x))
        if next_coords != -1:
            pipe_loop.append(pipe)
            break
    
    curr_coords = (pos_y,pos_x)
    while pipe_loop[-1].shape != "S":
        next_coords = pipe_loop[-1].in_to_out(curr_coords)
        for pipe in pipes:
            if pipe.position == next_coords:
                curr_coords = pipe_loop[-1].position
                pipe_loop.append(pipe)


    pipe_coords = []
    for pipe in pipe_loop:
        pipe_coords.append(pipe.position)

    ground_coords = []
    for y in range(len(lines)):
        for x in range(len(lines[y])):
            if (y,x) not in pipe_coords:
                if ray_tracing(y,x,pipe_coords):
                    ground_coords.append((y,x))
    return len(ground_coords)


def ray_tracing(x,y,poly):
    n = len(poly)
    inside = False
    p2x = 0.0
    p2y = 0.0
    xints = 0.0
    p1x,p1y = poly[0]
    for i in range(n+1):
        p2x,p2y = poly[i % n]
        if y > min(p1y,p2y):
            if y <= max(p1y,p2y):
                if x <= max(p1x,p2x):
                    if p1y != p2y:
                        xints = (y-p1y)*(p2x-p1x)/(p2y-p1y)+p1x
                    if p1x == p2x or x <= xints:
                        inside = not inside
        p1x,p1y = p2x,p2y

    return inside
"""
    removal_occured = True
    while removal_occured:
        bad_ground = []
        for ground in ground_coords:
            for why in range(ground[0]-1,ground[0]+2):
                for ex in range(ground[1]-1,ground[1]+2):
                    if (why,ex) not in pipe_coords and (why,ex) not in ground_coords:
                        bad_ground.append(ground)

        bad_ground = set(bad_ground)  
        removal_occured = False
        for ground in bad_ground:
           removal_occured = True
           ground_coords.remove(ground)

    print()
    for y in range(len(lines)):
        out_str = ""
        for x in range(len(lines[y])):
            if (y,x) in pipe_coords:
                out_str += lines[y][x]
            elif (y,x) in ground_coords:
                out_str += "█"
            else:
                out_str += " "
        print(out_str)
    print()
    print(ground_coords)
    return len(ground_coords)
"""
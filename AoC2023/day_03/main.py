import re

def solve_part_1(file):
    nums = ["0","1","2","3","4","5","6","7","8","9"]
    ret_sum = 0
    with open(file, "r") as input:
        engine_map = [line for line in input]
        for x in range(len(engine_map)):
            engine_map[x] = [charac for charac in engine_map[x]]


        for x in range(len(engine_map)):
            y = 0
            while y < len(engine_map[x]):
                good = False
                neigborhood = []
                curr_num = 0
                while engine_map[x][y] in nums:
                    
                    curr_num = (curr_num*10) + int(engine_map[x][y])
                    if x > 0:
                        neigborhood.append(engine_map[x-1][y])
                        if y > 0:
                            neigborhood.append(engine_map[x-1][y-1])
                        if y < len(engine_map[x])-2:
                            neigborhood.append(engine_map[x-1][y+1])
                    if x < len(engine_map)-2:
                        neigborhood.append(engine_map[x+1][y])
                        if y > 0:
                            neigborhood.append(engine_map[x+1][y-1])
                        if y < len(engine_map[x])-2:
                            neigborhood.append(engine_map[x+1][y+1])
                    if y > 0:
                        neigborhood.append(engine_map[x][y-1])
                    if y < len(engine_map[x])-2:
                        neigborhood.append(engine_map[x][y+1])
                    y = y + 1
                
                for charac in neigborhood:
                    if charac not in nums and charac != ".":
                        good = True
                if good:
                    ret_sum = ret_sum + curr_num
                y=y+1
        return ret_sum
                

        

def solve_part_2(file):
    nums = ["0","1","2","3","4","5","6","7","8","9"]
    possible_gears = {'(-100, -100)':""}
    ret_sum = 0
    with open(file, "r") as input:
        engine_map = [line for line in input]
        for x in range(len(engine_map)):
            engine_map[x] = [charac for charac in engine_map[x]]
        for x in range(len(engine_map)):
            y = 0
            while y < len(engine_map[x]):
                near_asterisks = []
                curr_num = 0
                while engine_map[x][y] in nums:
                    
                    curr_num = (curr_num*10) + int(engine_map[x][y])
                    if x > 0:
                        if engine_map[x-1][y] == '*':
                            near_asterisks.append((x-1,y))
                        if y > 0:
                            if engine_map[x-1][y-1] == '*':
                                near_asterisks.append((x-1,y-1))
                        if y < len(engine_map[x])-2:
                            if engine_map[x-1][y+1] == '*':
                                near_asterisks.append((x-1,y+1))
                    if x < len(engine_map)-2:
                        if engine_map[x+1][y] == '*':
                            near_asterisks.append((x+1,y))
                        if y > 0:
                            if engine_map[x+1][y-1] == '*':
                                near_asterisks.append((x+1,y-1))
                        if y < len(engine_map[x])-2:
                           if  engine_map[x+1][y+1] == '*':
                                near_asterisks.append((x+1,y+1))
                    if y > 0:
                        if engine_map[x][y-1] == '*':
                            near_asterisks.append((x,y-1))
                    if y < len(engine_map[x])-2:
                        if engine_map[x][y+1] == '*':
                            near_asterisks.append((x,y+1))
                    y = y + 1
                y = y + 1
                for idk_gear in set(near_asterisks):
                    idk_gear = str(idk_gear)
                    if idk_gear in possible_gears:
                        possible_gears[idk_gear] = possible_gears[idk_gear] + " " + str(curr_num)
                    else:
                        possible_gears.update({idk_gear:str(curr_num)})

        for key in possible_gears:
            if len(possible_gears[key].split(" ")) == 2:
                ret_sum = ret_sum + (int(possible_gears[key].split(" ")[0]) * int(possible_gears[key].split(" ")[1]))
        return ret_sum



def get_block_val(block):
    horzontal_ind = get_horizontal_mirror(block)
    if horzontal_ind != -1:
        return horzontal_ind * 100
    
    block_rotated = [''.join(s) for s in zip(*block)]
    horzontal_ind = get_horizontal_mirror(block_rotated)
    return horzontal_ind

def get_horizontal_mirror(block):
    horzontal_ind = -1
    for ind in range(1,len(block)):
        if block[ind] == block[ind-1]:
            ind_back = ind-2
            ind_forward = ind+1
            horzontal_ind = ind
            while ind_back>=0 and ind_forward < len(block):
                if block[ind_forward] != block[ind_back]:
                    horzontal_ind = -1
                    break
                ind_back = ind_back - 1
                ind_forward = ind_forward+1
        if horzontal_ind != -1:
            return horzontal_ind
    return horzontal_ind

def get_horizontal_mirror_exclude(block, exclude):
    horzontal_ind = -1
    for ind in range(1,len(block)):
        if block[ind] == block[ind-1]:
            ind_back = ind-2
            ind_forward = ind+1
            horzontal_ind = ind
            while ind_back>=0 and ind_forward < len(block):
                if block[ind_forward] != block[ind_back]:
                    horzontal_ind = -1
                    break
                ind_back = ind_back - 1
                ind_forward = ind_forward+1
        if horzontal_ind != -1 and horzontal_ind != exclude:
            return horzontal_ind
    if horzontal_ind != exclude:
        return horzontal_ind
    return -1

def de_smudged_mirror(block: list[str], ignore):
    for x in range(len(block)):
        for y in range(len(block[x])):
            li: list[str] = list(block[x])
            new_char = "."
            if li[y] == ".":
                new_char == "#"
            li[y] = new_char
            block_changed = block.copy()
            block_changed[x] = "".join(li)
            val = get_horizontal_mirror_exclude(block_changed,ignore)
            if  val != -1:
                return val
    

def get_cleaned_block_val(block):
    dirty_rot = [''.join(s) for s in zip(*block)]
    dirty_ind = get_horizontal_mirror(block)
    dirty_ind_rot = get_horizontal_mirror(dirty_rot)
    
    clean_ind = de_smudged_mirror(block,dirty_ind)
    clean_ind_rot = de_smudged_mirror(dirty_rot, dirty_ind_rot)

    if clean_ind is not None:
        return clean_ind * 100
    elif clean_ind_rot is not None:
        return clean_ind_rot
    



def dif_lines(line_1,line_2):
    difs = []
    if len(line_1) != len(line_2):
        return -1
    for ind in range(len(line_1)):
        if line_1[ind] != line_2[ind]:
            difs.append((ind,line_1[ind],line_2[ind]))
    return difs


def solve_part_1(file):
    lines = []
    blocks = []
    with open(file, "r") as input:
        lines = [item.strip() for item in input.readlines()]

    curr_block = []
    for ind, line in enumerate(lines):
        if line == "" or ind == len(lines)-1:
            blocks.append(curr_block.copy())
            curr_block = []
        else:
            curr_block.append(line.strip())

    sum = 0
    for block in blocks:
        sum += get_block_val(block)
    return sum

def print_block(block):
    for item in block:
        print(item)       

def solve_part_2(file):
    lines = []
    blocks = []
    with open(file, "r") as input:
        lines = [item.strip() for item in input.readlines()]

    curr_block = []
    for ind, line in enumerate(lines):
        if line == "" or ind == len(lines)-1:
            blocks.append(curr_block.copy())
            curr_block = []
        else:
            curr_block.append(line.strip())

    sum = 0
    for block in blocks:
        sum += get_cleaned_block_val(block)
    return sum

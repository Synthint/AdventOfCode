import re

def solve_part_1(file):
    with open(file, "r") as input:
        sum = 0
        for line in input.readlines():
            res = re.search(r"^[a-zA-Z]*([0-9]).*([0-9])[a-zA-Z]*$|^[a-zA-Z]*([0-9])[a-zA-Z]*$",line)
            if res is None:
                raise Exception(line,"Does Not Appear To Have Any Digits")
            if res.group(3) is not None:
                sum = sum + (int(res.group(3))*10 + int(res.group(3)))
            else:
                sum = sum + (int(res.group(1))*10 + int(res.group(2)))
        return sum

def str_to_int(input):
    convert = { "zero":0, "one":1, "two":2, "three":3,
                "four":4, "five":5, "six":6, "seven":7,
                "eight":8, "nine":9, "0":0, "1":1, "2":2,
                "3":3, "4":4, "5":5, "6":6, "7":7, "8":8, "9":9, }
    return convert[input]

def solve_part_2(file):
    with open(file, "r") as input:
        sum = 0
        for line in input.readlines():
            res = re.search(r"^[a-zA-Z]*?([0-9]|one|two|three|four|five|six|seven|eight|nine|zero).*([0-9]|one|two|three|four|five|six|seven|eight|nine|zero)[a-zA-Z]*?$|^[a-zA-Z]*([0-9]|one|two|three|four|five|six|seven|eight|nine|zero)[a-zA-Z]*?$",line)
            if res is None:
                raise Exception(line,"Does Not Appear To Have Any Numbers")
            if res.group(3) is not None:
                sum = sum + (str_to_int(res.group(3))*10 + str_to_int(res.group(3)))
            else:
                sum = sum + (str_to_int(res.group(1))*10 + str_to_int(res.group(2)))
        return sum


print("Part 1 Answer ->",solve_part_1("./input.txt"))
print("Part 2 Answer ->",solve_part_2("./input.txt"))
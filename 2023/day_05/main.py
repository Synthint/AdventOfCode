from collections import OrderedDict
from multiprocessing import Pool

class Almanac_Range:
    dest: int = 0
    source: int = 0
    range_len: int = 0
    
    def __init__(self,source,dest,range_len):
        self.source = source
        self.dest = dest
        self.range_len = range_len
    
    def convert_value(self, input: int) -> int:
        conversion = -1
        if input in range(self.source,self.source+self.range_len+1):
            conversion = (input - self.source) + self.dest
        return conversion
    
    def contains_value(self, input: int) -> bool:
       return (input in range(self.source,self.source+self.range_len+1))
    
    def reverse_convert_value(self, input: int) -> int:
        conversion = input
        if input in range(self.dest,self.dest+self.range_len+1):
            conversion = (input - self.dest) + self.source
        return conversion
    
    def reverse_contains_value(self, input: int) -> bool:
       return (input in range(self.dest,self.dest+self.range_len+1))
    
    def print_out(self):
        print(f"Convert: {self.source} -> {self.dest} over {self.range_len}")

class Almanac_Conversion:
    title: str = ""
    conversion_ranges : list[Almanac_Range] = []
    
    def __init__(self,title: str, conversion_ranges: list[Almanac_Range]):
        self.title=title
        self.conversion_ranges = conversion_ranges
    
    def add_range(self, item: Almanac_Range):
        self.conversion_ranges.append(item)
        
    def contains_value(self,input: int) -> bool:
        ret = False
        for range in self.conversion_ranges:
            if range.contains_value(input):
                ret = True
        return ret
    
    def convert_value(self,input: int) -> int:
        for range in self.conversion_ranges:
            if range.contains_value(input):
                return range.convert_value(input)
        return input
    
    def reverse_convert_value(self,input: int) -> int:
        for range in self.conversion_ranges:
            if range.reverse_contains_value(input):
                return range.reverse_convert_value(input)
        return input


    def print_out(self):
        print(f"\n===== {self.title} =====\n")
        for range in self.conversion_ranges:
            range.print_out()
        print(f"\n<><>===== {self.title} =====<><>\n\n")







def process_seeds(input: str) -> list[int]:
    in_seeds = input.split(" ")[1:]
    return [int(seed) for seed in (in_seeds)]


def process_seed_ranges(input: str) -> dict[int,int]:
    out_seeds : dict[int,int]= {}
    in_seeds = [int(seed) for seed in (input.split(" ")[1:])]
    index = 0
    while index < (len(in_seeds)-1):
        out_seeds.update({in_seeds[index] : in_seeds[index+1]})
        index += 2
    return out_seeds

def solve_part_1(file):
    seeds_list=[]
    conversions : dict[str,Almanac_Conversion]= {"seed-to-soil map:" : Almanac_Conversion("seed-to-soil map:", []),
             "soil-to-fertilizer map:" : Almanac_Conversion("soil-to-fertilizer map:", []),
             "fertilizer-to-water map:" : Almanac_Conversion("fertilizer-to-water map:", []),
             "water-to-light map:" : Almanac_Conversion("water-to-light map:", []),
            "light-to-temperature map:" : Almanac_Conversion("light-to-temperature map:", []),
            "temperature-to-humidity map:" : Almanac_Conversion("temperature-to-humidity map:", []),
            "humidity-to-location map:" : Almanac_Conversion("humidity-to-location map:", []),
            }
    all_lines : list[str] = []
    with open(file, "r") as input:
        for line in input:
            all_lines.append(line)
    seeds_list = (process_seeds(all_lines[0]))
    for index in range(1,len(all_lines)):
        if all_lines[index].strip() in conversions:
            mode = all_lines[index].strip()
            index += 1
            while index < len(all_lines) and all_lines[index].strip() not in conversions and all_lines[index].strip() != "":
                nums = [int(n.strip()) for n in all_lines[index].split(" ")]
                conversions[mode].add_range(Almanac_Range(nums[1], nums[0], nums[2]))
                index += 1
        
    min_loc = 999999999999999999999999999
    for seed in seeds_list:
        converted = seed
        for mode in conversions:
            converted = conversions[mode].convert_value(converted)
        if converted < min_loc:
            min_loc = converted
    return min_loc        


def min_convert_seed_range(seeds_start:int, seeds_range:int, conversions: dict[str,Almanac_Conversion]):
    min_loc = 9999999999999999999999999999999999999
    for seed in range(seeds_start,seeds_start+seeds_range):
        converted = seed
        for mode in conversions:
            converted = conversions[mode].convert_value(converted)
        if converted < min_loc:
                min_loc = converted
    return min_loc

def solve_part_2(file):
    conversions : dict[str,Almanac_Conversion]= {"seed-to-soil map:" : Almanac_Conversion("seed-to-soil map:", []),
             "soil-to-fertilizer map:" : Almanac_Conversion("soil-to-fertilizer map:", []),
             "fertilizer-to-water map:" : Almanac_Conversion("fertilizer-to-water map:", []),
             "water-to-light map:" : Almanac_Conversion("water-to-light map:", []),
            "light-to-temperature map:" : Almanac_Conversion("light-to-temperature map:", []),
            "temperature-to-humidity map:" : Almanac_Conversion("temperature-to-humidity map:", []),
            "humidity-to-location map:" : Almanac_Conversion("humidity-to-location map:", []),
            }
    all_lines : list[str] = []
    with open(file, "r") as input:
        for line in input:
            all_lines.append(line)
    seeds_list = (process_seed_ranges(all_lines[0]))
    for index in range(1,len(all_lines)):
        if all_lines[index].strip() in conversions:
            mode = all_lines[index].strip()
            index += 1
            while index < len(all_lines) and all_lines[index].strip() not in conversions and all_lines[index].strip() != "":
                nums = [int(n.strip()) for n in all_lines[index].split(" ")]
                conversions[mode].add_range(Almanac_Range(nums[1], nums[0], nums[2]))
                index += 1
    
    
    rev_conversions = OrderedDict(reversed(list(conversions.items())))
    step = 1_000         # nums per process
    start = 60_000_000  # min num to start at
    iter_map_thing = []
    for value in range(start,70_000_000,step):
        iter_map_thing.append((value,value+step,seeds_list,rev_conversions))
    
    try:  
        pool = Pool(16)
        results = pool.starmap(check_loc_for_seeds,iter_map_thing)
        
        res_grouped = []
        for res in results:
            res_grouped.extend(res)
        #print(res_grouped)
        return min(res_grouped)
    except:
        return -1
        
def full_convert_rev(seed: int, conversions: dict[str,Almanac_Conversion]):
    converted = seed
    for mode in conversions:
        #str = f"Converting {converted} over {mode}"
        converted = conversions[mode].reverse_convert_value(converted)
        #print(str + f" -> {converted}")
    return converted


def check_loc_for_seeds(range_min, range_max ,seeds_list, conversions: dict[str,Almanac_Conversion]):
    #print(f"Range Start: {range_min}")
    contained = []
    for location in range(range_min,range_max):
        seed = full_convert_rev(location,conversions)
        #print(f"Checking {i} -> {x}")
        for key in seeds_list:
            #print(f"Checking {i} -> {x} against {key} -> {key+seeds_list[key]}")
            if seed in range(key,key+seeds_list[key]+1):
                contained.append(location)
                #print(f"-> {x} in range {key} -> {key+seeds_list[key]+1}")
                break
    return contained

#70000000
#69841803

print("Part 1 Answer ->",solve_part_1("./input.txt"))
print("Part 2 Answer ->",solve_part_2("./input.txt"))
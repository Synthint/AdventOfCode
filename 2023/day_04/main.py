import re

def solve_part_1(file):
    with open(file, "r") as input:
        total_score = 0
        for card in input:
            card_score = 0
            card_info =card.split(":")[1]
            winners = card_info.split("|")[0].split(" ")
            
            winners = [int(num) for num in winners if num != ""]
            my_nums = card_info.split("|")[1].split(" ")
            my_nums = [int(num) for num in my_nums if num != ""]
            for num in my_nums:
                if num in winners:
                    if card_score == 0:
                        card_score = 1
                    else:
                        card_score = card_score * 2
            total_score = total_score + card_score
        return total_score
                    
        
def solve_part_2(file):
    cards_list = []
    with open(file, "r") as bad_var_name:
        for elem in bad_var_name:
            cards_list.append(elem)
    for card in cards_list:
        card_score = 0
        card_number = int(card.split(":")[0].split(" ")[-1])
        card_info = card.split(":")[1]
        winners = card_info.split("|")[0].split(" ")
        
        winners = [int(num) for num in winners if num != ""]
        my_nums = card_info.split("|")[1].split(" ")
        my_nums = [int(num) for num in my_nums if num != ""]
    
        for num in my_nums:
            if num in winners:
                card_score = card_score + 1
        for index in range(1,card_score+1):
            new_card = cards_list[(card_number -1)+ index]
            cards_list.append(new_card)
    return len(cards_list)
        


print("Part 1 Answer ->",solve_part_1("./input.txt"))
print("Part 2 Answer ->",solve_part_2("./input.txt"))
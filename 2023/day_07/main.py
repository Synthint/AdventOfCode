import collections
from functools import cmp_to_key

card_values = {
    "2" : 0,
    "3" : 1,
    "4" : 2,
    "5" : 3,
    "6" : 4,
    "7" : 5,
    "8" : 6,
    "9" : 7,
    "T" : 8,
    "J" : 9,
    "Q" : 10,
    "K" : 11,
    "A" : 12

}


card_values_part_2 = {
    "2" : 0,
    "3" : 1,
    "4" : 2,
    "5" : 3,
    "6" : 4,
    "7" : 5,
    "8" : 6,
    "9" : 7,
    "T" : 8,
    "J" : -1,
    "Q" : 10,
    "K" : 11,
    "A" : 12

}


def compare_two_hands(hand_1, hand_2):
    hand_1 = hand_1[:-1]
    hand_2 = hand_2[:-1]

    ind = 0
    while card_values[hand_2[ind]] == card_values[hand_1[ind]]:
        ind += 1
        if ind == 5:
            return 0

    value_list_1 = []
    high_1 = card_values[hand_1[ind]] 
    value_list_2 = []
    high_2 = card_values[hand_2[ind]] 
    for val in card_values:
        if hand_1.count(val) > 0:
            value_list_1.append(hand_1.count(val))
                
        if hand_2.count(val) > 0:
            value_list_2.append(hand_2.count(val))

    high_diff = high_1 - high_2
    value_list_1.sort()
    value_list_2.sort()

    value_list_1 = [i for i in value_list_1 if i != 1]
    value_list_2 = [i for i in value_list_2 if i != 1]

    while True:
        if len(value_list_1) == 0 and len(value_list_2) == 0:
            return int(high_diff / abs(high_diff))
        elif len(value_list_1) == 0 and len(value_list_2) != 0:
            return -1
        elif len(value_list_1) != 0 and len(value_list_2) == 0:
            return 1
        elif max(value_list_1) > max(value_list_2) :
            return 1
        elif max(value_list_2) > max(value_list_1):
            return -1
        elif len(value_list_1) == len(value_list_2) and len(value_list_1) == 1 and max(value_list_2) == max(value_list_1):
            return int(high_diff / abs(high_diff))
        else:
            value_list_1.pop()
            value_list_2.pop()


def compare_hand_types(hand_1, hand_2):
    hand_1 = hand_1[:-1]
    hand_2 = hand_2[:-1]

    value_list_1 = []
    value_list_2 = []
    for val in card_values:
        if hand_1.count(val) > 0:
            value_list_1.append(hand_1.count(val))
                
        if hand_2.count(val) > 0:
            value_list_2.append(hand_2.count(val))

    value_list_1.sort()
    value_list_2.sort()

    value_list_1 = [i for i in value_list_1 if i != 1]
    value_list_2 = [i for i in value_list_2 if i != 1]

    while True:
        if len(value_list_1) == 0 and len(value_list_2) == 0:
            return 0
        elif len(value_list_1) == 0 and len(value_list_2) != 0:
            return -1
        elif len(value_list_1) != 0 and len(value_list_2) == 0:
            return 1
        elif max(value_list_1) > max(value_list_2) :
            return 1
        elif max(value_list_2) > max(value_list_1):
            return -1
        elif len(value_list_1) == len(value_list_2) and len(value_list_1) == 1 and max(value_list_2) == max(value_list_1):
            return 0
        else:
            value_list_1.pop()
            value_list_2.pop()  


def compare_two_hands_jokers(hand_1, hand_2):
    if hand_1.count("J") == 0 and hand_2.count("J") == 0:
        return compare_two_hands(hand_1,hand_2)
    
    hand_1_possibilities = []
    hand_2_possibilities = []
    for val in card_values_part_2:
            if val != "J":
                hand_1_possibilities.append([card if card != "J" else val for card in hand_1])
                hand_2_possibilities.append([card if card != "J" else val for card in hand_2])

    hand_1_possibilities.sort(key=cmp_to_key(compare_two_hands))
    hand_2_possibilities.sort(key=cmp_to_key(compare_two_hands))


    comp = compare_hand_types(hand_1_possibilities[-1],hand_2_possibilities[-1])
    if comp != 0:
        return comp
    else:
        ind = 0
        while card_values_part_2[hand_2[ind]] == card_values_part_2[hand_1[ind]]:
            ind += 1
            if ind == 5:
                return 0
        high_1 = card_values_part_2[hand_1[ind]] 
        high_2 = card_values_part_2[hand_2[ind]] 
        high_diff = high_1 - high_2
        return int(high_diff / abs(high_diff))



def solve_part_1(file):
    hands = []
    with open(file, "r") as input:
        for line in input:
            cards = list(line.split(" ")[0])
            bid = line.split(" ")[1].strip()
            cards.append(int(bid))
            
            hands.append(cards)
    
    hands.sort(key=cmp_to_key(compare_two_hands))
    mult_sum = 0
    mult = 1
    for hand in hands:
        mult_sum += mult*hand[-1]
        mult += 1
    return mult_sum
        

def solve_part_2(file):
    hands = []
    with open(file, "r") as input:
        for line in input:
            cards = list(line.split(" ")[0])
            bid = line.split(" ")[1].strip()
            cards.append(int(bid))
            
            hands.append(cards)
    
    hands.sort(key=cmp_to_key(compare_two_hands_jokers))
    mult_sum = 0
    mult = 1
    for hand in hands:
        mult_sum += mult*hand[-1]
        mult += 1
    return mult_sum
        
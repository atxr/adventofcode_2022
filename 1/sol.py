#!/usr/bin/env python
from functools import reduce

"""
Aux function to reduce the lines
If the value is null, append 0 to the accumulator (new elf)
Otherwise, add the value to the last element of the list 
"""
def aux(acc, x):
    x = x.strip()
    if x == '':
        acc.append(0)
    else:
        acc[-1] += int(x)
    return acc

"""Return the calories carried by the strongest elves"""
def main():
    inp = open('input', 'r').readlines()
    calories = sorted(reduce(aux, inp, [0]), reverse=True)
    return calories[0], sum(calories[:3])

if __name__ == '__main__':
    top1, top3 = main()
    print(f"The strongest elf carries {top1} calories!")  # +1 because of python indexing
    print(f"The top 3 elves carry {top3} calories!")


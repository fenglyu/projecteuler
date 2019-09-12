#!/usr/bin/env python3
#
#   $ python3 14.py |sort -k2 -nr|head
#   837799 524
#   626331 508
#   939497 506
#   704623 503
#   927003 475
#   910107 475
#   511935 469
#   796095 467
#   767903 467
#   970599 457
#
def count(num):
    c = 0
    while True:
        if num <= 1:
            break
        if num % 2 == 0:
            num = num / 2
        else:
            num = num * 3 + 1
        c += 1
    return c

def main():
    for i in range(800000, 1000000):
        print((i, count(i)))

if __name__ == "__main__":
    main()

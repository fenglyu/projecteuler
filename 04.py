#!/usr/bin/env python3

def MAX(a, b):
    return a if a > b else b

def ispalind(num):
    ori, rev = num, 0
    while ori > 0:
        rev = rev * 10 + ori % 10
        ori = int(ori/10)
    return rev == num

def main():
    max = 0
    for i in range(999, 100, -1):
        limit = MAX(int(max/i), 100)
#       print("limit:%d, max:%d, i:%d" % (limit, max, i))
        for j in range(999, limit, -1):
            num = i * j
            if ispalind(num):
#                print("palind:%d" % num)
                max = MAX(max, num)
    print(max)

if __name__ == "__main__":
#    print(ispalind(91200219))
    main()

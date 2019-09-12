#!/usr/bin/env python3


def divisor(n):
    num = 0
    i = 1
    while i * i <= n:
        if n % i == 0:
            num += 2
        i += 1

    if i * i == n:
        num -= 1

    return num

def main():

    i = 1
    sum = 0
    while True:
        sum = divisor(i*(i+1)/2)
        if sum >=  500:
            print((sum, i*(i+1)/2))
            break
        i += 1

if __name__ == "__main__":
    main()

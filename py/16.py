#!/usr/bin/env python3


def main():

    a = 1

    b = 2**1000

    sum = 0

    while b != 0:
        sum = sum + int(b%10)
        b = int( b/10)
        print(sum, b)

    print(sum)


if __name__ == "__main__":
    main()

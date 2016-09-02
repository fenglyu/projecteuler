#!/usr/bin/env python3

def prime(n):
    i = 2;
    if n <= 1:
        return False

    while i * i <= n:
        if n % i == 0:
            return False
        i = i + 1
    return True

def main():

    sum = 2
    for i in range(3, 2000001, 2):
        if prime(i):
#            print(i)
            sum += i
#        print(sum)

    print("result is %d\n" % sum)

if __name__ == "__main__":
    main()


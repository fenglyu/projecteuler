#!/usr/bin/env python3


def prime(n):
    i = 2
    flag = True
    while i*i <= n:
        if n%i != 0:
            i += 1
        else:
            flag = False
#            i += 1
            break
    return flag

def index(m):
    i = 2
    count = 0
    while True:
        if prime(i):
            count += 1
    #        print(i, count)
            if count >= m:
                break;
        i += 1
    print(i, count)

if __name__ == "__main__":
#    print(prime(113))
#    print(prime(122))
    index(10001)

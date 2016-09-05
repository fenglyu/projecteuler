#!/usr/bin/env python3


def work(n):
    arr = []
    for i in range(1, n+1):
        if n % i == 0:
            arr.append(i)
#    print(arr)
    return len(arr)

def main():

    i = 1
    sum = 0
    count = 0
    while True:
#        print(i)
#        if (count >= 10):
#            break
        sum += i
        arr = work(sum)
        if arr >= 500:
            print(sum)
            print(arr)
            break
        i += 1
        count += 1

if __name__ == "__main__":
    main()

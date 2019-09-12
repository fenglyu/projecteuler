#!/usr/bin/python

arr=list()
prime, num = 3, 600851475143
while prime < 600851475143:
    if (prime * prime >= 600851475143):
        break
    if num % prime != 0:
        prime += 2
        continue
    else:
        arr.append(prime)
        num /= prime
        continue

print(arr)

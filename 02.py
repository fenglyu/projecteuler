#!/usr/bin/python


pre, cur = 0, 1
total = 0
while True:
    pre, cur = cur, pre + cur
    if cur >= 4000000:
        break
    if cur % 2 == 0:
        total += cur

print total

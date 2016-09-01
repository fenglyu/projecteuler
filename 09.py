#!/usr/bin/env python3

#
#   a^2 + b^2 = c^2
#   c = 1000 - (a+b)
#   results in : 1000*1000 - 2000*(a+b) + 2*a*b == 0

def main():
    for i in range(1,1000):
        for j in range(1,1000):
            if 1000*1000 - 2000*(i+j) + 2*i*j == 0:
                print(i * j * (1000 -i -j))

if __name__ == "__main__":
    main()

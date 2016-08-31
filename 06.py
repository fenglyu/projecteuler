#!/usr/bin/evn python3


def main():
    sum1 = 0
    for i in [ x*x for x in range(1, 101) ]:
        sum1 += i
#    print(sum1)

    sum2 = 0
    for i in [ x for x in range(1, 101) ]:
        sum2 += i

    sum2 = sum2 * sum2
    print(sum2 - sum1)


if __name__ == "__main__":
    main()

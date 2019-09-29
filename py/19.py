import time

def checkLeapYear(year):
    if ((year % 4 ==0) and (year % 100 != 0)) or (year % 400 == 0):
        return True
    return False

def getDaysInMonth (year, month):
    #      month   1   2   3   4   5   6   7   8   9  10  11  12
    daysInMonth_list = [31, 0, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]

    if month != 2:
        return daysInMonth_list[month-1]
    else:
        if True == checkLeapYear(year):
            return 29
        else:
            return 28

def getDaysInYear(year):
    if True == checkLeapYear(year):
        return 366
    else:
        return 365

def getDayOfWeek(year, month, day):
    totalDays = 0
    for y in range(1900, year):   # from 1900 to (year-1)
        totalDays += getDaysInYear(y)
    for m in range(1, month):     # from Jan to (month-1)
        totalDays += getDaysInMonth(year, m)
    totalDays += day              # from 1 to (day-1)
    dayOfWeek = totalDays % 7     # 1 Jan 1900 was a Monday

    return dayOfWeek

def getTotalOfSundays():
    totalOfSundays = 0
    for year in range(1901, 2001):
        for month in range(1, 13):
            # check whether the fist day of the month is Sunday
            if 0 == getDayOfWeek(year, month, 1):
                totalOfSundays += 1

    return totalOfSundays

def main():
    start = time.process_time()

    assert 1 == getDayOfWeek(1900, 1,  1)

    print(getTotalOfSundays(),"Sundays fell on the first of the", end='')
    print("month during the twentieth century.")

    end = time.process_time()
    print('CPU processing time : %.5f' %(end-start))

if  __name__ == '__main__':
    main()

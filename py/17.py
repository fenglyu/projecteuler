def getDigits(number):
    digits = [0]*4
    numOfDigits = 0
    while number >= 1:                           
        digits[numOfDigits] = number%10
        number      /= 10
        numOfDigits += 1
    return digits
        
def getLettersSumOfNumber(number):
    """
    C_blank = 0   # blank
    C_1 =     1   # one
    C_2 =     2   # two
    C_3 =     3   # three
    C_4 =     4   # four
    C_5 =     5   # five
    C_6 =     6   # six
    C_7 =     7   # seven
    C_8 =     8   # eight
    C_9 =     9   # nine
    C_10 =    10  # ten
    C_11 =    11  # eleven
    C_12 =    12  # twelve
    C_13 =    13  # thirteen
    C_14 =    14  # fourteen
    C_15 =    15  # fifteen
    C_16 =    16  # sixteen
    C_17 =    17  # seventeen
    C_18 =    18  # eighteen
    C_19 =    19  # nineteen
    C_20 =    20  # twenty
    C_30 =    21  # thirty 
    C_40 =    22  # forty
    C_50 =    23  # fifty
    C_60 =    24  # sixty
    C_70 =    25  # seventy
    C_80 =    26  # eighty
    C_90 =    27  # ninety
    C_100 =   28  # hundred
    C_1000 =  29  # thousand 
    C_and  =  30  # and
    """
    numberString = ["", "one", "two", "three", "four", "five",
        "six", "seven","eight", "nine", "ten", "eleven", "twelve", "thirteen", 
        "fourteen","fifteen","sixteen", "seventeen", "eighteen", "nineteen", 
        "twenty","thirty", "forty", "fifty", "sixty", "seventy", "eighty", 
        "ninety","hundred", "thousand", "and"]
    
    letters_list = [ len(item) for item in numberString ]
    
    digits = getDigits(number)

    unitDigit    = int(digits[0])
    decadeDigit  = int(digits[1])
    hundredDigit = int(digits[2])

    sumOfLetters = 0

    C_1, C_19, C_100, C_1000, C_and = 1, 19, 28, 29, 30

    if 1000 == number:
        sumOfLetters += letters_list[C_1000] + letters_list[C_1]
    else:
        if hundredDigit > 0:
            sumOfLetters += letters_list[hundredDigit]+letters_list[C_100]
            if decadeDigit+unitDigit > 0:
                sumOfLetters += letters_list[C_and]

        if decadeDigit > 1: 
            sumOfLetters += letters_list[decadeDigit+C_19-1]
        elif decadeDigit == 1:
            sumOfLetters += letters_list[10+unitDigit]

        if unitDigit > 0 and decadeDigit != 1:
            sumOfLetters += letters_list[unitDigit]

    return  sumOfLetters

def getSumOfLetters(max_number):
    sumOfLetters = 0;
    for i in range (1, max_number+1):
        sumOfLetters += getLettersSumOfNumber(i)
    return sumOfLetters
    
def main():
    assert 20 == getLettersSumOfNumber(115)
    assert 23 == getLettersSumOfNumber(342)
    print("The letters", getSumOfLetters(1000), "are used for all the numbers ")
    print("from 1 to 1000 (one thousand) inclusive in words.") 

if  __name__ == '__main__':
    main()


#!/usr/bin/python3


fil = {}
for i in range(2, 101):
    for  j in range(2, 101):
        s = i**j
#        if fil.get(s) and fil[s] > 1:
#            print(i, j)
#        else:
#            fil[s] = 0
        if fil.get(s):
            fil[s]['c'] +=1 
            print(fil[s]['p'], "=>", i, j)
        else:
            fil[s] = {}
            fil[s]['c'] = 1 
            fil[s]['p'] = (i,j)



#for k in fil.keys():
#    v = fil[k]
#    if v > 1:
#        print(k, v)

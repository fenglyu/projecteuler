## Compile project 29
```
flv@zenyata:~/dev/projecteuler/c$ echo $CXX
clang++-10
flv@zenyata:~/dev/projecteuler/c$ $CXX 29_01.cc
flv@zenyata:~/dev/projecteuler/c$ ./a.out
9183

```


```
 in sweep()).


Here's the theory behind my code:

For a1 b1 to be the same number as a2 b2, a1 and a2 must be powers of some common base a0. If a1 = a0 m1 and a2 = a0 m2, then we also have m1 b1 = m2 b2. For example, 46 = (22)6 = 212 = (23)4 = 84 (2 is common base here).

Suppose now that we check powers of a where a goes from 2 upwards. If we come to a value of a that is a power of some lower number, say a = a0n, then we must have calculated a as this power earlier (for the smaller base a0). So we could mark some powers of a as duplicates. The way I do this might seem a bit lengthy, but it works...

When I get a value ab that is in the interval 2 to 100 (e.g. 23 = 8), I study  powers of ab: a2b, a3b, a4b, ..., a100b. For each of these I check if it could be written as a power of some other power of a; if m b has a factor f, then amb = (af)mb/f. If now af < 100, then we have duplicates, in which case I mark (af)mb/f as a duplicate [it's a duplicate of amb].
The procedure in this paragraph is briefly what mark_duplicates() does.


Actually, I tried to do in Haskell what bitRAKE did in Mathematica, but I ran into problems; the garbage collector couldn't release enough memory.
```

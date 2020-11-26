#include <iostream>

using namespace std;
 
const int MAX = 100;
 
bool duplicate[MAX][MAX];
 
void init()
{
    for(int a=2; a<=MAX; a++)
    for(int b=2; b<=MAX; b++)
        duplicate[a-1][b-1] = false;
}
 
int power(int a, int b)
{
    int res = 1;
    for(int i=0; i<b; i++)
        res *= a;
        
    return res;
}
 
void mark_duplicates(int a1, int b1)
{
    for(int b2=2; b2<=MAX; b2++) {
        int b3 = b1 * b2;
        for(int f1=1; f1<b1; f1++) {
            if(b3 % f1 == 0) {
                int f2 = b3 / f1;
              int a2 = power(a1, f1);
              if(a2 <= MAX) {
                    if(f2 <= MAX) {
                        duplicate[a2-1][f2-1] = true;
                  }
              }
            }
        }
    }
}
 
void sweep(int a)
{
    for(int b=2; b<=MAX; b++) {
        int c = power(a, b);
        if(c > MAX)
            break;
        mark_duplicates(a, b);
    }
}
 
int count_nondups()
{
    int count = 0;
    for(int a=2; a<=MAX; a++)
        for(int b=2; b<=MAX; b++)
            if(!duplicate[a-1][b-1]) {
                count++;
            }
                
    return count;
}
 
int main(int argc, char *argv[])
{
    init();
  
    for(int a=2; a<=MAX; a++)
        sweep(a);
    
    std::cout << count_nondups() << std::endl;
    return 0;
}


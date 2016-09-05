#include <stdio.h>
/*
 *   gcc -O3 -o 12 12.c # -g -W
 */

unsigned long long arr[600] = {0};

int prime(unsigned long long n){
    unsigned long long i = 1;
    int count = 0;

    while(i <= n){
        if( n % i  == 0){
            arr[count++] = i;
        }
        i++;
    }
    return count; 
}

void display(unsigned long long arr[], int size){
    int i;
    printf("[ ");
    for(i = 0; i < size; i++){
        if(arr[i] > 0 ){
            printf("%ld ", arr[i]);
        }
    }
    printf(" ]\n");
}

int main(){

    unsigned long i = 1;
    unsigned long long sum = 0;
    
    while(1){
        sum += i;
        int s =  prime(sum);
        display(arr, 600);
        if( s >= 500){
            printf("result is %d\n ", sum);
            //display(arr, 600);
            break;
        } 
        i++;
    }
}

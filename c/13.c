#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>


#define LEN 80
#define COL 50

#define err_abort(str...) do{ \
    printf("%s", str); \
    exit(1); \
}while(0)\

int char2int(char c){
    int i ;
    switch (c){
        case '0': i = 0; break;
        case '1': i = 1; break;
        case '2': i = 2; break;
        case '3': i = 3; break;
        case '4': i = 4; break;
        case '5': i = 5; break;
        case '6': i = 6; break;
        case '7': i = 7; break;
        case '8': i = 8; break;
        case '9': i = 9; break;
        default:
            i = -1;
            break;
    }
    return i;
}

int main(){
    
    char arr[100][50];
    int n=0;
    FILE *fp;
    char * line = NULL;
    size_t len = 0;
    ssize_t read;
    
    fp = fopen("13.txt", "r");
    if (fp == NULL)
        exit(EXIT_FAILURE);

    while ((read = getline(&line, &len, fp)) != -1){
//        printf("Retrieved line of length %zu :\n", read);
//        printf("%s", line);
        memmove(arr[n++], line, COL);
    } 


    if (line)
        free(line);

    printf("echo :%i", INT_MAX);
    int i,j;
    int sqare[100][50];
    unsigned long long  tmp = 0;
    int res[100] = {0}; 
//    for(i = 0; i< 100; i++){
//        for(j =0; j<50; j++){
//            //printf(" %c", arr[i][j]);
//            printf("%d", char2int(arr[i][j]));
//   //         sqare[i][j] = char2int(arr[i][j]);
//            tmp = char2int(arr[i][j]); 
//        }
//        printf("\n");
//    }

    for (i = 0; i < 100; i++){
        for (j = 0; j < 50; j++){
            tmp +=  tmp * 10 + char2int(arr[i][j]);
        }
        res[i] = tmp;
        printf("%ld\n", tmp);
    }

}

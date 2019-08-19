#include<stdio.h>
#include<limits.h>
#include<stdint.h>
int reverse(int x) {        
    int64_t res = 0;
    int64_t val = x;
    while(val != 0){
        res = 10 * res + val % 10;
        if (res > INT_MAX || res < INT_MIN){
            return 0;
        }
        val = val/10;
    }
    return res;
}
int main(){
    printf("%d \n",reverse(-12003000));
    return 0;
}

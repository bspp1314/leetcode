#include <stdio.h>
#include <malloc.h>
int* findErrorNums(int* nums, int numsSize, int* returnSize) {
    int *result = (int *)malloc(sizeof(int)*2);
    int map[20000] = {0};

    for(int i = 0;i < numsSize;i++){
        if (map[nums[i]] == 1) {
            result[0] = nums[i];
        }else{
            map[nums[i]] = 0;
        }
    }
    for(int i = 0;i < numSize;i++){
        if (map[i+1] == 0) {
            result[1] = i+1;
        }
    }
    //*returnSize = 2; 
    return result;
}

int main(){

    int array[10] = {1,2,3,4,4,9,8,7,6,10};
    int *val = 
}

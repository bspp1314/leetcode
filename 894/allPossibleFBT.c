#include<stdio.h>
#include<malloc.h>
#include<stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};


    struct TreeNode** 
allPossibleFBT(int N, int* returnSize)
{
    struct TreeNode **res = (struct TreeNode **)(malloc(sizeof(struct TreeNode *)));
    if (N %2 == 0) {
        return res;
    }else if (N == 1) {
        struct TreeNode *root = (struct TreeNode*)malloc(sizeof(struct TreeNode));
        root->left = NULL;
        root->right = NULL;
        root->val =  0;
        res = realloc(res, sizeof(struct TreeNode) * ((*returnSize) + 1));
        res[(*returnSize)++] = root;
        return res; 
    }else{
        for (int i = 1;i < N/2 +1;i+=2){
            int lsize = 0;
            int rsize = 0;
            struct TreeNode **lres = allPossibleFBT(i,&lsize);
            struct TreeNode **rres = allPossibleFBT(N-i-1,&rsize);
            for (int j = 0;j < lsize;j++){
                for (int k = 0;k < rsize;k++){
                    struct TreeNode *root = (struct TreeNode*)malloc(sizeof(struct TreeNode));
                    root->left = NULL;
                    root->right = NULL;
                    root->val =  0;
                    root->left = lres[j];
                    root->right = rres[k];
                    res = realloc(res, sizeof(struct TreeNode) * ((*returnSize) + 1));
                    res[(*returnSize)++] = root;
                    if (i != (N-i-1)){
                        struct TreeNode *root2 = (struct TreeNode*)malloc(sizeof(struct TreeNode));
                        root2->left = NULL;
                        root2->right = NULL;
                        root2->val =  0;
                        root2->right = lres[j];
                        root2->left = rres[k];
                        res = realloc(res, sizeof(struct TreeNode) * ((*returnSize) + 1));
                        res[(*returnSize)++] = root2;
                    }
                }
            } 
        }
        return res;
    }
}

int
main(int argc,char *argv[]){
    int returnSize = 0;
    for int i = 0;i <= 20;i++ {
        allPossibleFBT(i,&returnSize);
        printf("returnSize %d \n",returnSize);
        *returnSize = 0;
    }
}

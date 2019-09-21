#include<stdio.h>
#include<malloc.h>
#include<math.h>


int* twoSum(int* nums, int numsSize, int target) 
{
    for(int i = 0;i < numsSize;i++){
        for (int j = i+1;j < numsSize;j++){
            if ((nums[i]+nums[j]) == target) {
                int  *result = (int *)malloc(sizeof(int) * 2); 
                result[0] = nums[i];
                result[1] = nums[j];
                return result;
            }   
        }   
    }   

    return NULL;

}
typedef struct node{
  int key;
  int value;
  struct node *next;
}node_t;

typedef struct hash_table {
  int capacity;
  node_t **hashtab;
}hash_table_t;

hash_table_t *hash_table_create(int capacity)
{
  hash_table_t *hash_table = (hash_table_t *)malloc(sizeof(hash_table_t));
  hash_table->hashtab = (node_t **)malloc(sizeof(node_t*) * capacity);
   for(int i = 0;i < capacity;i++)
       hash_table->hashtab[i] = NULL;
   
  hash_table->capacity = capacity;
    return  hash_table;
}

int get_hash_key(int n, int capacity){
  return abs((65535 * n + 1048575)) % capacity;
}

int hash_table_value(hash_table_t *hash_table,int key)
{
  int index = get_hash_key(key,hash_table->capacity);
  node_t *curr = hash_table->hashtab[index];
  while(curr != NULL){
    if(curr->key == key)
      return curr->value;
    curr = curr->next;
  }

  return -1; 
}
void hash_table_add(hash_table_t *hash_table,int key,int value)
{
  int index = get_hash_key(key,hash_table->capacity);//get hash key

  node_t *node = (node_t*)malloc(sizeof(node_t));
  node->next = NULL;
  node->next   = hash_table->hashtab[index];
  node->key    = key;
  node->value  = value;
  hash_table->hashtab[index] = node;
}

void hash_table_free(hash_table_t* hash_table)
{
  for(int i=0;i<hash_table->capacity;i++){
    node_t* curr = hash_table->hashtab[i];
    while(curr!=NULL){
      node_t* hold = curr;
      curr->next = curr;
      free(hold);
    }   
  }
  free(hash_table->hashtab);
  free(hash_table);
}

int* twoSum2(int* nums, int numsSize, int target)
{
  hash_table_t* hash_table = hash_table_create(numsSize/2);

  for(int index=0;index<numsSize;index++){
    int idx = hash_table_value(hash_table, target - nums[index]);

    if(idx>=0){
      int* result = (int*)malloc(sizeof(int)*2);
      result[1] = index;
      result[0] = idx;
      return result;

    }

    hash_table_add(hash_table, nums[index], index);

  }

  return NULL;
}
int main(){
	return 0; 
}

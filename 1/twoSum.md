## 题目
给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。

你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

示例:给定 nums = [2, 7, 11, 15], target = 9因为 nums[0] + nums[1] = 2 + 7 = 所以返回 [0, 1]
## 暴力法
简单来说只要遍历每个元素x,并查找是否存在一个值和
target-x相等就可以了。但是缺点也是很明显的，它的时间复杂度达到了O(n2),这显然不
是我们要的。
### C语言版
```c
twoSum(int* nums, int numsSize, int target) 
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

```
### Go语言版
```go
func twoSum(nums []int, target int) []int {
    l := len(nums)
    result := make([]int, 2, 2)
    for i := 0; i < l; i++ {
        for j := i + 1; j < l; j++ {
            if (nums[i] + nums[j]) == target {
                result[0] = i 
                result[1] = j 
                return result
            }   
        }   
    }   
    return result
}
```
## Hash法
 　实际上，当我们遍历数组去查找和target-x相等的数时候，每一次都要从头开始遍历。而其实
在x前面的数其实已经被访问过多次了，这种行为确实是比较浪费时间的。而这时候我们需要一个
表来记录已经访问过的数据，并且访问这个表中的数据的时间复杂度越低越好，而Hash表显然符合
这个性质，并且其访问数据在理论在可以达到O(1)。
### C语言版
```C
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

int* twoSum(int* nums, int numsSize, int target)
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
```
### Go语言版
```go
//Hash
func twoSum(nums []int, target int) []int {
    l := len(nums)
    result := make([]int, 2, 2)
    tmp := make(map[int]int)                                                                                                                                                                                        
       
    for i := 0; i < l; i++ {
        if j, ok := tmp[target-nums[i]]; ok {
            result[0] = j 
            result[1] = i 
            return result
        } else {
            tmp[nums[i]] = i 
        }   
    }  
    return result
}   


```
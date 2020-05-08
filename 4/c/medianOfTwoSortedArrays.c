#include<stdio.h>
#include<limit.h>

//n1 + n2 >= k 
//寻找另个数组中第k大的数
int findKthSmallest(int* arr1, int n1, int* arr2, int n2,int k){
    if (n1 == 0 ){
        return arr2[k-1]
    }

    if (n2 == 0){
        return arr1[k-1]
    }

    if (k == 1 ){
        return arr1[0] < arr2[0]? arr1[0]:arr2[0];
    }

    //分摊到arr1上的第k位 
    int idx1 = (l1/(l1+l2)) * (k-1);
    int idx2 = k-(idx1+1)-1;

    if (arr1[idx] == arr2[idx]) {
        return arr1[idx1];
    }else if (arr1[idx1] < arr2[idx2]){ //k in left idx1 and righ idx2 
        return findKthSmallest(arr1+idx1+1, n1-idx1-1, arr2, idx2+1, k-idx1-1);
    }else {
        return findKthSmallest(arr1, idx1+1, arr2+idx2+1, n2-idx2-1, k-idx2-1);
    }
}
double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size) {
    int mid = 0;

    mid = (nums1Size + nums2Size)/2;
    if((nums1Size + nums2Size)%2 == 1)
        return findKthSmallest(nums1,nums1Size,nums2,nums2Size,mid+1);
    else{
        int m1 = findKthSmallest(nums1,nums1Size,nums2,nums2Size,mid);
        int m2 = findKthSmallest(nums1,nums1Size,nums2,nums2Size,mid+1);
        return (m1+m2)/2.0;

    }


}

double findMedianSortedArrays2(int* nums1, int m, int* nums2, int n) {
    if (m > n){
        findMedianSortedArrays2(nums2,n,nums1,m);
    }

    int low = 0;
    int high = nums1Size;
    while(low < high){
        int i = low + (high-low) /2 ; // i is num1  left len 
        int j =  (m+n+1)/2 -i ;       // j is num2 left len 
        int leftMaxA = i == 0 ? MIN_INT:nums[i-1]; 
        int rightMinA = i == m ? MAX_INT:nums[i]; 
        
        int leftMaxB = j == 0 ? MIN_INT:nums2[j-1]; 
        int rightMinB = j == n ? MAX_INT:nums2[j];
        
        if (leftMaxA <=rightMinB && leftMaxB <= rightMinA){
            
        }
    
    }

}


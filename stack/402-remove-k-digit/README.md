

/*
 * To tackle this problem, I break it to 4 parts:
 *
 * I: Think about the simplest case where k = 1:
 * In this case, we need to remove the first "peak" element, where "peak"
 * is defined as A[p] > A[p-1] && A[p] < A[p+1].
 * Proof: Suppose the original value is A
 *        removing non-peak digit at index p yields to optimal value A1
 *        since A[p] is non peak value, let's say A[p] < A[p-1]
 *        removing index p-1 yields to value A2.
 * 	      Therefore, A1 and A2 differs only by one digit, where in A1 it's A[p-1] and
 *        in A2 it's A[p] (idx p is removed in A1 and idx p-1 is removed in A2)
 *        Since A1 is optimal, A1 < A2; but since A[p-1] > A[p], A1 > A2. This contradicts
 *        with original assumption.
 *        Hence, removing peak element will yields to optimal solution.
 *
 * II: Removing k digits is just repeating problem (I) for k times.
 *
 * III: To speed up (II), when removing an element when it's peak, we compare the previous element
 * and see if it's also a peak element. For example:
 * 1 2 3 6 8 9 5 7 4 1 1, k = 5:
 *    1) 9 is inspected and removed when 5 is seen;
 *    2) 8 is inspected and removed while processing 5;
 *    3) 6 is inspected and removed while processing 5;
 *    4) 3 is inspected and is good;
 *    5) 7 is inspected and removed while processing 4;
 *    6) 5 is inspected and removed while processing 4;
 *    7) 5 chars had been removed, append rest to result;
 *    To solve (III), a stack is built to track all not-removed element. When processing a new digit, inspect
 *    top of stack and remove it while current digit is smaller than top of stack.
 *    Note that for case 1 2 3 4 5 6 and k = 3, the peak element is 6 5 4, an additional iteration of removing
 *    peak element needs to be performed after processing last digit of input to use up all remaining k. This
 *    is done by only picking up num.length() - k elements in result list.
 *
 * IV: Trim all 0s at beginning.
 */




### 其他答案
1. https://leetcode.com/problems/remove-k-digits/discuss/88678/Two-algorithms-with-detailed-explaination
2. 
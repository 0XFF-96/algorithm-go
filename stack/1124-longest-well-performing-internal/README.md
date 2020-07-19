

1。 Make a new array A of +1/-1s corresponding to if hours[i] is > 8 or not. The goal is to find the longest subarray with positive sum



2。 
Using prefix sums (PrefixSum[i+1] = A[0] + A[1] + ... + A[i]), you need to find for each j, the smallest i < j with PrefixSum[i] + 1 == PrefixSum[j].



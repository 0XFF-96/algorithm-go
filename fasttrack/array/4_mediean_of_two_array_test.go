package array


// 非常难和重点的题目
// 边界条件特别多
// https://www.youtube.com/watch?v=LPFhl65R7ww&t=5s
//


// 周末补一下这个吧
//
// 补课： 1、permuation a string https://www.youtube.com/watch?v=GCm7m5671Ps
// 2、https://www.youtube.com/watch?v=0NAASclUm4k 【数学问题】
// 3、https://www.youtube.com/watch?v=KukNnoN-SoY


//
//Free 5-Day Mini-Course: https://backtobackswe.com
//Try Our Full Platform: https://backtobackswe.com/pricing
//Join Early Beta For OfferFeed (Coming May 2020): https://offerfeed.io
//
//Question: Given a permutation of a sequence, calculate the next permutation in that sequence (if the permutation given is the last one, just return an empty array since it has no next permutation...it is the last permutation).
//
//The code: https://github.com/bephrem1/backtobac...
//
//
//Approach 1 (Brute Force)
//
//We can compute all permutations and stop on the permutation right after the permutation given.
//
//This can have us generate n! permutations in the worst case (the permutation given is the last one).
//
//
//Approach 2 (Use Intuition & Patterns)
//
//Permutation Sequence Example:
//
//[0, 1, 2]
//
//1.)   [0, 1, 2]
//2.)   [0, 2, 1]
//3.)   [1, 0, 2]
//4.)   [1, 2, 0]
//5.)   [2, 0, 1]
//6.)   [2, 1, 0]
//
//This approach will use the idea that we notice patterns.
//
//Notice how we plant the 0, the 1, then the 2 in the first slot as we go through the sequence.
//
//Also, notice how the last element is the array completely reversed.
//
//These may not matter but we are just piecing things together right now.
//
//Let's formulate a plan.
//
//[6, 2, 1, 5, 4, 3, 0]
//
//We know to get to this permutation we
//
//rooted 6
//rooted 2
//rooted 1
//
//If you remember back to generating permutations it was all about planting items and then picking from a pool of remaining items.
//
//When we plant 1 we have a pool like so [5, 4, 3, 0].
//
//Remember how the last permutation was the original pool reversed? We will look for the longest reversed pool because we know that it is the last permutation for that particular rooting.
//
//This is [5, 4, 3, 0] in the example given. This is a maximum suffix after the planted 1.
//
//1 is of interest since this means we have exhausted the possibilities of permutations with 1 rooted where it is since it's suffix following it is strictly decreasing.
//
//SO TO GET THE NEXT PERMUTATION WE SWAP 1 AND THE SMALLEST NEXT ELEMENT TO MINIMIZE CHANGE IN THE PERMUTATION.
//
//We are considering back to the rooting [6, 2] which has a possible pool of [0, 1, 2, 3, 5]
//
//0 got it's turn in index 2
//1 got it's turn and is finished
//
//SO THE NEXT LARGEST ELEMENT IN 1's SUFFIX IS NEXT TO BE PLANTED THERE.
//
//Swapping yields [6, 2, 3, 5, 4, 1, 0].
//
//NOT DONE.
//
//Now the prefix has been advanced in the smallest way possible. BUT the suffix of our choice to plant 3 may not be the smallest it could be to be the next permutation.
//
//NOTICE: The first permutation was in STRICTLY INCREASING order. We need to get our suffix in strictly increasing order
//We don't need to use a sorting algorithm since that is expensive.
//
//Remember, the suffix was in strictly decreasing order before, so all we need to do is reverse it and we will have a sorted suffix.
//
//When you really grasp how permutations are made then this will make a lot more sense and you probably could get this in an interview.
//
//Complexities:
//
//Time: O( n )
//We are just doing linear time passes through all of this so we stay linear in time throughout.
//
//Space: O( 1 )
//We just use a few local variables, our space usage does not scale as input size gets very large.
//
//++++++++++++++++++++++++++++++++++++++++++++++++++
//
//HackerRank: https://www.youtube.com/channel/UCOf7...
//
//Tuschar Roy: https://www.youtube.com/user/tusharro...
//
//GeeksForGeeks: https://www.youtube.com/channel/UC0Rh...
//
//Jarvis Johnson: https://www.youtube.com/user/VSympathyV
//
//Success In Tech: https://www.youtube.com/channel/UC-vY...
//
//++++++++++++++++++++++++++++++++++++++++++++++++++
//
//This question on Leetcode: https://leetcode.com/problems/next-pe...
//
//This question is number 6.10 in "Elements of Programming Interviews" by Adnan Aziz, Tsung-Hsien Lee, and Amit Prakash.

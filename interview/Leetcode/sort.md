
- 多数元素，利用抵消的思想。
- AC once within 5 minutes 

```
func majorityElement(nums []int) int {
    if len(nums) == 0 {
        return 0 
    }
    curNum := nums[0]
    curCount := 1 
    for i:=1; i < len(nums); i++ {
        if curCount == 0 {
            curNum = nums[i]
            curCount++
            continue 
        }

        if nums[i] == curNum {
            curCount ++ 
        } else {
            curCount --
        }
    }
    return curNum
}
```

	
剑指 Offer 51. 数组中对逆序对

- 利用归并排序的思想，进行逆序对统计
- 利用冒泡排序的思想，进行逆序对统计

```
例子：
左边数组：2，3，5，7
右边数组：1，4，6，8
合并这两个数组时：
2 和 1 比较，1 更小，说明 1 可以和左边数组中的所有剩余数字组成逆序对，逆序对数量 +4
2 和 4 比较，2 更小，逆序对数量不增加
3 和 4 比较，3 更小，逆序对数量不增加
5 和 4 比较，4 更小，说明 4 可以和左边数组中所有剩余数字组成逆序对，逆序对数量 +2
5 和 6 比较，5 更小，逆序对数量不增加
7 和 6 比较，6 更小，说明 6 可以和左边数组中所有剩余数字组成逆序对，逆序对数量 +1
7 和 8 比较，7 更小，逆序对数量不增加

左边数组遍历完毕，将 8 添加到合并后的数组末尾，逆序对数量不增加
```



- 计数排序

```
public static void countingSort(int[] arr) {
    // 判空及防止数组越界
    if (arr == null || arr.length <= 1) return;
    // 找到最大值，最小值
    int max = arr[0];
    int min = arr[0];
    for (int i = 1; i < arr.length; i++) {
        if (arr[i] > max) max = arr[i];
        else if (arr[i] < min) min = arr[i];
    }
    // 确定计数范围
    int range = max - min + 1;
    // 建立长度为 range 的数组，下标 0~range-1 对应数字 min~max
    int[] counting = new int[range];
    // 遍历 arr 中的每个元素
    for (int element : arr) {
        // 将每个整数出现的次数统计到计数数组中对应下标的位置，这里需要将每个元素减去 min，才能映射到 0～range-1 范围内
        counting[element - min]++;
    }
    // 记录前面比自己小的数字的总数
    int preCounts = 0;
    for (int i = 0; i < range; i++) {
        // 当前的数字比下一个数字小，累计到 preCounts 中
        preCounts += counting[i];
        // 将 counting 计算成当前数字在结果中的起始下标位置。位置 = 前面比自己小的数字的总数。
        counting[i] = preCounts - counting[i];
    }
    int[] result = new int[arr.length];
    for (int element : arr) {
        // counting[element - min] 表示此元素在结果数组中的下标
        result[counting[element - min]] = element;
        // 更新 counting[element - min]，指向此元素的下一个下标
        counting[element - min]++;
    }
    // 将结果赋值回 arr
    for (int i = 0; i < arr.length; i++) {
        arr[i] = result[i];
    }
}

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/sort-algorithms/ozyo63/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
```

- 元素相对排序。
- Leetcode 1122 题
```
func relativeSortArray(arr1 []int, arr2 []int) []int {

    maxValue := 0 
    for _, i := range arr1 {
        if i > maxValue {
            maxValue = i 
        }
    }

    // 初始化 0 值，不需要 capacity
    freqCounter := make([]int, maxValue+1)

    // 计数
    for _, x := range arr1 {
        freqCounter[x]++
    }

    // 初始化最终的数组
    res := make([]int, len(arr1))
    idx := 0
    for _, i := range arr2 {
        for j := 0; j < freqCounter[i]; j++ {
            res[idx] = i 
            idx++
        }
        // 清空 freq , 但感觉是不需要的。
        // 这一步找出哪些元素不在 arr2 , 下一个循环就可以只遍历这些元素了。
        freqCounter[i] = 0 
    }

    // 不在 arr2 的元素，放在末尾升序排列。
    for i:=0; i <= maxValue; i++{
        for j :=0; j < freqCounter[i]; j++ {
            res[idx] = i
            idx++
        }
    }
    return res 
}
```



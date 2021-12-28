# 排序算法

- golang: https://github.com/asong2020/go-algorithm/tree/master/sort 

### 排序相关算法的特点

### 核心思想
### 先构造分界点，然后去左右子数组构造分界点

void sort(int[] nums, int lo, int hi) {
    /****** 前序遍历位置 ******/
    // 通过交换元素构建分界点 p
    int p = partition(nums, lo, hi);
    /************************/

    sort(nums, lo, p - 1);
    sort(nums, p + 1, hi);
}

void sort(int[] nums, int lo, int hi) {
    int mid = (lo + hi) / 2;
    sort(nums, lo, mid);
    sort(nums, mid + 1, hi);

    /****** 后序遍历位置 ******/
    // 合并两个排好序的子数组
    merge(nums, lo, mid, hi);
    /************************/
}s


### 冒泡排序

```
func sortBubble(numbers uint64Slice)  {
 length := len(numbers)
 if length == 0{
  return
 }
 flag := true

 for i:=0;i<length && flag;i++{
  flag = false
  for j:=length-1;j>i;j--{
   if numbers[j-1] > numbers[j] {
    numbers.swap(j-1,j)
    flag = true // 有数据才交换
   }
  }
 }
}

// 交换方法
func (numbers uint64Slice)swap(i,j int)  {
 numbers[i],numbers[j] = numbers[j],numbers[i]
}
```


### 快速排序


```
func quickSort(numbers uint64Slice,start,end int)  {
 var middle int
 tempStart := start
 tempEnd := end

 if tempStart >= tempEnd{
  return
 }
 pivot := numbers[start]
 for start < end{
  for start < end && numbers[end] > pivot{
   end--
  }
  if start<end{
   numbers.swap(start,end)
   start++
  }
  for start < end && numbers[start] < pivot{
   start++
  }
  if start<end{
   numbers.swap(start,end)
   end--
  }
 }
 numbers[start] = pivot
 middle = start

 quickSort(numbers,tempStart,middle-1)
 quickSort(numbers,middle+1,tempEnd)

}
```

### 插入排序

```
func insertSort(numbers uint64Slice)  {
 for i:=1; i < len(numbers); i++{
  tmp := numbers[i]
  // 从待排序序列开始比较,找到比其小的数
  j:=i
  for j>0 && tmp<numbers[j-1] {
   numbers[j] = numbers[j-1]
   j--
  }
  // 存在比其小的数插入
  if j!=i{
   numbers[j] = tmp
  }
 }
}

```

### 选择排序

```
func selectSort(numbers uint64Slice)  {
 for i := 0; i < len(numbers) - 1; i++{
  // 记录最小值位置
  min := i

  for j:= i+1; j<len(numbers);j++{
   if numbers[j] < numbers[min]{
    min = j
   }
  }
  if i != min{
   numbers.swap(i,min)
  }
 }
}
```

### 归并排序

```
func merge(left uint64Slice,right uint64Slice) uint64Slice {
 result := make(uint64Slice,0)
 for len(left) != 0 && len(right) != 0 {
  if left[0] <= right[0] {
   result = append(result, left[0])
   left = left[1:]
  } else {
   result = append(result, right[0])
   right = right[1:]
  }
 }
```

### 计算排序

```
func countSort(numbers uint64Slice,maxValue uint64) {
 bucketLen := maxValue + 1
 bucket := make(uint64Slice,bucketLen) // 初始都是0的数组
 sortedIndex := 0

 for _,v:= range numbers{
  bucket[v] +=1
 }
 var j uint64
 for j=0;j<bucketLen;j++{
  for bucket[j]>0{
   numbers[sortedIndex] = j
   sortedIndex +=1
   bucket[j] -= 1
  }
 }
}
```

### 链表排序

- 使用归并
link: https://leetcode-cn.com/problems/sort-list/solution/sort-list-gui-bing-pai-xu-lian-biao-by-jyd/ 。 



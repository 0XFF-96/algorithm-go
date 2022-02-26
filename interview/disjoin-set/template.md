

```

func countComponents(n int, edges [][]int) int {


    ans := 0 
    fa := make([]int, n)
    //MakeSet 建立并查集
    for i:=0;i<n;i++{
        fa[i] = i
    }

    for _, edge := range edges{
        x := edge[0]
        y := edge[1]

        unionSet(fa, x, y)
    }

    for i:=0;i<n;i++{

        //遇到并查集的根，答案加1
        if (Find(fa, i) == i){
            ans++
        }
    }
    return ans 

}

//并查集 路径压缩模板
func Find(fa []int, x int ) int {
    if (fa[x]==x){return  x}
     fa[x] = Find(fa, fa[x])
     return fa[x]
}  

func unionSet(fa []int,  x, y int ) {
    x, y = Find(fa, x), Find(fa, y) 
    if (x !=y){
        fa[x] = y
    }
}

```




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


```
func countComponents(n int, edges [][]int) int {
    uf := NewUF(n, edges)
    for i:=0; i<len(edges); i++ {
        edge := edges[i]
        if !uf.Connected(edge[0], edge[1]) {
            uf.Union(edge[0], edge[1])
        }
    }
    return uf.Count()
}

type UF struct {
    parent map[int]int
    size map[int]int
    count int
}

func NewUF(n int, edges [][]int) *UF {
    uf := UF{}
    uf.parent = make(map[int]int, n)
    uf.size = make(map[int]int, n)
    uf.count = n
    for i:=0; i<n; i++ {
        uf.parent[i] = i
        uf.size[i] = 1
    }
    return &uf
}

func (uf *UF) Find(a int) int {
    for a != uf.parent[a] {
        uf.parent[a] = uf.parent[uf.parent[a]]
        a = uf.parent[a]
    }
    return a
}

func (uf *UF) Connected(a, b int) bool {
    return uf.Find(a) == uf.Find(b)
}

func (uf *UF) Union(a, b int) {
    pA := uf.Find(a)
    pB := uf.Find(b)
    if uf.size[pA] > uf.size[pB] {
        uf.parent[pB] = pA
        uf.size[pA] += uf.size[pB]
    } else {
        uf.parent[pA] = pB
        uf.size[pB] += uf.size[pA]
    }
    uf.count--
}

func (uf *UF) Count() int {
    return uf.count
}

```

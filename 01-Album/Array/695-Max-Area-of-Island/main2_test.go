package main 

func maxAreaOfIslandV2(grid [][]int) int {
    var maxArea int
    weight:=len(grid[0])
    hight:=len(grid)
   
    for i:=0;i<hight;i++{
        
        for j:=0;j<weight;j++{
            
            area:=islandArea(j,i,grid,weight,hight)
            

            if area>maxArea{
                maxArea=area
            }
        }
    }

    return maxArea
    
}

func islandArea(x int,y int,grid [][]int,weight int,hight int) int{
	var area int 
	// 这种叫后置判断条件
	// 另一种，叫做前置判断条件。
    if x<0||x>=weight||y<0||y>=hight||grid[y][x]==0{
        return area
    }
    grid[y][x]=0
    area+=1
    for a:=-1;a<2;a++{
      for b:=-1;b<2;b++{

		  // 从坐标轴的思维出发
		  // 就可以知道为什么条件是 a*b 
          if a*b==0{
              area+=islandArea(x+a,y+b,grid,weight,hight)
              
          }
      
      }  
    }
    return area    
}



def main():
    inputStr1 = input("")
    inputStr2 = input("")

    arr1 = list(inputStr1.split(" "))
    arr2 = list(inputStr2.split(" "))

    # print(arr1)
    # print(arr2)
    # print("-----------")

    h1, d1, t1 = int(arr1[0]), int(arr1[1]), int(arr1[2])
    h2, d2, t2 = int(arr2[0]), int(arr2[1]), int(arr2[2])

    # print(h1, d1, t1)
    # print(h2, d2, t2)
    # print("------------")

    h1 -= d2 
    h2 -= d1 
    time = 0.5
    while h1 > 0 and h2 > 0:
        time += 0.5 
        if time % t1 == 0: 
            h2 -= d1 
        
        if time % t2 == 0:
            h1 -= d2

    if h1 <= 0 and h2 <= 0:
        print("draw")
    elif h1 <=0:
        print("player two")
    else:
        print("player one")

main()
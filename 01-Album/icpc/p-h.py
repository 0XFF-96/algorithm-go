n, k = input().split(' ')
n = int(n)
k = int(k)
sum = 0
s = str()
for i in range(1, n + 1):
    s += str(i)
    num = int(s)
    if (num % k == 0):
        sum += 1
    s = str(num % k)
print(sum)
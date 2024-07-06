n = input()
n = int(n)
for i in range(0, n):
    volume = input()
    n_volume = float(volume[:-1])
    print(hash(n_volume))
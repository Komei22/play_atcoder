# bsのlistは降順にソートされていることが前提。targetは検索する値。
def bs(list, target):
    left = 0
    right = len(list)-1

    # leftとrightが同値のときは、探索を終えたことを意味する。
    while left <= right:
        middle = (left+right)//2
        val = list[middle]
        if val == target:
            return middle
        # 検索する値targetが現在の中央の値valより小さい時、右側には解は存在しないので、rightを更新。逆に、targetの方が大きい場合は左側には解は存在しないのでleftを更新。
        if val > target:
            right = middle - 1
        else:
            left = middle + 1
    return None


numbers = []
for i in range(10000):
    numbers.append(i)
idx = bs(numbers, 7)
print(idx)

# すごろく問題
# スタートからゴールまでがnマス目の場合、ゴールぴったりに上がる上がり方は何通りか？

# 与えられたnからサイコロのでた目を引いていく。最終的に0になったらそのパターンはゴールぴったり。マイナス値になったらゴールを通り過ぎる。

import sys

cache = {}


def deep_first_search(num):
    global cache

    # 0のときは、現在見ているサイコロの出目の組み合わせでちょうどゴールにたどり着いたことを意味する
    if num == 0:
        return 1
    # 0以下のときは、現在のサイコロの目ではちょうどたどり着けなかったことを意味する
    if num < 0:
        return 0
    # これまでに計算したパターンが出現したら、キャッシュから呼び出す
    if cache.get(num) is not None:
        return cache[num]

    result = 0
    for i in range(1, 6):
        # 与えられた数字に対して、出た目を引いた値を再帰的に呼び出す
        result += deep_first_search(num-i)

    cache[num] = result
    return result


if __name__ == "__main__":
    sys.setrecursionlimit(20000)
    N = int(sys.argv[1])
    pattern_num = deep_first_search(N)
    print(pattern_num)

# アルファベット A, B, C からなる、3文字以下の文字列の組み合わせをすべて取得する。
# stackで実装すると、木の右側から探索する


def is_empty(stack):
    return True if stack else False


def bfs(patterns):
    # スタックの初期状態の定義。スタックがノードに対応。
    stack = [""]
    while is_empty(stack):
        # スタックから要素を取り出す
        string = stack.pop()

        if len(string) > 3:
            continue

        patterns.append(string)
        for rune in ['A', 'B', 'C']:
            stack.append(string+rune)


if __name__ == "__main__":
    patterns = []
    bfs(patterns)
    print(patterns)

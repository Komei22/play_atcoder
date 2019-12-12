# アルファベット A, B, C からなる、3文字以下の文字列の組み合わせをすべて取得する。
# stackで実装すると、木の右側から探索する

def bfs(patterns)
  # スタックとなる配列を用意し、初期状態である空文字列を格納する
  stack = [""]
  until stack.empty?
    # スタックから1つ取り出す（末尾から取り出すのでpopを使う）
    str = stack.pop

    # 停止条件: 文字列調査3文字より大きい
    next if str.length > 3

    # 値の格納する
    patterns << str unless str.empty?

    # 次のノードへの遷移処理
    ['A', 'B', 'C'].each do |rune|
      stack << str+rune
    end
  end
end

patterns = [] # 結果を格納する配列を用意しておく
bfs(patterns)
p patterns # 組み合わせの表示

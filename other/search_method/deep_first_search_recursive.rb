# アルファベット A, B, C からなる、3文字以下の文字列の組み合わせをすべて取得する。
# 再気による実装では、木の左側から探索を行う

def dfs(patterns, str)
  # 停止条件
  if str.length > 3
    return
  end

  # 値格納
  unless str.empty?
    patterns.push(str)
  end

  # 次のノードへ移動
  ['A','B','C'].each do |rune|
    dfs(patterns, str+rune)
  end

  return
end

patterns = []
dfs(patterns, "")
p(patterns)
printf("pattern length: %d", patterns.length)

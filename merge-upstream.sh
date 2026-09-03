#!/bin/bash
set -e

UPSTREAM_BRANCH="${1:-main}"

echo "正在從 upstream2 合併 $UPSTREAM_BRANCH..."
git merge "upstream2/$UPSTREAM_BRANCH" --no-commit || true

echo "正在自動將上游路徑替換為你的路徑..."
find . -type f \( -name "*.go" -o -name "go.mod" \) -not -path './.git*' -exec sed -i '' 's|github.com/nuomiiiii/lite|github.com/raymao96/komari|g' {} +

git add .
echo "合併與路徑對齊已完成！請輸入 git commit 完成合併。"

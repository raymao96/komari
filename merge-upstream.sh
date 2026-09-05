#!/bin/bash
set -e

UPSTREAM_BRANCH="${1:-upstream2/main}"

echo "正在從 $UPSTREAM_BRANCH 進行完整合併（採用上游優先策略）..."
git merge "$UPSTREAM_BRANCH" -X theirs --no-commit || true

echo "正在全面將整個倉庫內的所有檔案與路徑徹底替換..."

python3 -c '
import os

old_str = "github.com/nuomiiiii/lite"
new_str = "github.com/raymao96/komari"
old_short = "nuomiiiii/lite"
new_short = "raymao96/komari"

for root, dirs, files in os.walk("."):
    if ".git" in dirs:
        dirs.remove(".git")
    if "merge-upstream.sh" in files:
        files.remove("merge-upstream.sh")
        
    for file in files:
        filepath = os.path.join(root, file)
        try:
            with open(filepath, "r", encoding="utf-8", errors="ignore") as f:
                content = f.read()
            
            if old_str in content or old_short in content:
                new_content = content.replace(old_str, new_str).replace(old_short, new_short)
                with open(filepath, "w", encoding="utf-8") as f:
                    f.write(new_content)
        except Exception:
            pass
'

# 將所有變更加入暫存區
git add -A

echo "合併與全域路徑替換完成！請執行 git status 檢查變更狀況。"

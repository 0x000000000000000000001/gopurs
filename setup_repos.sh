#!/usr/bin/env bash
set -e

HTDOCS="/Users/0x1/Documents/htdocs"
cd "$HTDOCS"

for dir in "$HTDOCS"/phpurs-*; do
  if [ ! -d "$dir" ]; then continue; fi
  
  basename=$(basename "$dir")
  suffix=${basename#phpurs-}
  
  gopurs_dir="$HTDOCS/gopurs-$suffix"
  
  if [ -d "$gopurs_dir" ]; then
    echo "Directory $gopurs_dir already exists. Skipping."
    continue
  fi
  
  echo "=> Processing $suffix"
  
  upstream=""
  if [ -f "$dir/bower.json" ]; then
    upstream=$(grep -o '"url": *"[^"]*"' "$dir/bower.json" | cut -d'"' -f4 | head -n 1)
  fi
  
  if [ -z "$upstream" ]; then
    upstream=$(git -C "$dir" remote get-url upstream 2>/dev/null || echo "")
  fi
  
  if [ -z "$upstream" ]; then
    echo "Could not determine upstream for $suffix. Skipping."
    continue
  fi
  
  echo "Found upstream: $upstream"
  
  mkdir "$gopurs_dir"
  cd "$gopurs_dir"
  git init
  
  git remote add upstream "$upstream"
  git fetch upstream
  
  # Figure out default branch
  branch="master"
  if git ls-remote --exit-code --heads upstream main >/dev/null 2>&1; then
    branch="main"
  fi
  
  git checkout -b "$branch" "upstream/$branch"
  
  # Copy configuration files
  if [ -f "$dir/spago.yaml" ]; then
    sed 's/phpurs/gopurs/g' "$dir/spago.yaml" > spago.yaml
    git add spago.yaml
    git commit -m "chore: add spago.yaml for gopurs" || true
  fi
  
  if [ -f "$dir/bower.json" ]; then
    sed 's/phpurs/gopurs/g' "$dir/bower.json" > bower.json
    git add bower.json
    git commit -m "chore: update bower.json for gopurs" || true
  fi
  
  unset GITHUB_TOKEN
  echo "Creating GitHub repo 0x000000000000000000001/gopurs-$suffix"
  if gh repo create "0x000000000000000000001/gopurs-$suffix" --public --source=. --remote=origin --push; then
     echo "Success for $suffix"
  else
     echo "Failed to create repo for $suffix on GitHub"
  fi
  
  echo "-----------------"
  
  sleep 3
done

echo "All repos processed!"

#!/bin/bash

# 转换Windows路径到WSL路径
BACKEND_PATH="/mnt/d/Ubuntu_Home/bubblevan/asian_games/application/server"
FRONTEND_PATH="/mnt/d/Ubuntu_Home/bubblevan/asian_games/application/web"

# 启动后端
cd "$BACKEND_PATH"
nohup go run main.go > backend.log 2>&1 &

# 启动前端
cd "$FRONTEND_PATH"
nohup npm run dev > frontend.log 2>&1 &

echo "前后端应用已启动"
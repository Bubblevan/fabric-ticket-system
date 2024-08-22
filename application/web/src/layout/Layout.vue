<template>
  <el-container style="height: 100vh;">
    <el-aside width="200px" class="sidebar">
      <el-scrollbar>
        <el-menu :default-openeds="['1', '2', '3', '4']" class="el-menu-vertical-demo" @select="handleSelect">
          <el-menu-item index="/main">
            <el-icon><home-filled /></el-icon>
            <span slot="title">首页</span>
          </el-menu-item>
          <el-sub-menu index="1">
            <template #title>
              <el-icon><calendar /></el-icon>比赛管理
            </template>
            <el-menu-item-group title="">
              <el-menu-item index="/events">比赛列表</el-menu-item>
              <el-menu-item index="/events/create">创建比赛</el-menu-item>
              <el-menu-item index="/blockchain-explorer">浏览器</el-menu-item>
            </el-menu-item-group>
          </el-sub-menu>
          <el-sub-menu index="2">
            <template #title>
              <el-icon><user /></el-icon>用户管理
            </template>
            <el-menu-item-group title="">
              <el-menu-item index="/users">用户列表</el-menu-item>
            </el-menu-item-group>
          </el-sub-menu>
          <el-sub-menu index="3">
            <template #title>
              <el-icon><setting /></el-icon>个人中心
            </template>
            <el-menu-item-group title="">
              <el-menu-item index="/profile">个人信息</el-menu-item>
              <el-menu-item index="/orders">我的购买</el-menu-item>
            </el-menu-item-group>
          </el-sub-menu>
        </el-menu>
      </el-scrollbar>
    </el-aside>

    <el-container>
      <el-header class="header">
        <h1 class="title">杭州第19届亚运会</h1>
        <el-button type="danger" @click="logout">注销</el-button>
      </el-header>
      <el-main class="content">
        <slot></slot> <!-- 用于渲染子组件内容 -->
        <Breadcrumb />
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import { HomeFilled, Calendar, Ticket, User, Setting } from '@element-plus/icons-vue';
import Breadcrumb from '@/components/Breadcrumb.vue';

export default {
  components: {
    Breadcrumb,
    HomeFilled,
    Calendar,
    Ticket,
    User,
    Setting
  },
  methods: {
    handleSelect(key, keyPath) {
      this.$router.push(key);
    },
    logout() {
      // 注销逻辑
    }
  }
};
</script>

<style scoped>
@font-face {
  font-family: 'YeZiGongChangZhaoPaiTi-2';
  src: url('../assets/YeZiGongChangZhaoPaiTi-2.ttf') format('truetype');
}
.title {
  font-family: 'YeZiGongChangZhaoPaiTi-2';
  font-size: 2.5em;
  font-weight: bold;
  background: linear-gradient(to right, #f4eff9, #ffffff);
  -webkit-background-clip: text; /* 将背景剪裁为文字 */
  -webkit-text-fill-color: transparent; /* 将文字填充为透明 */
}
.el-container {
  height: 100vh;
}

.sidebar {
  background-color: #f0f2f5;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
  background-color: #1890ff;
  color: white;
}

.content {
  padding: 20px;
}

.project-name {
  padding: 20px;
  font-size: 18px;
  font-weight: bold;
}

.el-menu-vertical-demo {
  border-right: none;
}

.el-menu-item {
  padding-left: 20px !important;
}
</style>

<template>
    <Layout>
      <div class="profile-container">
        <h1>个人信息</h1>
        <el-form :model="user" label-width="120px">
          <el-form-item label="用户名">
            <el-input v-model="user.username" disabled></el-input>
          </el-form-item>
          <el-form-item label="电子邮件">
            <el-input v-model="user.email" disabled></el-input>
          </el-form-item>
          <el-form-item label="全名">
            <el-input v-model="user.fullName" disabled></el-input>
          </el-form-item>
          <el-form-item label="角色">
            <el-input v-model="user.roleText" disabled></el-input>
          </el-form-item>
        </el-form>
      </div>
    </Layout>
  </template>
  
  <script>
  import Layout from '../layout/Layout.vue';
  import axios from 'axios';
  
  export default {
  components: {
    Layout
  },
  data() {
    return {
      user: {
        username: '',
        email: '',
        fullName: '',
        roleText: ''
      }
    };
  },
  created() {
    this.fetchProfile();
  },
  methods: {
  async fetchProfile() {
    try {
      const token = sessionStorage.getItem('token');
      if (token) {
        console.log('Profile Token:', token); // 输出 sessionStorage 中的 token
      } else {
        console.log('No token found in sessionStorage');
      }

      const response = await axios.get('http://localhost:8080/profile', {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });
      console.log('Response:', response.data);
      if (response.data && response.data.data) {
        const userData = response.data.data;
        this.user.username = userData.Username;
        this.user.email = userData.Email;
        this.user.fullName = userData.FullName;
        this.user.roleText = userData.Role ? '亚运主办方' : '游客';
      } else {
        this.$message.error('获取个人信息失败: 数据格式不正确');
      }
    } catch (error) {
      console.error('Error fetching profile:', error);
      this.$message.error('获取个人信息失败');
    }
  }
}
};
  </script>
  
  <style scoped>
  .profile-container {
    max-width: 600px;
    margin: 50px auto;
    padding: 20px;
    background: #fff;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    border-radius: 8px;
  }
  
  h1 {
    text-align: center;
    margin-bottom: 20px;
  }
  
  .el-form-item {
    margin-bottom: 20px;
  }
  </style>
  
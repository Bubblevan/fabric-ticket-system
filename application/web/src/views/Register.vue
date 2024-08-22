<template>
    <div class="register-container">
      <h1 class="title">注册</h1>
      <form @submit.prevent="register" class="register-form">
        <div class="form-group">
          <label for="newUsername">用户名</label>
          <input type="text" id="newUsername" v-model="newUsername" placeholder="请输入用户名" class="form-control" />
        </div>
        <div class="form-group">
          <label for="newPassword">密码</label>
          <input type="password" id="newPassword" v-model="newPassword" placeholder="请输入密码" class="form-control" />
        </div>
        <div class="form-group">
          <label for="newEmail">电子邮件</label>
          <input type="email" id="newEmail" v-model="newEmail" placeholder="请输入电子邮件" class="form-control" />
        </div>
        <div class="form-group">
          <label>角色</label>
          <div>
            <input type="radio" id="visitor" value="false" v-model="role" />
            <label for="visitor">游客</label>
            <input type="radio" id="organizer" value="true" v-model="role" />
            <label for="organizer">亚运主办方</label>
          </div>
        </div>
        <button type="submit" class="btn btn-primary">注册</button>
        <button type="button" @click="goToLogin" class="btn btn-secondary">返回登录</button>
      </form>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue';
  import axios from 'axios';
  import { useRouter } from 'vue-router';
  
  export default defineComponent({
    data() {
      return {
        newUsername: '',
        newPassword: '',
        newEmail: '',
        role: 'false' // 默认角色为游客
      };
    },
    setup() {
      const router = useRouter();
      return { router };
    },
    methods: {
      async register() {
        try {
          const response = await axios.post('http://localhost:8080/register', {
            username: this.newUsername,
            password: this.newPassword,
            email: this.newEmail,
            role: this.role === 'true'  // 发送角色字段
          });
          if (response.data.code === 0) { // 成功
            alert(response.data.msg);
            this.router.push('/login');
          } else { // 失败
            alert(response.data.msg);
          }
        } catch (error: any) {
          alert('注册失败');
        }
      },
      goToLogin() {
        this.router.push('/');
      }
    }
  });
  </script>
  
  
  <style scoped>
  .register-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    background-color: #f0f2f5;
  }
  
  .title {
    margin-bottom: 20px;
    color: #333;
  }
  
  .register-form {
    width: 300px;
    padding: 20px;
    background-color: #fff;
    border-radius: 8px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  }
  
  .form-group {
    margin-bottom: 15px;
  }
  
  .form-group label {
    display: block;
    margin-bottom: 5px;
    color: #666;
  }
  
  .form-control {
    width: 100%;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 4px;
  }
  
  .btn {
    width: 100%;
    padding: 10px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 16px;
  }
  
  .btn-primary {
    background-color: #1890ff;
    color: #fff;
    margin-bottom: 10px;
  }
  
  .btn-secondary {
    background-color: #f0f2f5;
    color: #333;
  }
  </style>
  
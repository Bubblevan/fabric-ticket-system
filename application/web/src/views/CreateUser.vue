<template>
  <Layout>
    <div class="create-user-container">
      <h1 class="title">创建用户</h1>
      <form @submit.prevent="createUser" class="create-user-form">
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
        <button type="submit" class="btn btn-primary">创建</button>
        <button type="button" @click="goToUsers" class="btn btn-secondary">返回用户管理</button>
      </form>
    </div>
  </Layout>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';
import Layout from '../layout/Layout.vue'; // 确保你已经导入了Layout组件

export default defineComponent({
  components: {
    Layout
  },
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
    async createUser() {
      try {
        const response = await axios.post('http://localhost:8080/register', {
          username: this.newUsername,
          password: this.newPassword,
          email: this.newEmail,
          role: this.role === 'true'  // 发送角色字段
        });
        if (response.data.code === 0) { // 成功
          alert(response.data.msg);
          this.router.push('/users');
        } else { // 失败
          alert(response.data.msg);
        }
      } catch (error: any) {
        alert('创建用户失败');
      }
    },
    goToUsers() {
      this.router.push('/users');
    }
  }
});
</script>

<style scoped>
.create-user-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.title {
  text-align: center;
}

.form-group {
  margin-bottom: 15px;
}

.form-control {
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
}

.btn {
  padding: 10px 15px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.btn-primary {
  background-color: #007bff;
  color: white;
}

.btn-secondary {
  background-color: #6c757d;
  color: white;
  margin-left: 10px;
}
</style>
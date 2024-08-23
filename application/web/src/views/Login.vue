<template>
  <div class="login-container">
    <h1 class="title">登录</h1>
    <form @submit.prevent="login" class="login-form">
      <div class="form-group" v-if="loginBy === 'username'">
        <label for="username">用户名</label>
        <input type="text" id="username" v-model="username" placeholder="请输入用户名" class="form-control" />
      </div>
      <div class="form-group" v-if="loginBy === 'email'">
        <label for="email">邮箱</label>
        <input type="email" id="email" v-model="email" placeholder="请输入邮箱" class="form-control" />
      </div>
      <div class="form-group">
        <label for="password">密码</label>
        <input type="password" id="password" v-model="password" placeholder="请输入密码" class="form-control" />
      </div>
      <div class="form-group">
        <label>角色</label>
        <div>
          <input type="radio" id="visitor" :value="false" v-model="role" />
          <label for="visitor">游客</label>
          <input type="radio" id="organizer" :value="true" v-model="role" />
          <label for="organizer">亚运主办方</label>
        </div>
      </div>
      <button type="submit" class="btn btn-primary">登录</button>
    </form>
    <p class="register-link" @click="goToRegister">没有账号？点击注册</p>
    <p class="login-switch" @click="toggleLoginBy">
      {{ loginBy === 'username' ? '通过邮箱登录' : '通过用户名登录' }}
    </p>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

// 导入图片
import bg6 from '@/assets/bg-6.png';

export default defineComponent({
  data() {
    return {
      username: '',
      email: '',
      password: '',
      role: false, // 默认角色为游客
      loginBy: 'username', // 默认通过用户名登录
      backgroundImages: [
        bg6
      ],
      currentImageIndex: 0
    };
  },
  setup() {
    const router = useRouter();
    return { router };
  },
  mounted() {
    this.startBackgroundRotation();
  },
  methods: {
    async login() {
      try {
        const loginData = this.loginBy === 'username'
          ? { username: this.username, password: this.password, role: this.role }
          : { email: this.email, password: this.password, role: this.role };

        const response = await axios.post('http://localhost:8080/login', loginData);
        alert(response.data.msg); // 修改为 msg
        if (response.data.code === 0) { // 确保检查code
          const token = response.data.data.token;
          sessionStorage.setItem('token', token); // 确保存储token
          sessionStorage.setItem('role', this.role ? 'organizer' : 'visitor'); // 确保存储角色
          sessionStorage.setItem('userID', response.data.userID); // 确保存储userID
          console.log(token, 'vue login的token');
          this.router.push('/main');
        } else {
          alert(response.data.msg);
        }
      } catch (error) {
        alert((error as any).response.data.msg);
      }
    },
    goToRegister() {
      this.router.push('/register');
    },
    toggleLoginBy() {
      this.loginBy = this.loginBy === 'username' ? 'email' : 'username';
    },
    startBackgroundRotation() {
      setInterval(() => {
        this.currentImageIndex = (this.currentImageIndex + 1) % this.backgroundImages.length;
        document.body.style.backgroundImage = `url(${this.backgroundImages[this.currentImageIndex]})`;
      }, 5000);
    }
  }
});
</script>

<style scoped>
.login-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center;
  transition: background-image 1s ease-in-out;
  background-blend-mode: overlay; 
  animation: floatBackground 20s linear infinite; /* 应用动画 */
}

@keyframes floatBackground {
  0% {
    background-position: 50% 100%;
  }
  100% {
    background-position: 50% -100%;
  }
}

.title {
  margin-bottom: 20px;
  color: #333;
}

.login-form {
  width: 300px;
  padding: 20px;
  background-color: rgba(255, 255, 255, 0.9);
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

.register-link, .login-switch {
  color: #1890ff;
  cursor: pointer;
  margin-top: 10px;
}
</style>
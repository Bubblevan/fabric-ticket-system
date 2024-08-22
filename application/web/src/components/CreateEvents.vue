<template>
  <Layout>
    <div class="form-container">
      <h1>创建比赛</h1>
      <el-form :model="ticket" @submit.prevent="createTicket" label-width="120px">
        <el-form-item label="活动名称">
          <el-input v-model="ticket.eventName"></el-input>
        </el-form-item>
        <el-form-item label="场地">
          <el-input v-model="ticket.venue"></el-input>
        </el-form-item>
        <el-form-item label="活动日期">
          <el-date-picker v-model="ticket.eventDate" type="date"></el-date-picker>
        </el-form-item>
        <el-form-item label="价格">
          <el-input v-model="ticket.price" type="number"></el-input>
        </el-form-item>
        <el-form-item label="数量">
          <el-input v-model="ticket.num" type="number"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="createTicket">创建</el-button>
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
      ticket: {
        eventName: '',
        venue: '',
        eventDate: '',
        price: 0,
        num: 0
      },
      token: ''
    };
  },
  mounted() {
    this.token = sessionStorage.getItem('token');
    console.log('Token in mounted:', this.token);
  },
  methods: {
    async createTicket() {
      try {
        const token = this.token;
        console.log('Token in createTicket:', token);

        const response = await axios.post('http://localhost:8080/createTicket', {
          eventName: this.ticket.eventName,
          venue: this.ticket.venue,
          eventDate: this.ticket.eventDate,
          price: parseFloat(this.ticket.price), // 确保是数字
          num: parseInt(this.ticket.num) // 确保是数字
        }, {
          headers: {
            'Authorization': `Bearer ${token}` // 添加JWT token到请求头
          }
        });

        console.log('Response:', response);

        if (response.data.code === 0) {
          this.$message.success('创建门票成功');
          this.$router.push('/events');
        } else {
          this.$message.error(response.data.msg);
        }
      } catch (error) {
        console.error('Error:', error);
        this.$message.error('创建门票失败');
      }
    }
  }
};
</script>
  
  <style scoped>
  .form-container {
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
  
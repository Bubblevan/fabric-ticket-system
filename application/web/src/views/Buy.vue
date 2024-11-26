<template>
    <Layout>
      <div class="buy-container">
        <h1>购买门票</h1>
        <el-form :model="ticket" label-width="120px">
          <el-form-item label="活动名称">
            <el-input v-model="ticket.EventName" disabled></el-input>
          </el-form-item>
          <el-form-item label="场地">
            <el-input v-model="ticket.Venue" disabled></el-input>
          </el-form-item>
          <el-form-item label="活动日期">
            <el-input v-model="formattedDate" disabled></el-input>
          </el-form-item>
          <el-form-item label="价格">
            <el-input v-model="ticket.Price" disabled></el-input>
          </el-form-item>
          <el-form-item label="剩余票数">
            <el-input v-model="ticket.Num" disabled></el-input>
          </el-form-item>
          <el-form-item label="购买数量">
            <el-input v-model="quantity"></el-input>
          </el-form-item>
        </el-form>
        <el-button type="primary" @click="buyTicket">购买</el-button>
      </div>
    </Layout>
  </template>
  
  <script>
  import Layout from '../layout/Layout.vue';
  import axios from 'axios';
  import { format } from 'date-fns';
  import { v4 as uuidv4 } from 'uuid'; // 引入UUID库
  
  export default {
    components: {
      Layout
    },
    data() {
      return {
        ticket: {},
        quantity: 1,
        formattedDate: '',
        userID: null // 存储用户ID
      };
    },
    async created() {
      await this.fetchUserProfile();
      this.fetchTicket();
    },
    methods: {
      async fetchUserProfile() {
        try {
          const token = sessionStorage.getItem('token');
          const response = await axios.get('http://localhost:8080/profile', {
            headers: {
              'Authorization': `Bearer ${token}`
            }
          });
          this.userID = response.data.data.ID;
          console.log('User ID:', this.userID); // 调试信息，确保用户ID正确获取
        } catch (error) {
          this.$message.error('获取个人信息失败');
          console.error('Error fetching user profile:', error);
        }
      },
      async fetchTicket() {
        const id = this.$route.params.id;
        try {
          const response = await axios.get(`http://localhost:8080/tickets/${id}`);
          this.ticket = response.data.data;
          this.formattedDate = format(new Date(this.ticket.EventDate), 'yyyy-MM-dd');
          console.log(this.ticket); // 调试信息，确保数据正确获取
        } catch (error) {
          this.$message.error('获取门票详情失败');
        }
      },
      async buyTicket() {
        if (this.quantity > this.ticket.Num) {
          this.$message.error('购买数量超过剩余票数');
          return;
        }
  
        const order = {
          id: uuidv4(), // 生成唯一的orderID
          userID: this.userID, // 使用获取到的用户ID
          ticketID: this.ticket.ID, // 使用正确的字段名
          num: this.quantity, // 注意这里的字段名应该是 "num" 而不是 "quantity"
          orderDate: new Date().toISOString() // 转换为ISO 8601格式
        };
  
        try {
          const token = sessionStorage.getItem('token');
          console.log('Sending order:', order); // 调试信息，确保订单数据正确
          await axios.post('http://localhost:8080/createOrder', order, {
            headers: {
              'Authorization': `Bearer ${token}`,
              'Content-Type': 'application/json'
            }
          });
          this.$message.success('购买成功');
          this.$router.push('/events');
        } catch (error) {
          this.$message.error('购买失败');
          console.error('Error creating order:', error); // 调试信息，查看错误详情
        }
      }
    }
  };
</script>
  
  <style scoped>
  .buy-container {
    padding: 20px;
  }
  </style>
  
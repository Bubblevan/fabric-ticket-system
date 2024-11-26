<template>
  <Layout>
    <template v-slot:default>
      <h1>购买历史</h1>
      <el-table :data="orders" class="custom-table">
        <!-- 根据实际数据结构添加列 -->
        <el-table-column prop="Key" label="订单ID"></el-table-column>
        <el-table-column prop="Record.userID" label="用户ID"></el-table-column>
        <el-table-column prop="Record.ticketID" label="票ID"></el-table-column>
        <el-table-column prop="Record.num" label="票数"></el-table-column>
        <el-table-column prop="Record.totalPrice" label="总价"></el-table-column>
        <el-table-column prop="Record.orderDate" label="订单日期"></el-table-column>
      </el-table>
    </template>
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
      orders: []
    };
  },
  async created() {
    await this.fetchUserProfile();
    this.fetchOrders();
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
        const userID = response.data.data.ID;
        sessionStorage.setItem('userID', userID);
      } catch (error) {
        this.$message.error('获取个人信息失败');
        console.error('Error fetching user profile:', error);
      }
    },
    async fetchOrders() {
      try {
        const token = sessionStorage.getItem('token');
        const userID = sessionStorage.getItem('userID');

        // 调试信息，确保 userID 存在且格式正确
        console.log('userID:', userID);

        const response = await axios.get('http://localhost:8080/queryOrdersByUserID', {
          params: {
            userID: userID
          },
          headers: {
            'Authorization': `Bearer ${token}`
          }
        });

        // 解析后端返回的JSON字符串
        const ordersData = JSON.parse(response.data.result);
        this.orders = ordersData;
      } catch (error) {
        this.$message.error('获取购买历史失败');
        console.error('Error fetching orders:', error);
      }
    }
  }
};
</script>
  <style scoped>
  .purchase-container {
    max-width: 800px;
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
  
  .el-table {
    width: 100%;
  }
  </style>
  
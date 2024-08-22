<template>
  <Layout>
    <template v-slot:default>
      <h1>比赛列表</h1>
      <el-table :data="tickets" class="custom-table">
        <el-table-column prop="EventName" label="活动名称"></el-table-column>
        <el-table-column prop="Venue" label="场地"></el-table-column>
        <el-table-column label="活动日期" :formatter="formatDate"></el-table-column>
        <el-table-column prop="Price" label="价格"></el-table-column>
        <el-table-column prop="Num" label="票数"></el-table-column>
        <el-table-column label="操作">
          <template v-slot="scope">
            <el-button type="primary" @click="goToBuy(scope.row)">购买</el-button>
          </template>
        </el-table-column>
      </el-table>
    </template>
  </Layout>
</template>

<script>
import Layout from '../layout/Layout.vue';
import axios from 'axios';
import { format } from 'date-fns';

export default {
  components: {
    Layout
  },
  data() {
    return {
      tickets: []
    };
  },
  created() {
    this.fetchTickets();
  },
  methods: {
    async fetchTickets() {
      try {
        const response = await axios.get('http://localhost:8080/tickets');
        // 进行字段映射
        this.tickets = response.data.data.map(ticket => ({
          ...ticket,
          id: ticket.ID, // 将 ID 映射为 id
          EventDate: ticket.EventDate // 确保 EventDate 字段一致
        }));
        console.log(this.tickets); // 调试信息，确保数据正确获取
      } catch (error) {
        this.$message.error('获取门票列表失败');
      }
    },
    formatDate(row) {
      const cellValue = row.EventDate; // 确保 EventDate 是正确的属性名
      if (cellValue) {
        return format(new Date(cellValue), 'yyyy-MM-dd');
      }
      return '';
    },
    goToBuy(ticket) {
      if (!ticket.id) {
        this.$message.error('票的ID缺失');
        return;
      }
      this.$router.push({ name: 'Buy', params: { id: ticket.id } });
    }
  }
};
</script>

<style scoped>
.custom-table {
  width: 100%;
  margin: 0 auto;
  border-collapse: collapse;
}

.custom-table th,
.custom-table td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

.custom-table th {
  background-color: #f2f2f2;
}

.custom-table tr:nth-child(even) {
  background-color: #f9f9f9;
}

.custom-table tr:hover {
  background-color: #f5f5f5;
}
</style>

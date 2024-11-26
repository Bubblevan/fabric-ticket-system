<template>
  <Layout>
    <div class="users-container">
      <h1>用户管理</h1>
      <el-input
        v-model="searchQuery"
        placeholder="输入用户名进行搜索"
        @keyup.enter="fetchUsers"
      />
      <el-button type="primary" @click="fetchUsers">搜索</el-button>
      <el-button type="success" @click="goToCreateUser">添加用户</el-button>
      <el-table :data="users" style="width: 100%">
        <el-table-column prop="Username" label="用户名" width="180" />
        <el-table-column prop="Email" label="电子邮件" />
        <el-table-column prop="FullName" label="全名" />
        <el-table-column prop="Role" label="角色">
          <template v-slot="scope">
            {{ scope.row.Role ? '主办方' : '游客' }}
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template v-slot="scope">
            <el-button
              type="primary"
              size="small"
              @click="editUser(scope.row)"
            >编辑</el-button>
            <el-button
              type="danger"
              size="small"
              @click="deleteUser(scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog :visible.sync="editUserDialogVisible" title="编辑用户">
      <el-form :model="editedUser">
        <el-form-item label="用户名">
          <el-input v-model="editedUser.username" />
        </el-form-item>
        <el-form-item label="电子邮件">
          <el-input v-model="editedUser.email" />
        </el-form-item>
        <el-form-item label="全名">
          <el-input v-model="editedUser.fullName" />
        </el-form-item>
        <el-form-item label="角色">
          <el-select v-model="editedUser.role">
            <el-option label="游客" :value="false" />
            <el-option label="主办方" :value="true" />
          </el-select>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="editUserDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveUser">保存</el-button>
      </span>
    </el-dialog>
  </Layout>
</template>

<script>
import axios from 'axios';
import Layout from '../layout/Layout.vue';

export default {
  components: {
    Layout
  },
  data() {
    return {
      users: [],
      searchQuery: '',
      newUser: {
        username: '',
        email: '',
        fullName: '',
        role: false,
      },
      editUserDialogVisible: false, 
      editedUser: {
        username: '',
        email: '',
        fullName: '',
        role: false,
      },
    };
  },
  methods: {
    fetchUsers() {
      axios.get('http://localhost:8080/users')
        .then((response) => {
          console.log('Response from API:', response.data); // 添加调试信息
          this.users = response.data.data;
        })
        .catch((error) => {
          console.error('Error fetching users:', error); // 添加调试信息
        });
    },
    goToCreateUser() {
    this.$router.push({ path: '/users/create' });
    },
    showAddUserDialog() {
      this.addUserDialogVisible = true;
    },
    addUser() {
      axios.post('http://localhost:8080/users', this.newUser)
        .then(() => {
          this.addUserDialogVisible = false;
          this.fetchUsers();
        })
        .catch((error) => {
          console.error('Error adding user:', error); // 添加调试信息
        });
    },
    editUser(user) {
      this.editedUser = { ...user };
      this.editUserDialogVisible = true;
    },
    saveUser() {
      axios.put(`http://localhost:8080/users/${this.editedUser.ID}`, this.editedUser)
        .then(() => {
          this.editUserDialogVisible = false;
          this.fetchUsers();
        })
        .catch((error) => {
          console.error('Error updating user:', error); // 添加调试信息
        });
    },
    deleteUser(user) {
      axios.delete(`http://localhost:8080/users/${user.ID}`)
        .then(() => {
          this.fetchUsers();
        })
        .catch((error) => {
          console.error('Error deleting user:', error); // 添加调试信息
        });
    },
  },
  created() {
    this.fetchUsers();
  },
};
</script>

<style>
.users-container {
  padding: 20px;
}
</style>
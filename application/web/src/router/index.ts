import { createRouter, createWebHistory } from 'vue-router';
import Login from '../views/Login.vue';
import Register from '../views/Register.vue';
import Main from '../views/Main.vue';
import Events from '../views/Events.vue';
import CreateEvents from '../components/CreateEvents.vue';
import Profile from '../views/Profile.vue';
import Purchased from '../views/Purchase.vue'
import Users from '../views/Users.vue'
import Buy from '../views/Buy.vue'
import CreateUser from '../views/CreateUser.vue';
import BlockchainExplorer from '@/components/BlockchainExplorer.vue';
import component from 'element-plus/es/components/tree-select/src/tree-select-option.mjs';

const routes = [
  { path: '/', component: Login },
  { path: '/register', component: Register },
  { path: '/main', component: Main },
  { path: '/events', name: 'Events', component: Events },
  { path: '/events/create', name: 'CreateEvents', component: CreateEvents },
  { path: '/events/buy/:id', name: 'Buy', component: Buy, props: true }, // 确保这个路径在 /events 之后
  { path: '/users', name: 'Users', component: Users },
  { path: '/users/create', name: 'CreateUser', component: CreateUser },
  { path: '/profile', name: 'Profile', component: Profile },
  { path: '/orders', name: 'Purchased', component: Purchased }, // 起到区块链浏览器的作用
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;

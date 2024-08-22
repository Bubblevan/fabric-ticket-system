<template>
    <div class="breadcrumb">
      <router-link v-for="(breadcrumb, index) in breadcrumbs" :key="index" :to="breadcrumb.path">
        {{ breadcrumb.name }}
        <span v-if="index < breadcrumbs.length - 1"> / </span>
      </router-link>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent, computed } from 'vue';
  import { useRoute } from 'vue-router';
  
  export default defineComponent({
    name: 'Breadcrumb',
    setup() {
      const route = useRoute();
  
      const breadcrumbs = computed(() => {
        const paths = route.path.split('/').filter(path => path);
        return paths.map((path, index) => {
          return {
            name: path.charAt(0).toUpperCase() + path.slice(1), // 将路径的首字母大写作为名称
            path: '/' + paths.slice(0, index + 1).join('/')
          };
        });
      });
  
      return { breadcrumbs };
    }
  });
  </script>
  
  <style scoped>
  .breadcrumb {
    padding: 10px 0;
    font-size: 14px;
  }
  
  .breadcrumb a {
    color: #1890ff;
    text-decoration: none;
  }
  
  .breadcrumb a:hover {
    text-decoration: underline;
  }
  
  .breadcrumb span {
    margin: 0 5px;
  }
  </style>
  
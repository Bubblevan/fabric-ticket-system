<template>
    <div>
      <h1>Blockchain Explorer</h1>
      <iframe
        src="http://localhost:8081/#/"
        width="100%"
        height="800px"
        frameborder="0"
      ></iframe>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue';
  
  export default defineComponent({
    name: 'BlockchainExplorer',
  });
  </script>
  
  <style scoped>
  iframe {
    border: none;
  }
  </style>
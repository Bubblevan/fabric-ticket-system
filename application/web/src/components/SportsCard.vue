<template>
    <div class="sports-card">
      <img :src="image" :alt="nameEn" class="sports-image" />
      <div class="sports-info">
        <h3 class="sports-name">{{ name }}</h3>
        <p class="sports-name-en">{{ nameEn }}</p>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'SportsCard',
    props: {
      image: {
        type: String,
        required: true
      },
      name: {
        type: String,
        required: true
      },
      nameEn: {
        type: String,
        required: true
      }
    }
  };
  </script>
  
  <style scoped>
  .sports-card {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 15px;
    border: 1px solid #ddd;
    border-radius: 10px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    transition: transform 0.2s ease;
    cursor: pointer;
  }
  
  .sports-card:hover {
    transform: translateY(-5px);
  }
  
  .sports-image {
    width: 100px;
    height: 100px;
    object-fit: cover;
    border-radius: 50%;
    margin-bottom: 10px;
  }
  
  .sports-info {
    text-align: center;
  }
  
  .sports-name {
    font-size: 1.5em;
    font-weight: bold;
    margin: 0;
  }
  
  .sports-name-en {
    font-size: 1em;
    color: #888;
    margin: 0;
  }
  </style>
  
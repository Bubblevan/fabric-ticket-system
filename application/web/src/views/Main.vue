<template>
  <Layout>
    <div class="header-image"></div>
    <div class="content">
      <div class="news-section">
        <div class="news-header">
          <span>赛场快讯</span>
          <button @click="showAllNews">全部</button>
        </div>
        <div class="news-list">
          <div v-for="(news, index) in displayedNews" :key="index" class="news-item">
            <div class="news-line">
              <strong class="news-title">{{ news.title }}</strong>
              <span class="news-content">{{ news.content }}</span>
              <span class="news-date">{{ news.date }}</span>
            </div>
          </div>
        </div>
      </div>
      <div class="sports-section">
        <div class="sports-header">
          <span>亚运项目</span>
          <button @click="showAllGames">全部</button>
        </div>
        <div class="sports-grid">
          <SportsCard
            v-for="(sport, index) in sports"
            :key="index"
            :image="sport.image"
            :name="sport.name"
            :nameEn="sport.nameEn"
          />
        </div>
      </div>
    </div>
  </Layout>
</template>

<script>
import Layout from '../layout/Layout.vue';
import newsData from '../assets/news.json'; // 假设你的news.json文件在assets目录下
import SportsCard from '../components/SportsCard.vue';

import athleticsImage from '@/assets/athletics.png';
import swimmingImage from '@/assets/swimming.png';
import tableTennisImage from '@/assets/table-tennis.png';
import esportsImage from '@/assets/esports.png';
import badmintonImage from '@/assets/badminton.png';
import weightliftingImage from '@/assets/weightlifting.png';

export default {
  components: {
    Layout,
    SportsCard
  },
  data() {
    return {
      sports: [
        { name: '田径', nameEn: 'Athletics', image: athleticsImage },
        { name: '游泳', nameEn: 'Swimming', image: swimmingImage },
        { name: '乒乓球', nameEn: 'Table Tennis', image: tableTennisImage },
        { name: '电子竞技', nameEn: 'Esports', image: esportsImage },
        { name: '羽毛球', nameEn: 'Badminton', image: badmintonImage },
        { name: '举重', nameEn: 'Weightlifting', image: weightliftingImage },
      ],
      allNews: [],
      displayedNews: []
    };
  },
  mounted() {
    this.loadNews();
  },
  methods: {
    loadNews() {
      // 读取json文件内容并解析
      this.allNews = newsData;
      this.randomizeNews();
    },
    randomizeNews() {
      // 随机选择5条新闻
      this.displayedNews = this.allNews.sort(() => 0.5 - Math.random()).slice(0, 5);
    },
    showAllNews() {
      // 按照时间顺序展示所有新闻
      this.displayedNews = this.allNews.sort((a, b) => {
        const dateA = new Date(a.date);
        const dateB = new Date(b.date);
        return dateA - dateB;
      });
    },
    showAllGames() {
      this.$router.push({ name: 'Events' });
    }
  }
};
</script>

<style scoped>
@font-face {
  font-family: 'YeZiGongChangZhaoPaiTi-2';
  src: url('../assets/YeZiGongChangZhaoPaiTi-2.ttf') format('truetype');
}
.header-image {
  width: 100%;
  height: 500px; /* 根据需要调整高度 */
  background-image: url('../assets/header.png');
  background-size: cover;
  background-position: center;
}

.content {
  padding: 20px;
}

.news-section {
  margin-top: 20px;
}

.news-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
  padding: 10px;
  background-color: #f5f5f5;
  border-radius: 5px;
}

.news-header span {
  font-family: 'YeZiGongChangZhaoPaiTi-2';
  font-size: 3em;
  font-weight: bold;
  background: linear-gradient(to right, #8e2de2, #4a00e0); /* 紫色渐变背景 */
  -webkit-background-clip: text; /* 将背景剪裁为文字 */
  -webkit-text-fill-color: transparent; /* 将文字填充为透明 */
}

.news-header button {
  padding: 8px 16px;
  font-size: 1em;
  color: #fff;
  background-color: #007bff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.news-header button:hover {
  background-color: #0056b3;
}

.news-list {
  border: 1px solid #ccc;
  padding: 10px;
}

.news-item {
  margin-bottom: 20px; /* 增加间距 */
  padding-bottom: 20px; /* 增加间距 */
  border-bottom: 1px solid #ddd; /* 淡横线 */
}

.news-item:last-child {
  border-bottom: none; /* 去掉最后一条新闻的横线 */
  padding-bottom: 0; /* 去掉最后一条新闻的间距 */
}

.news-line {
  display: flex;
  justify-content: space-between;
  align-items: baseline;
}

.news-title {
  font-size: 1.5em;
  font-weight: bold;
  margin-right: 10px;
}

.news-content {
  font-size: 1.3em;
  flex-grow: 1;
  margin-right: 10px;
  text-align: left;
}

.news-date {
  text-align: right;
  white-space: nowrap;
}

.sports-grid {
  display: flex;
  flex-wrap: wrap; /* 如果有太多的卡片，换行显示 */
  justify-content: space-around; /* 卡片之间留出空间，或者用 'space-between' 或 'center' */
  gap: 20px; /* 设置卡片之间的间距 */
}
.sports-section {
  margin-top: 20px;
}

.sports-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
  padding: 10px;
  background-color: #f5f5f5;
  border-radius: 5px;
}

.sports-header span {
  font-family: 'YeZiGongChangZhaoPaiTi-2';
  font-size: 3em;
  font-weight: bold;
  background: linear-gradient(to right, #8e2de2, #4a00e0); /* 紫色渐变背景 */
  -webkit-background-clip: text; /* 将背景剪裁为文字 */
  -webkit-text-fill-color: transparent; /* 将文字填充为透明 */
}
.sports-header button {
  padding: 8px 16px;
  font-size: 1em;
  color: #fff;
  background-color: #007bff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}
</style>
<template>
  <div class="ranking">
    <TheHeader />
    <main class="main-content">
      <div class="container">
        <h1>排行榜</h1>
        <!-- 后续添加排行榜内容 -->
      </div>
    </main>
  </div>
</template>

<script>
import TheHeader from '@/components/layout/TheHeader.vue'

export default {
  name: 'RankingPage',
  components: {
    TheHeader
  }
}
</script>

<style scoped>
.ranking {
  min-height: 100vh;
}

.main-content {
  padding-top: 64px;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

h1 {
  color: #333;
  text-align: center;
  margin: 2rem 0;
}
</style>
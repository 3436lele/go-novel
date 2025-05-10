<template>
  <div class="home">
    <TheHeader />
    <main class="main-content">
      <div class="container">
        <h1>欢迎来到兔缘书界网</h1>
        <!-- 这里后续可以添加更多首页内容 -->
      </div>
    </main>
  </div>
</template>

<script>
import TheHeader from '@/components/layout/TheHeader.vue'

export default {
  name: 'HomePage',
  components: {
    TheHeader
  }
}
</script>

<style scoped>
.home {
  min-height: 100vh;
}

.main-content {
  padding-top: 64px; /* 与 header 高度相同 */
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
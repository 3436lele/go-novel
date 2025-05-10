<template>
  <div class="novels">
    <TheHeader />
    <main class="main-content">
      <div class="container">
        <h1>小说列表</h1>
        <!-- 后续添加小说列表内容 -->
      </div>
    </main>
  </div>
</template>

<script>
import TheHeader from '@/components/layout/TheHeader.vue'

export default {
  name: 'NovelsPage',
  components: {
    TheHeader
  }
}
</script>

<style scoped>
.novels {
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
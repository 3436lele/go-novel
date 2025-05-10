<template>
  <header class="header">
    <div class="container">
      <div class="header-left">
        <router-link to="/" class="logo">
          <img src="@/assets/logo.png" alt="兔缘书界网" class="logo-img">
          <span class="logo-text">兔缘书界网</span>
        </router-link>
      </div>

      <nav class="nav-menu">
        <router-link to="/" class="nav-item">首页</router-link>
        <router-link to="/novels" class="nav-item">小说</router-link>
        <router-link to="/worldview" class="nav-item">世界观</router-link>
        <router-link to="/ranking" class="nav-item">排行榜</router-link>
      </nav>

      <div class="header-right">
        <div class="search-box">
          <el-input
            v-model="searchQuery"
            placeholder="搜索小说..."
            class="search-input"
          >
            <template #prefix>
              <el-icon><Search /></el-icon>
            </template>
          </el-input>
        </div>

        <div class="auth-buttons">
          <el-button type="primary" plain>登录</el-button>
          <el-button type="primary">注册</el-button>
        </div>
      </div>

      <!-- 移动端菜单按钮 -->
      <el-button class="mobile-menu-btn" @click="toggleMobileMenu">
        <el-icon><Menu /></el-icon>
      </el-button>
    </div>

    <!-- 移动端菜单 -->
    <div class="mobile-menu" :class="{ 'is-active': isMobileMenuOpen }">
      <nav class="mobile-nav">
        <router-link to="/" class="mobile-nav-item" @click="closeMobileMenu">首页</router-link>
        <router-link to="/novels" class="mobile-nav-item" @click="closeMobileMenu">小说</router-link>
        <router-link to="/worldview" class="mobile-nav-item" @click="closeMobileMenu">世界观</router-link>
        <router-link to="/ranking" class="mobile-nav-item" @click="closeMobileMenu">排行榜</router-link>
      </nav>
      <div class="mobile-auth">
        <el-button type="primary" plain block>登录</el-button>
        <el-button type="primary" block>注册</el-button>
      </div>
    </div>
  </header>
</template>

<script>
import { ref } from 'vue'
import { Search, Menu } from '@element-plus/icons-vue'

export default {
  name: 'TheHeader',
  components: {
    Search,
    Menu
  },
  setup() {
    const searchQuery = ref('')
    const isMobileMenuOpen = ref(false)

    const toggleMobileMenu = () => {
      isMobileMenuOpen.value = !isMobileMenuOpen.value
    }

    const closeMobileMenu = () => {
      isMobileMenuOpen.value = false
    }

    return {
      searchQuery,
      isMobileMenuOpen,
      toggleMobileMenu,
      closeMobileMenu
    }
  }
}
</script>

<style scoped>
.header {
  background-color: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
  height: 64px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.header-left {
  display: flex;
  align-items: center;
}

.logo {
  display: flex;
  align-items: center;
  text-decoration: none;
  color: #2c3e50;
}

.logo-img {
  height: 48px;
  margin-right: 12px;
  border-radius: 8px;
}

.logo-text {
  font-size: 2.2rem;
  font-weight: bold;
  color: #2c3e50;
  text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.1);
  letter-spacing: 1px;
}

.nav-menu {
  display: flex;
  gap: 2rem;
}

.nav-item {
  text-decoration: none;
  color: #34495e;
  font-size: 1rem;
  padding: 0.5rem 0;
  position: relative;
  transition: color 0.3s;
}

.nav-item:hover {
  color: #409EFF;
}

.nav-item::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 0;
  height: 2px;
  background-color: #409EFF;
  transition: width 0.3s;
}

.nav-item:hover::after {
  width: 100%;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

.search-box {
  width: 240px;
}

.search-input :deep(.el-input__wrapper) {
  border-radius: 20px;
}

.search-input :deep(.el-input__inner) {
  color: #34495e;
}

.search-input :deep(.el-input__inner::placeholder) {
  color: #95a5a6;
}

.auth-buttons {
  display: flex;
  gap: 1rem;
}

.auth-buttons .el-button {
  border-radius: 20px;
  padding: 8px 20px;
}

.auth-buttons .el-button--primary {
  color: #fff;
}

.auth-buttons .el-button--primary.is-plain {
  color: #409EFF;
}

/* 移动端菜单按钮 */
.mobile-menu-btn {
  display: none;
  font-size: 24px;
  padding: 8px;
  border: none;
  background: none;
  color: #34495e;
  cursor: pointer;
}

/* 移动端菜单 */
.mobile-menu {
  display: none;
  position: fixed;
  top: 64px;
  left: 0;
  right: 0;
  background: #fff;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  transform: translateY(-100%);
  transition: transform 0.3s ease;
}

.mobile-menu.is-active {
  transform: translateY(0);
}

.mobile-nav {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  margin-bottom: 1rem;
}

.mobile-nav-item {
  text-decoration: none;
  color: #34495e;
  font-size: 1.1rem;
  padding: 0.5rem 0;
  border-bottom: 1px solid #eee;
}

.mobile-auth {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

/* 响应式布局 */
@media screen and (max-width: 1024px) {
  .container {
    padding: 0 15px;
  }

  .logo-text {
    font-size: 1.8rem;
  }

  .nav-menu {
    gap: 1.5rem;
  }

  .search-box {
    width: 200px;
  }
}

@media screen and (max-width: 768px) {
  .nav-menu,
  .header-right {
    display: none;
  }

  .mobile-menu-btn {
    display: block;
  }

  .mobile-menu {
    display: block;
  }

  .logo-text {
    font-size: 1.5rem;
  }

  .logo-img {
    height: 40px;
  }
}

@media screen and (max-width: 480px) {
  .logo-text {
    font-size: 1.2rem;
  }

  .logo-img {
    height: 36px;
    margin-right: 8px;
  }

  .container {
    padding: 0 10px;
  }
}
</style> 
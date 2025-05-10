import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '@/views/Home.vue'

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomePage
    },
    {
      path: '/novels',
      name: 'novels',
      component: () => import('@/views/Novels.vue')
    },
    {
      path: '/worldview',
      name: 'worldview',
      component: () => import('@/views/Worldview.vue')
    },
    {
      path: '/ranking',
      name: 'ranking',
      component: () => import('@/views/Ranking.vue')
    }
  ]
})

export default router 
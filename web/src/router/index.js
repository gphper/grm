import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'home',
    component: () => import('../views/HomeView.vue'),
    children:[{
      path: 'monitor/:sk',
      name: 'monitor',
      component: () => import('../views/MonitorView.vue')
    }]
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../views/LoginView.vue'),
  },
]

const router = createRouter({
  history: createWebHashHistory(process.env.BASE_URL),
  routes
})


router.beforeEach((to, from, next) => {
  if (to.name !== 'login' && sessionStorage.getItem("auth") == null){
    next({ name: 'login' })
  }else{
    next()
  }
})

export default router

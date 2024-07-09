import {createRouter, createWebHistory} from 'vue-router';
import home from "./components/app/home.vue";
import LoginPage from "./components/app/admin/auth/login-page.vue";
import {tokenNotExpired} from "./components/app/admin/auth/auth.jwt.ts";

const routes = [
    {
        path: '/login',
        name: 'login',
        component: LoginPage,
    },
    {
        path: '/home',
        name: 'home',
        component: home,
    },
    // 其他路由配置
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

// 添加全局路由守卫
router.beforeEach((to, _, next) => {
    const isAuthenticated = tokenNotExpired();
    if (to.path == '/' && !isAuthenticated) next({ name: 'login' })
    else if (to.path == '/' && isAuthenticated) next({ name: 'home' })
    else if (to.name !== 'login' && !isAuthenticated) next({ name: 'login' })
    else next()
})

export default router;
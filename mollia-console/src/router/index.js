import {createRouter, createWebHistory } from 'vue-router';
import { useProfileStore } from '../stores/ProfileStore';

const routes = [
    {
        path: '/login',
        name: 'Login',
        component: () => import('../views/auth/Login.vue'),
    },
    {
        path: '/',
        component: () => import('../views/home/Home.vue')
    },
    {
        path: '/article/create',
        component: () => import('../views/article/ArticleEditor.vue')
    },
    {
        path: '/article/:id',
        component: () => import('../views/article/ArticleEditor.vue')
    },
    {
        path: '/article/list',
        component: () => import('../views/article/ArticleList.vue')
    },
    {
        path: '/article/category',
        component: () => import('../views/article/ArticleCategory.vue')
    },
    {
        path: '/article/tags',
        component: () => import('../views/article/TagManagement.vue')
    },
    {
        path: '/option/theme',
        component: () => import('../views/option/ThemeOption.vue'),
    },
    {
        path: '/option/profile',
        component: () => import('../views/option/ProfileOption.vue'),
    },
    {
        path: '/option/about',
        component: () => import('../views/option/AboutOption.vue'),
    },
    {
        path: '/storage/image',
        component: () => import('../views/storage/ImageStorage.vue'),
    },
    {
        path: '/storage/attach',
        component: () => import('../views/storage/AttachStorage.vue')
    }
];

const router = new createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    const profileStore = useProfileStore();
    const isAuthenticated = !!profileStore.token;

    if (to.name !== 'Login' && !isAuthenticated) {
        next({ name: 'Login' });
    } else if (to.name === 'Login' && isAuthenticated) {
        next('/'); // 如果已登录，访问登录页则重定向到首页
    } else {
        next();
    }
});

export default router;
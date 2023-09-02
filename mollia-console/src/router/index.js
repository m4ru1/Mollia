import {createRouter, createWebHistory } from 'vue-router';

const router = new createRouter({
    history: createWebHistory(),
    routes: [
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
    ]
});

export default router;
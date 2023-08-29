import {createRouter, createWebHistory } from 'vue-router';

const router = new createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            component: () => import('../views/home/Home.vue'),
            children: [
                {
                    path: 'article',
                    component: () => import('../views/notfound/NotFound.vue'),
                    children: [
                        {
                            path: '/create',
                            component: () => import('../views/article/ArticleEditor.vue')
                        },
                        {
                            path: '/:id',
                            component: () => import('../views/article/ArticleEditor.vue')
                        },
                        {
                            path: '/list',
                            component: () => import('../views/article/ArticleCategory.vue')
                        },
                        {
                            path: '/category',
                            component: () => import('../views/article/ArticleCategory.vue')
                        }
                    ]
                },
                {
                    path: 'option',
                    component: () => import('../views/notfound/NotFound.vue'),
                    children: [
                        {
                            path: '/theme',
                            component: () => import('../views/option/ThemeOption.vue'),
                        },
                        {
                            path: '/profile',
                            component: () => import('../views/option/ProfileOption.vue'),
                        },
                        {
                            path: '/about',
                            component: () => import('../views/option/AboutOption.vue'),
                        },
                    ]
                },
                {
                    path: 'storage',
                    component: () => import('../views/notfound/NotFound.vue'),
                    children: [
                        {
                            path: '/image',
                            component: () => import('../views/storage/ImageStorage.vue'),
                        },
                        {
                            path: '/attach',
                            component: () => import('../views/storage/AttachStorage.vue')
                        }
                    ]
                }
            ],
        }
    ]
});

export default router;
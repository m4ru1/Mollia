import { createRouter, createWebHistory } from 'vue-router'

const router = new createRouter({
    history: createWebHistory(),
    routes: [
        { path: "/", component: () => import("../views/home/HomePage.vue")},
        { path: "/article", component: () => import("../views/article/ArticlePage.vue")},
        { path: "/tag", component: () => import("../views/tag/TagPage.vue")},
        { path: "/about", component: () => import("../views/about/AboutPage.vue")},
        { path: "/notfound", component: () => import("../views/notfound/NotFound.vue")}
    ]
})

export default router
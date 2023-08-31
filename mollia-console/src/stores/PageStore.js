import { defineStore } from "pinia";

export const usePageStore = defineStore('page-store', {
    state: () => {
        return {
            pages: [{
                path: '/',
                curr: true,
                title: undefined
            }]
        }
    },
    getters: {
        routeMap() {
            return {
                '/': {pageName: '首页', pagePath: ['首页']},
                '/article/create': {pageName: '发布文章', pagePath: ['文章', '发布文章']},
                '/article/list': {pageName: '文章列表', pagePath: ['文章', '文章列表']},
                '/article/category': {pageName: '文章分类', pagePath: ['文章', '文章分类']},
                '/storage/image': {pageName: '图床管理', pagePath: ['存储', '图床管理']},
                '/storage/attach': {pageName: '附件管理', pagePath: ['存储', '附件管理']},
                '/option/profile': {pageName: '网站信息', pagePath: ['设置', '网站信息']},
                '/option/theme': {pageName: '主题设置', pagePath: ['设置', '主题设置']},
                '/option/about': {pageName: '关于页面', pagePath: ['设置', '关于页面']},
            }
        },

        // 返回当前被选中的页面对象
        currPage: (state) => state.pages.find((item) => item.curr),

        // 返回当前被选中的页面对象的路径，用于面包屑展示
        currPagePath() {
            // title不为undefined的页面对象, 视为文章编辑页面
            if(this.currPage.title){
                return ['文章', '编辑文章'];
            }else{
                return this.routeMap[this.currPage.path].pagePath;
            }
        },

        // 返回给定页面对象的页面名称, title用来处理特定文章的编辑页面, 当title存在时将作为页面名称显示在HeaderTag上
        // 调用方式: state.pageName(pageObject);
        pageName(){
            return (item) => {
                if(item.title){
                    return item.title;
                }else{
                    return this.routeMap[item.path].pageName;
                }
            };
        } 
    }
});
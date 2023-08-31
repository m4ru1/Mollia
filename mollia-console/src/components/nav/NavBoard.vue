<template>
    <div class="nav-board">
        <el-menu
            default-active="1"
            class="nav-content"
            :collapse="themeStore.isCollapse"
            popper-effect="dark"
        >
            <div class="nav-header"> 
                <span v-if="!themeStore.isCollapse"> Mollia </span>
                <span v-else> Mo </span>
            </div>
            <el-menu-item index="1" @click="handleClick">
                <template #title>
                    <el-icon><House /></el-icon>
                    <span> 仪表盘 </span>
                </template>
            </el-menu-item>
            <el-sub-menu index="2" :show-timeout="100" :hide-timeout="10">
                <template #title>
                    <el-icon><document /></el-icon>
                    <span> 文章 </span>
                </template>
                <el-menu-item index="2-1" @click="handleClick">
                    <el-icon><edit /></el-icon>
                    <span> 发布文章 </span>
                </el-menu-item>
                <el-menu-item index="2-2" @click="handleClick">
                    <el-icon><Menu /></el-icon>
                    <span> 文章列表 </span>
                </el-menu-item>
                <el-menu-item index="2-3" @click="handleClick">
                    <el-icon><price-tag /></el-icon>
                    <span> 文章分类 </span>
                </el-menu-item>
            </el-sub-menu>
            <el-sub-menu index="3" :show-timeout="100" :hide-timeout="10">
                <template #title>
                    <el-icon><Files /></el-icon>
                    <span> 存储 </span>
                </template>
                <el-menu-item index="3-1" @click="handleClick">
                    <el-icon><PictureFilled /></el-icon>
                    <span> 图床管理 </span>
                </el-menu-item>
                <el-menu-item index="3-2" @click="handleClick">
                    <el-icon><Management /></el-icon>
                    <span> 附件管理 </span>
                </el-menu-item>            
            </el-sub-menu>
            <el-sub-menu index="4" :show-timeout="100" :hide-timeout="10">
                <template #title>
                    <el-icon><Operation /></el-icon>
                    <span> 设置 </span>
                </template>
                <el-menu-item index="4-1" @click="handleClick">
                    <el-icon><Avatar /></el-icon>
                    <span> 网站信息 </span>
                </el-menu-item>
                <el-menu-item index="4-2" @click="handleClick">
                    <el-icon><Sunrise /></el-icon>
                    <span> 主题设置 </span>
                </el-menu-item>  
                <el-menu-item index="4-3" @click="handleClick">
                    <el-icon><Notification /></el-icon>
                    <span> 关于页面 </span>
                </el-menu-item>
            </el-sub-menu>
        </el-menu>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { usePageStore } from '../../stores/PageStore';
import { useThemeStore } from '../../stores/ThemeStore';
import { useRouter } from 'vue-router';
import { House, Document, Edit, PriceTag, Menu,
         Files, PictureFilled, Management, Operation,
         Avatar, Sunrise, Notification} 
         from '@element-plus/icons-vue';

let themeStore = useThemeStore();
let pageStore = usePageStore();

const router = useRouter();
const routeMap = {
    '1': '/',
    '2-1': '/article/create',
    '2-2': '/article/list',
    '2-3': '/article/category',
    '3-1': '/storage/image',
    '3-2': '/storage/attach',
    '4-1': '/option/profile',
    '4-2': '/option/theme',
    '4-3': '/option/about',
}

function handleClick(elMenuItem){

    // 1. 获取路由路径
    let routePath = routeMap[elMenuItem.index];
    // 2. 路由变更
    if(routePath){
        router.push(routePath);
    }
    // 3. page-store内容更新
    // 检查pages数组是否存在选中页面，如果存在则将其标记为当前页面.
    // 若不存在，则将其添加到pages中，再重复上一步骤.
    pageStore.currPage.curr = false;

    let selectedPage = pageStore.pages.find((item) => item.path === routePath);
    if (selectedPage){
        selectedPage.curr = true;
    }else{
        pageStore.pages.push({
            path: routePath,
            curr: true
        });
    }
}
</script>

<style>

.nav-board{
    /* min-width: 250px; */
    transition: all .2s ease-in-out;
    --el-menu-bg-color: #333;
    --el-menu-text-color: white;
    --el-menu-hover-bg-color: rgba(16, 185, 129, .4);
    --el-transition-duration: .18s;
    user-select: none;
    .nav-header{
        padding: 30px 0px;
        box-shadow: 0px 1px 1px #10b981;
        margin-bottom: 15px;
        & > span{
            display: block;
            font-size: 35px;
            font-weight: 900;
            text-align: center;
            color: #10b981;
        }
    }
    .nav-content{
        height: 100vh;
        border-right: 0px;
    }
}
</style>
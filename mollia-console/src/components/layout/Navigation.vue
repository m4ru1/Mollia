<template>
    <div class="navigation">
        <el-icon class="collapse-btn">
            <expand v-if="themeStore.isCollapse" @click="handleExpand" />
            <fold v-else @click="handleCollapse" /> 
        </el-icon>
        <el-breadcrumb :separator-icon="ArrowRight" class="">
            <el-breadcrumb-item class="breadcrumb-item" v-for="(item, index) in pageStore.currPagePath" :key="index">
                {{ item }}
            </el-breadcrumb-item>
        </el-breadcrumb>
        <div class="avatar-area">

        </div>
    </div>
</template>

<script setup>
import { usePageStore } from '../../../stores/PageStore';
import {
    Expand, Fold
} from '@element-plus/icons-vue';
import { useThemeStore } from '../../../stores/ThemeStore';

let themeStore = useThemeStore();
let pageStore = usePageStore();

function handleCollapse(){
    themeStore.isCollapse = true;
}

function handleExpand(){
    themeStore.isCollapse = false;
}
</script>

<style scoped>
.navigation{
    display: flex;
    flex-direction: row;
    .collapse-btn{
        font-size: 21px; 
        margin: 2px 10px 2px 10px;
        cursor: pointer;
        transition: .5s all ease-in;
    }
    .breadcrumb-item{
        padding: 6px 0;
    }
}
</style>
<template>
<div class="page-dock">
    <div class="dock-content">
        <PageTag v-for="(item, index) in pageStore.pages" :key="index"
                :page-name="pageStore.pageName(item)">
        </PageTag>
    </div>
</div>
</template>

<script setup>
import { onMounted } from 'vue';
import { usePageStore } from '../../../stores/PageStore';
import PageTag from './components/PageTag.vue';

let pageStore = usePageStore();

onMounted(() => {
    addEventListener('wheel', (event) => {
        let dock = document.querySelector('.page-dock .dock-content');
        if (dock){
            dock.setAttribute('translateX', '5px');
        }else{
            console.log('undefined dock-content');
        }
    });    
})
</script>

<style scoped>
.page-dock{
    --el-text-color-regular: black;
    padding: 5px 30px;
    .dock-content{
        display: flex;
        overflow: hidden;
    }
}
</style>
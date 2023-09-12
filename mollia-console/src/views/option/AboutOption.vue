<template>
    <div class="about-option">
        <PageContent class="page-content">
            <template #page-title>
                <h2>关于页面编辑</h2>
            </template>
            <template #page-view>
                <div id="editor"></div>
                <div class="opt-board">
                    <el-button type="primary" style="width: 50px;" @click="postHandle">
                        发布
                    </el-button>
                    <el-button type="primary" style="width: 50px;" @click="resetHandle">
                        重置
                    </el-button>
                </div>
            </template>
        </PageContent>
    </div>
</template>

<script setup>
import Cherry from 'cherry-markdown';
import PageContent from '../../components/layout/PageContent.vue';
import { onMounted } from 'vue';
import 'cherry-markdown/dist/cherry-markdown.css';

// 创建编辑器实例
const config = {
    id: 'editor',
    value: '',
    callback: {
        afterchange: function(md, html){
            console.log('change');
        }
    }
};


let cherry;

onMounted(() => {
    cherry = new Cherry(config); 
});

// 发布逻辑
function postHandle(){

}

// 重置逻辑
function resetHandle(){

}
</script>

<style scoped>
.about-option{
    display: flex;
    overflow: hidden;
    position: relative;
    background-color: var(--theme-pagecontent-bg-color);
    .page-content{
        display: flex;
        flex-direction: column;
        border-radius: 20px;
        #editor{
            max-width: 2600px;
            max-height: 70vh;
            padding: 15px;
        }
        .opt-board{
            display: flex;
            justify-content: end;
            padding: 5px;
            margin-right: 20px;
            margin-bottom: 10px;
        }
    }
}
</style>
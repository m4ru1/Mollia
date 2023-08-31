<template>
<div :class="{'page-tag': true, 'page-tag-selected': pageStore.pageName(pageStore.currPage) === props.pageName}">
    <span @click="handleClick"> {{ props.pageName }} </span>
    <div class="close-btn" @click="handleClose" v-show="pageStore.pages.length > 1">
        <svg t="1693471106027" class="close-icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4005" xmlns:xlink="http://www.w3.org/1999/xlink" width="200" height="200"><path d="M801.645714 170.666667l51.833905 51.590095L565.150476 511.951238l288.353524 289.670095-51.833905 51.614477-288.109714-289.450667L225.426286 853.23581 173.592381 801.621333l288.329143-289.670095L173.592381 222.256762 225.426286 170.666667l288.109714 289.426285L801.645714 170.666667z" p-id="4006"></path></svg>
    </div>
</div>
</template>

<script setup>
import { usePageStore } from '../../../../stores/PageStore';
import { useRouter } from 'vue-router';

const props = defineProps({
    pageName: String,
})

let router = useRouter();
let pageStore = usePageStore();

function handleClick(){

    let pageObject = pageStore.pages.find((item) => pageStore.pageName(item) === props.pageName);
    if(pageObject){
        
        // 切换标签页
        pageStore.currPage.curr = false;
        pageObject.curr = true;
        
        // 路由变更
        router.push(pageObject.path);
    }
}

function handleClose(){
    // 仅当剩余打开标签页数量大于1时生效
    if(pageStore.pages.length > 1){
        // 标签页关闭
        let selectedPage;
        let pageIndex = pageStore.pages.findIndex((item) => pageStore.pageName(item) === props.pageName);
        if (pageIndex != -1){
            let pageObject = pageStore.pages[pageIndex];
            if (pageObject.curr){ // 若被删除页面为当前被选中页面，则应选择已打开页面中的一项作为下一个被选中页面
                if (pageIndex === (pageStore.pages.length - 1)){
                    selectedPage = pageStore.pages[pageIndex - 1];
                }else if (pageIndex === 0){
                    selectedPage = pageStore.pages[1];
                }else{
                    selectedPage = pageStore.pages[pageIndex + 1];
                }
            }
            pageStore.pages.splice(pageIndex, 1);
            selectedPage.curr = true;
        }
        // 拓展：文章缓存数据丢弃
        // TODO: 
    }
}
</script>

<style scoped>

.page-tag, .page-tag-selected{
    /* 变量区 */
    --pagetag-bg-color: white;

    /* ***** */
    min-width: 84px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    border: 1px solid #ccc;
    border-radius: 8px;
    padding: 0px 5px;
    margin: 0 0 0 15px;
    background-color: var(--pagetag-bg-color);
    font-size: 14px;
    & > span{
        padding: 10px 15px 10px 15px;
        text-align: center;
        cursor: pointer;
    }
    .close-btn{
        margin-top: 2px;
        margin-left: 10px;
        transition: all .5s ease-in-out;
        .close-icon{
            height: 16px;
            width: 16px;
            border-radius: 5px;
            border: 1px solid rgba(255, 255, 255, 0);
            transition: all .1s ease-in-out;
        }
        &:hover .close-icon{
            fill: #10b981;
            border-color: #ccc;
        }
        &:active .close-icon{
            fill: white;
            background-color: #bbb;
        }
    }
}

.page-tag-selected{
    background-color: #eee;
    border-color: var(--theme-primary-color);
    box-shadow: 0px 1px 3px var(--theme-primary-color);
}
</style>
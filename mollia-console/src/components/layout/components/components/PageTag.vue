<template>
<div :class="{'page-tag': true, 'page-tag-selected': pageStore.pageName(pageStore.currPage) === props.pageName}">
    <span @click="handleClick"> {{ props.pageName }} </span>
    <div class="close-btn" @click="handleClose">
        <svg t="1693471106027" class="close-icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4005" xmlns:xlink="http://www.w3.org/1999/xlink" width="200" height="200"><path d="M801.645714 170.666667l51.833905 51.590095L565.150476 511.951238l288.353524 289.670095-51.833905 51.614477-288.109714-289.450667L225.426286 853.23581 173.592381 801.621333l288.329143-289.670095L173.592381 222.256762 225.426286 170.666667l288.109714 289.426285L801.645714 170.666667z" p-id="4006"></path></svg>
    </div>
</div>
</template>

<script setup>
import { usePageStore } from '../../../../stores/PageStore';

const props = defineProps({
    pageName: String,
})

let pageStore = usePageStore();

function handleClick(){
    // 切换标签页

    // 路由变更

}

function handleClose(){
    // 标签页关闭

    // 拓展：文章缓存数据丢弃

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
    padding: 5px 15px;
    margin: 0 0 0 15px;
    background-color: var(--pagetag-bg-color);
    font-size: 14px;
    & > span{
        cursor: pointer;
        text-align: center;
    }
    .close-btn{
        transition: all .5s ease-in-out;
        .close-icon{
            height: 16px;
            width: 16px;
            margin-top: 2px;
            margin-left: 10px;
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
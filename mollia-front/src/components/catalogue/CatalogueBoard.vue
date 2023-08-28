<template>
    <div class="catalogue-board">
        <div class="catalogue-header">
            <p>目录</p>
        </div>
        <div class="catalogue-content">
            <li :class="['li' + '-' + item.type, {'curr-tag': currTag === item.name}]" v-for="(item, index) in props.catalogue" :key="index">
                <a :class="'a' + '-' + item.type" @click="scrollTo(item.name)">
                    {{ item.name }}
                </a>
            </li>
        </div>
    </div>
</template>

<script setup>
import {ref, onMounted, onUnmounted} from "vue";

const props = defineProps({
    catalogue: Array,
})

let currTag = ref('');  // 当前锚点
let shiftLock = true;   // 锁：避免两个切换锚点事件同时执行，造成页面紊乱。
                        //     在一个切换任务未执行完成时，无法触发新的切换锚点任务。 

onMounted(() => {
    addEventListener('scroll', onScroll);
}),

onUnmounted(() => {
    /* 容器卸载时，注销监听器 */
    removeEventListener('scroll', onScroll);
});

/* 目录逻辑 */
/* 目录标签的高亮与定位跟随页面展示而变化，例如当前看到标题一，那么目录中定位到标题一且予以高亮*/
// 思路借鉴：https://juejin.cn/post/6844903940539023367
// - scrollTop: 一个元素的 scrollTop 值是这个元素的顶部到视口可见内容（的顶部）的距离的度量。
//              当一个元素的内容没有产生垂直方向的滚动条，那么它的 scrollTop 值为0。
// - offsetTop: HTMLElement.offsetTop 为只读属性，它返回当前元素相对于其 offsetParent 元素的顶部内边距的距离。
// - offsetParent: 与当前元素最近的经过定位( position 不等于 static )的父级元素

function onScroll(){

    // 获取所有锚点共同offsetParent的scrollTop
    // 注：离SpecificArticlePage组件中的文章锚点最近的定位元素，是ArticleArea Div>
    // 但是main和article area都没有产生滚动条，scrollTop为0，因此直接选用文档流的scrollTop
    let documentScrollTop = document.documentElement.scrollTop || document.body.scrollTop;
    // 获取当前所有锚点元素的offsetTop TODO: 目录栏也应当随着选中项的变化而跳转到指定tag上
    // 考虑到锚点元素的offsetParent是<div .article-area>，因此与documentScrollTop比较时需加上article-area的offsetTop
    // 为了避免完全没过anchor后对应tag才被选中，因此还需减去目标anchor元素的高度，即offsetHeight
    // 这样的使用体验会达到最佳.
    props.catalogue.some(item => {
        let targetElem = document.getElementById(`${item.name}`);
        let offsetParent = document.querySelector(`.article-reader .article-area`);
        if(documentScrollTop >= targetElem.offsetTop + offsetParent.offsetTop - targetElem.offsetHeight){
            currTag.value = item.name;
        }else{
            return true;
        }
    });
}

/* 锚点滚动逻辑 */
// 点击目录标签，平滑滚动至目标点，且距离越远，滚动速度越快
function scrollTo(targetAnchor){
    
    if(!shiftLock){
        return;
    }
    shiftLock = false;

    /* 1. 获取文档流的当前的scrollTop值，并计算目标scrollTop值 */
    let currScrollTop = document.documentElement.scrollTop || document.body.scrollTop;
    let targetElem = document.getElementById(targetAnchor);
    let offsetParent = document.querySelector(`.article-reader .article-area`);
    let dummyScrollTop = targetElem.offsetTop + offsetParent.offsetTop - targetElem.offsetHeight;
    // 单次滚动像素，为了画面滚动平滑又不至于太慢，需要根据目标锚点与当前位置距离进行计算
    let step = Math.max(Math.floor(Math.abs((currScrollTop - dummyScrollTop) / 60)), 30); 

    /* 2. 开滚 */
    stepScroll();
    function stepScroll(){
        currScrollTop = document.documentElement.scrollTop || document.body.scrollTop;
        if(currScrollTop !== dummyScrollTop){
            let delta = dummyScrollTop - currScrollTop;
            let movePixels = Math.abs(delta) > step ? step *  ((delta) / Math.abs(delta)): delta;
            document.documentElement.scrollTop += movePixels;
            document.body.scrollTop += movePixels;
            window.requestAnimationFrame(stepScroll);
        }else{
            shiftLock = true;
        }
    }
}



</script>

<style>
.catalogue-board{
    display: flex;
    flex-direction: column;
    width: 100%;
    height: 100%;
    border-radius: 20px;
    background-color: #111827;
    opacity: .7;
    color: white;
    .catalogue-header{
        text-align: center;
        padding: 32px 24px 24px 24px;
        margin-bottom: 10px;
        border-radius: 20px;
        box-shadow: 0 0 1.1px white;
        font-size: 24px;
        & > p{
            margin: 0;
            user-select: none;
        }
    }
    .catalogue-content{
        display: flex;
        flex-direction: column;
        overflow: auto;
        padding-bottom: 10px;
        & > li{
            display: flex;
            list-style: none;
            padding: 10px 5px 5px 0;
        }
        & > li:hover, & > .curr-tag{
            background-color: black;
        }
        .a-h1, .a-h2, .a-h3{
            width: 100%;
            color: white;
            text-decoration: none;
            cursor: pointer;
            user-select: none;
        }
        .a-h1{
            padding-left: 15px;
        }
        .a-h2{
            padding-left: 30px;
        }
        .a-h3{
            padding-left: 60px;
        }
        /*滚动条美化，非标准语法，Firefox PC/Mobile以及Safari均不支持*/
        &::-webkit-scrollbar{ /*滚动条宽度*/
            width: 2px;
            height: 1px;
        }
        &::-webkit-scrollbar-track{ /*滚动条轨道*/
            background-color: transparent;
            border-radius: 2em;
        }
        &::-webkit-scrollbar-thumb{ /*滚动条滑块*/
            background-color: rgb(255, 255, 255, .7);
            border-radius: 2em;
        }
    }
}
</style>
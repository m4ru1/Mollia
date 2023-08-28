<template>
    <div class="article-reader">
        <div class="article-info">
            <div class="title-area">
                <a href="#">{{ articleTitle }}</a>
            </div>
            <div class="info-area">
                <i class="date"> {{ date }} </i>
                <span class="divider">|</span>
                <i class="tag"> {{ tag }} </i>
                <span class="divider">|</span>
                <i class="word-count"> {{ wordCount.toString() + (wordCount > 1 ? ' words' : ' word') }} </i>
                <span class="divider">|</span>
                <i class="read-time"> {{ readTime.toString() + (readTime > 1 ? ' minitus' : ' minitu') }} </i>
            </div>
        </div>
        <div class="article-area">
            <div v-html="marked(articleData)" class="article-data"></div>
        </div>
        <div id="img-view">
            <img src="#" alt="">
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import hljs from 'highlight.js';
import { marked } from 'marked';
import 'highlight.js/styles/atom-one-dark-reasonable.css';

// 加载文章数据
const props = defineProps({
    articleTitle: String,
    date: String,
    tag: String,
    wordCount: Number,
    readTime: Number,
    articleData: String 
})

// 文章数据渲染
let articleData = ref(props.articleData);
onMounted(() => {
    /* 手动渲染代码高亮 */
    Array.from(document.getElementsByTagName('pre')).forEach(element => {
        Array.from(element.children).forEach((innerItem) => {hljs.highlightElement(innerItem);});
    });

    /* 为所有图像设置click监听事件 */
    document.querySelectorAll('.article-data img').forEach((imgItem) => {
        imgItem.onclick = function(){
            document.querySelector('#img-view img').setAttribute('src', this.getAttribute('src'));
            document.querySelector('#img-view').style.display = 'flex';
        };
    });

    /* 为图像展示容器设置click监听事件 */
    let imgView = document.getElementById('img-view');
    if (imgView){
        imgView.onclick = function(){
            document.querySelector('#img-view img').setAttribute('src', '#');
            this.style.display = 'none';
        }
    }
});
</script>

<style>

/* 文章展示容器样式 */
.article-reader{
    padding: 40px 20px;
    background-color: rgba(255, 255, 255, .8);
    border-radius: 20px;
}

/* 文章样式 */
.article-area{
    position: relative;
    .article-data  p{
        font-family: "Lato";
        margin: 1.5em auto 0 auto;
        line-height: 1.5em;
    }
    .article-data  ol  li{
        margin-top: 1em;
        line-height: 1.5em;
    }
    .article-data img{
        position: relative;
        display: block;
        width: 100%;
        margin: 10px auto 10px auto;
        pointer-events: auto;
        cursor: url("../../assets/icon/enlarge.ico"), auto;
    }
    .article-data > p > img{
        margin: 0px auto 10px auto;
    }

    .article-data .hljs{
        border-radius: 15px;
    }

    .article-data code:not(.hljs){
        color: red;
        text-shadow: .5px 0px 0px red;
    }
}


/* 图像展示容器样式 */
.article-reader #img-view{
    display: none;
    position: fixed;
    height: 100%;
    width: 100%;
    top: 0%;
    left: 0%;
    background-color: rgba(0, 0, 0, .5);
    text-align: center;
    align-items: center;
    justify-content: center;
    z-index: 100;
    cursor: url("../../assets/icon/shrink.ico"), auto;
    & > img{  /* 图像全局居中放大 */
        user-select: none;
        position: fixed;
    }
}

/* 文章标题简介的样式复用自ArticleEntry组件，
   请仅关心上面部分文章样式 :) */
.article-reader .article-info{ 
    /*TODO: 整个article-info的样式拷贝自Article-entry, 后期改进SCSS，实现样式复用*/
    display: flex;
    flex-grow: 5;
    padding: 0 20px 20px 20px;
    flex-direction: column;
    justify-content: start;
    text-align: center;
}

/*文章标题样式*/
.title-area{
    display: flex;
    justify-content: center;
    margin-top: 30px;
    overflow: hidden;
}

.title-area > a{
    position: relative;
    display: block;
    font-size: 35px;
    max-width: 710px;
    max-height: 80px;
    color: black;
    font-size: 2em;
    text-decoration: none;
    user-select: none;
    transition: color .3s ease-in-out;
}

.title-area > a:hover{
    color: #10b981;
    text-shadow: 0px 0px .2px #10b981;
}

.title-area > a::after{/*文章标题下划线动效*/
    content: "";
    position: absolute;
    width: 0%;
    height: 3px;
    top: 35px;
    left: 50%;
    right: 50%;
    color: #10b981;
    background-color: black;
    transition: all .3s ease-in-out;
}

.title-area > a:hover::after{
    left: 0%;
    right: 0%;
    width: 100%;
}

/*文章简介信息演示设置*/
.info-area{
    display: inline;
    /* min-width: 365px; */
    margin: 10px auto 10px auto;
}

.info-area > i{ /*简介内容*/
    color: rgb(130, 130, 130);
    font-style: normal;
}

.info-area > .divider{ /*简介分割线*/
    margin: 0 8px 0 8px;
    color: black;
    user-select: none;
}

/*简介信息前置图标*/
.info-area > .date::before{
    font-family: 'FontAwesome';
    content: "\f133";
    margin-right: 5px;
}

.info-area > .tag::before{
    font-family: 'FontAwesome';
    content: "\f114";
    margin-right: 5px;
}

.info-area > .word-count::before{
    font-family: 'FontAwesome';
    content: "\f1c2";
    margin-right: 5px;
}

.info-area > .read-time::before{
    font-family: 'FontAwesome';
    content: "\f017";
    margin-right: 5px;
}
</style>
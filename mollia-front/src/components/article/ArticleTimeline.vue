<template>
<div class="article-timeline">
    <div class="article-box">
        <div class="timeline-header">
            <h2>
                {{ `${total}` + `${total > 1 ? ' articles' : ' article'}` + ' in total.' }}
            </h2>
        </div>
        <div class="year-area" v-for="(yearData, year) in groupedArticles" :key="year">
            <div class="article-tag">
                <h2>
                    {{ year }}
                </h2>
            </div>
            <a href="#" class="article-node" v-for="article in yearData" :key="article.id">
                <h1>
                    {{ `${new Date(article.createdAt).getMonth() + 1}-${new Date(article.createdAt).getDate()}`}}&emsp;
                    <small> {{ article.title }} </small>
                </h1>
            </a>
        </div>
    </div>
</div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { getArticles } from '../../api';

const articles = ref([]);
const total = ref(0);

const groupedArticles = computed(() => {
    const groups = {};
    articles.value.forEach(article => {
        const year = new Date(article.createdAt).getFullYear();
        if (!groups[year]) {
            groups[year] = [];
        }
        groups[year].push(article);
    });
    return groups;
});


onMounted(async () => {
    try {
        const response = await getArticles({ page: 1, size: 100 }); // 获取前100篇文章作为示例
        articles.value = response.list;
        total.value = response.total;
    } catch (error) {
        console.error("获取文章列表失败:", error);
    }
});

function getYearStr(date){
    return `${date.getYear() + 1900}`;
}

function getMonthStr(date){
    return `${date.getMonth()}`;
}

function getDayStr(date){
    return `${date.getDay()}`;
}
</script>

<style>
.article-timeline{
    box-sizing: border-box;
    background-color: #eeeeeeca;
    border-radius: 20px;
    & >.article-box{
        position: relative;
        max-width: 42vw;
        height: auto;
        margin-left: 55px;
        padding-top: 40px;
        padding-bottom: 40px;
        border-radius: 20px;
        user-select: none;
    }
}

.article-timeline .article-box .timeline-header{
    position: relative;
    & > h2{
        font-size:17px;
        margin: 0 0 15px 20px;
    }
}

.article-timeline .article-box{
    .year-area{
        margin-bottom: 15px;
    }
    .article-tag{
        position: relative;
        & > h2{
            font-size: 24px;
            margin: 0 0 15px 20px;
        }
    }
    .article-node{
        position: relative;
        display: block;
        margin: 30px 0;
        text-decoration: none;
        line-height: 1.5em;
    }
    .timeline-header::before,
    .article-tag::before,
    .article-node::before{
        background-color: #777;
        content: " ";
        position: absolute;
        left: 0%;
        top: 50%;
        border-radius: 50%;
        z-index: 1;
    }
    .timeline-header::before{
        margin-left: -4px;
        margin-top: -4px;
        width: 8px;
        height: 8px;
    }
    .article-tag::before{
        margin-left: -4px;
        margin-top: -5px;
        width: 8px;
        height: 8px;
    }
    .article-node::before{
        margin-left: -2px;
        margin-top: -3px;
        width: 5px;
        height: 5px;
        transition: all .2s ease-in;
    }
    .article-node::after{ /* 条目下划线 */
        position: absolute;
        content: "";
        height: 1px;
        top: 1.5em;
        left: 20px;
        right: 100%;
        background-color: #555;
        transition: all .2s ease-in;
    }
    .article-node:hover::after{
        right: 5%;
    }

    .article-node h1{
        color: #555;
        font-size: 15px;
        margin: 0 0 0 20px;
    }

    .article-node h1 small{
        font-size: 17px;
        transition: all .3s ease-in;
    }

    .article-node:hover h1{
        color: #000;
    }
    .article-node:hover::before{
        background-color: #10b981;
        margin-left: -2.5px;
        width: 6px;
        height: 6px;
    }
}

.article-timeline .article-box::after{ /* 左侧长竖线 */
    position: absolute;
    content: " ";
    background-color: #55555530;
    top: 4.5%;
    left: 0;
    margin-left: -1px;
    width: 3px;
    height: 91%;
    z-index: 0;
}

</style>
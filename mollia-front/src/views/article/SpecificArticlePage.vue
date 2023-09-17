<template>
    <div class="specific-article-page">
        <div class="left-area">
            <ArticleReader
                :article-title="articleTitle"
                :date="date"
                :tag="tag"
                :word-count="wordCount"
                :read-time="readTime"
                :article-data="testData"
            >
            </ArticleReader>
        </div>
        <div class="right-area">
            <ProfileBoard></ProfileBoard>
            <CatalogueBoard v-if="isCatalogueDone" s="11"
            :catalogue="flatCatalogueData"></CatalogueBoard>
            <CatalogueBoard v-else s="22"></CatalogueBoard>
        </div>
    </div>
</template>

<script setup>
import ArticleReader from '../../components/article/ArticleReader.vue'
import CatalogueBoard from '../../components/catalogue/CatalogueBoard.vue';
import ProfileBoard from '../../components/profile/ProfileBoard.vue';

import { testData } from './testData';
import { ref, reactive, onMounted, onUnmounted } from 'vue';

let articleTitle = ref("记录：改了网安国密库的一个小Bug并提交PR");
let date = ref("2023/08/15");
let tag = ref("开源项目");
let wordCount = ref(125);
const readTime = Math.max(1, Math.floor(wordCount.value / 150));
let isCatalogueDone = ref(false);

// 文章数据渲染
let catalogueData = reactive([]);     // 结构化目录数据
let flatCatalogueData = reactive([]); // 扁平化目录数据 [√]
onMounted(() => {
    /* 目录数据生成, 以及目录定位标签插入 */
    let titleCollection = document.querySelectorAll('.article-reader .article-area .article-data h1,h2,h3');
    let currH1Pointer, currH2Pointer = null;
    for(let i = 0; i < titleCollection.length; i++){
        let item = titleCollection[i];
        switch (item.nodeName) { // 对于H1外的H2\H3，H2外的H3，视为非法结构，不予记入目录
            case 'H1':
                catalogueData.push({type: 'h1', name: item.innerText, h2: []});
                flatCatalogueData.push({type: 'h1', name: item.innerText});
                currH1Pointer = catalogueData[catalogueData.length - 1];
                currH2Pointer = null;
                item.setAttribute('id', `${item.innerText}`);
                break;
            case 'H2':
                if (currH1Pointer){
                    currH1Pointer.h2.push({type: 'h2', name: item.innerText, h3:[]});
                    flatCatalogueData.push({type: 'h2', name: item.innerText});
                    currH2Pointer = currH1Pointer.h2[currH1Pointer.h2.length - 1];
                    item.setAttribute('id', `${item.innerText}`);
                }
                break;
            case 'H3':
                if(currH2Pointer){
                    currH2Pointer.h3.push({type: 'h3', name: item.innerText});
                    flatCatalogueData.push({type: 'h3', name: item.innerText});
                    item.setAttribute('id', `${item.innerText}`);
                }
                break;
    }
    isCatalogueDone.value = true;
}});


</script>
<style scoped>

.specific-article-page{
    display: flex;
    justify-content: center;
    & > .left-area, & > .right-area {
        display: flex;
        flex-direction: column;
        margin: 30px;
    }
    .left-area{
        display: flex;
        flex-grow: 0;
        flex-basis: 35vw;
        max-width: 45vw;
        flex-shrink: 0;

    }
    .right-area{
        flex-grow: 0;
        flex-basis: 17vw;
        flex-shrink: 0;
        transition: all .5s ease-in-out;
        .catalogue-board {
            position: sticky;
            top: 65px;
            margin: 30px auto 0 auto;
            max-height: 350px;
        }
    }
}

/*针对移动窄屏设备的响应式设计*/
@media screen and (max-width: 800px){
    .specific-article-page{
        flex-direction: column;
        & > .left-area, & > .right-area {
            flex-grow: 1;
            max-width: 100%;
            margin-left: 15px;
            margin-right: 15px;
        }
    }
    .catalogue-board{
        margin-right: 50px;
    }
}

@media screen and (max-width: 1300px) and (min-width: 800px){
    .specific-article-page{
        flex-direction: column;
        align-items: center;
        & > .left-area, & > .right-area {
            flex-grow: 1;
            max-width: 80vw;
            margin-left: 10px;
            margin-right: 10px;
        }
        .left-area{
            width: 100%;
            flex-grow: 1;
        }
        .right-area{
            width: 100%;
            flex-direction: row;
            .profile-board{
                width: 50%;
                margin: 0;
            }
            .catalogue-board{
                width: 50%;
                margin: 0 0 0 10px;
            }
        }
    }
}
</style>
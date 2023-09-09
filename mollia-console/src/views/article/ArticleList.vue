<template>
    <div class="article-list">
        <PageContent class="page-content">
            <template #page-title>
                <h2>文章列表</h2>
            </template>
            <template #page-view>
                <div class="filter-bar">
                    <el-menu mode="horizontal" :ellipsis="false">
                        <span style="margin: auto 0;">文章状态：</span>
                        <el-menu-item index="all">全部</el-menu-item>
                        <el-menu-item index="public">公开</el-menu-item>
                        <el-menu-item index="private">私密</el-menu-item>
                        <el-menu-item index="drafts">草稿箱</el-menu-item>
                        <el-menu-item index="recycle">回收站</el-menu-item>
                    </el-menu>
                </div>
                <div class="search-bar">
                    <span> 标题：</span>
                    <el-input v-model="articleKeyWord" 
                        style="width: 200px; margin-right: 20px;"
                        placeholder="请输入文章标题" clearable>
                    </el-input>
                    <span> 类别：</span>
                    <el-select v-model="articleCategory"
                        style="width: 200px; margin-right: auto;"
                        placeholder="请输入文章类别" clearable>
                        <el-option 
                            v-for="(item, index) in articleCategoryChoices"
                            :key = "index"
                            :label = "item"
                            :value = "item"
                        />
                    </el-select>
                    <el-button @click="searchHandle" type="primary">
                        检索
                    </el-button>
                </div>
                <div class="article-table">
                    <el-table v-if="articleData.length" :data="articleData" border style="width: 100;">
                        <el-table-column prop="title" label="标题"></el-table-column>
                        <el-table-column prop="category" label="分类"></el-table-column>
                        <el-table-column prop="viewNumber" label="浏览量"></el-table-column>
                        <el-table-column prop="createDate" label="创建日期"></el-table-column>
                        <el-table-column label="置顶">
                            <template v-slot="{ row }">
                                <el-switch v-model="row.isTop" @change="topHandle(row)" :loading="row.topLoading"/>
                            </template>
                        </el-table-column>
                        <el-table-column prop="" label="操作">
                            <template v-slot="{ row }">
                                <el-button @click="editHandle(row)" type="primary"> 编辑 </el-button>
                                <el-button @click="deleteHandle(row)" type="danger"> 删除 </el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                    <el-empty v-else description="这里空空如也"></el-empty>
                </div>
                <div class="pagination-box">
                    <el-pagination background layout="prev, pager, next" :total="articleData.length"></el-pagination>
                </div>
            </template>
        </PageContent>
    </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import PageContent from '../../components/layout/PageContent.vue';

// 文章状态筛选
let perPage = ref(7);  // 文章数据列表
let articleData = ref([]);  // 文章数据列表
let activeArticleType = ref('all'); // 当前展示文章类型 【全部，公开，私密，草稿箱，回收站】
let activeArticleData = computed(()=>{ // 当前展示文章列表
    return articleData.value.filter((item) => item.type === activeArticleType.value || activeArticleType.value === 'all');
});

// 文章检索
let articleKeyWord = ref('');          // 待检索文章关键字
let articleCategory = ref('');         // 待检索文章类别
let articleCategoryChoices = ref([]);  // 现有文章类别

function searchHandle(){ // 检索逻辑

}

// 文章操作
function topHandle(row){ // 置顶逻辑
    console.log(row);
}

function editHandle(row){  // 打开编辑页面逻辑
    console.log(row);
}

function deleteHandle(row){  // 删除页面逻辑
    console.log(row);
}

// 测试数据加载
for( let i = 0; i < 50; i++){
    articleData.value.push({
        title: 'Python异步编程入门', 
        category: 'Python', 
        viewNumber: `32${i}`,
        createDate: `2022-08-${i}`, 
        isTop: false,
        topLoading: false
    });
}
// articleData.value.push({});
// articleData.value.push({});
// articleData.value.push({});
// articleData.value.push({});
// articleData.value.push({});
// articleData.value.push({});
// articleData.value.push({});
// articleCategoryChoices.value.push();
// articleCategoryChoices.value.push();
// articleCategoryChoices.value.push();
// articleCategoryChoices.value.push();
// articleCategoryChoices.value.push();
</script>

<style scoped>
.article-list{
    display: flex;
    overflow: hidden;
    position: relative;
    background-color: var(--theme-pagecontent-bg-color);
    .page-content{
        .filter-bar{
            padding: 15px 15px 0 15px;
            user-select: none;
        }
        .search-bar{
            display: flex;
            flex-direction: row;
            padding: 15px 20px;
            margin: 15px;
            border-radius: 5px;
            background-color: #eee;
            user-select: none;
            & > span{
                margin: auto 0;
            }
            .search-btn{
                background-color: var(--theme-primary-color);
                width: 80px;
                border-radius: 8px;
            }
        }
        .article-table{
            margin: 15px 15px;
            max-width: 90vw;
            min-width: 1000px;
        }
        .pagination-box{
            display: flex;
            justify-content: end;
            padding: 5px 30px 15px 15px;
        }
    }
}
</style>
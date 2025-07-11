<template>
    <div class="article-list">
        <PageContent class="page-content">
            <template #page-title>
                <div class="page-title-container">
                    <h2>文章列表</h2>
                    <el-button type="primary" @click="createHandle">新建文章</el-button>
                </div>
            </template>
            <template #page-view>
                <!-- filter-bar and search-bar are kept for future implementation -->
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
                    <el-table v-loading="loading" :data="articles" border style="width: 100%;">
                        <el-table-column prop="title" label="标题"></el-table-column>
                        <el-table-column prop="status" label="状态">
                             <template v-slot="{ row }">
                                <el-tag :type="row.status === 'published' ? 'success' : 'info'">
                                    {{ row.status === 'published' ? '已发布' : '草稿' }}
                                </el-tag>
                            </template>
                        </el-table-column>
                        <el-table-column prop="createdAt" label="创建日期">
                            <template v-slot="{ row }">
                                {{ new Date(row.createdAt).toLocaleString() }}
                            </template>
                        </el-table-column>
                        <el-table-column prop="updatedAt" label="更新日期">
                            <template v-slot="{ row }">
                                {{ new Date(row.updatedAt).toLocaleString() }}
                            </template>
                        </el-table-column>
                        <el-table-column label="操作" width="200">
                            <template v-slot="{ row }">
                                <el-button @click="editHandle(row)" type="primary" size="small"> 编辑 </el-button>
                                <el-button @click="deleteHandle(row)" type="danger" size="small"> 删除 </el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                    <el-empty v-if="!loading && !articles.length" description="这里空空如也"></el-empty>
                </div>
                <div class="pagination-box">
                    <el-pagination 
                        background 
                        layout="prev, pager, next, total" 
                        :total="total"
                        v-model:current-page="currentPage"
                        :page-size="pageSize"
                        @current-change="fetchArticles"
                    />
                </div>
            </template>
        </PageContent>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessageBox, ElMessage } from 'element-plus';
import PageContent from '../../components/layout/PageContent.vue';
import { getAdminArticles, deleteArticle } from '../../api';

const router = useRouter();

const articles = ref([]);
const total = ref(0);
const currentPage = ref(1);
const pageSize = ref(10);
const loading = ref(false);

// TODO: The following refs are for future filter implementation
let articleKeyWord = ref('');          // 待检索文章关键字
let articleCategory = ref('');         // 待检索文章类别
let articleCategoryChoices = ref([]);  // 现有文章类别

function searchHandle(){ // 检索逻辑

}

// 获取文章列表
const fetchArticles = async () => {
    loading.value = true;
    try {
        const response = await getAdminArticles({ page: currentPage.value, size: pageSize.value });
        articles.value = response.list;
        total.value = response.total;
    } catch (error) {
        ElMessage.error('获取文章列表失败: ' + error.message);
    } finally {
        loading.value = false;
    }
};

// 在组件挂载时获取数据
onMounted(fetchArticles);

// 文章操作
function topHandle(row){ // 置顶逻辑
    console.log(row);
}

// 操作处理
const createHandle = () => {
    router.push('/article/create');
};

const editHandle = (row) => {
    router.push(`/article/${row.id}`);
};

const deleteHandle = (row) => {
    ElMessageBox.confirm(`确定要删除文章《${row.title}》吗？`, '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
    }).then(async () => {
        try {
            await deleteArticle(row.id);
            ElMessage.success('删除成功');
            fetchArticles(); // 刷新列表
        } catch (error) {
            ElMessage.error('删除失败: ' + error.message);
        }
    }).catch(() => {
        // 用户取消操作
    });
};
</script>

<style scoped>
.page-title-container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
}
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
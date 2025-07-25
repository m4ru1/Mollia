<template>
    <div class="article-category">
        <PageContent class="page-content">
            <template #page-title>
                <div class="page-title-container">
                    <h2>文章分类</h2>
                    <el-button type="primary" @click="handleCreate">创建分类</el-button>
                </div>
            </template>
            <template #page-view>
                <div class="category-table">
                    <el-table v-loading="loading" :data="categories" border style="width: 100%;">
                        <el-table-column prop="id" label="ID" width="100"></el-table-column>
                        <el-table-column prop="name" label="类别名"></el-table-column>
                        <el-table-column prop="createdAt" label="创建日期">
                            <template v-slot="{ row }">
                                {{ new Date(row.createdAt).toLocaleString() }}
                            </template>
                        </el-table-column>
                        <el-table-column label="操作" width="200">
                            <template v-slot="{ row }">
                                <el-button @click="handleEdit(row)" type="primary" size="small">编辑</el-button>
                                <el-button @click="handleDelete(row)" type="danger" size="small">删除</el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                    <el-empty v-if="!loading && !categories.length" description="这里空空如也"></el-empty>
                </div>
            </template>
        </PageContent>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import PageContent from '../../components/layout/PageContent.vue';
import { getCategories, createCategory, updateCategory, deleteCategory } from '../../api';

const categories = ref([]);
const loading = ref(false);

// 获取分类列表
const fetchCategories = async () => {
    loading.value = true;
    try {
        categories.value = await getCategories();
    } catch (error) {
        ElMessage.error('获取分类列表失败: ' + error.message);
    } finally {
        loading.value = false;
    }
};

onMounted(fetchCategories);

// 创建分类
const handleCreate = () => {
    ElMessageBox.prompt('请输入新的分类名称', '创建分类', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputPattern: /.+/,
        inputErrorMessage: '分类名称不能为空',
    }).then(async ({ value }) => {
        try {
            await createCategory({ name: value });
            ElMessage.success('创建成功');
            fetchCategories();
        } catch (error) {
            ElMessage.error('创建失败: ' + error.message);
        }
    });
};

// 编辑分类
const handleEdit = (row) => {
    ElMessageBox.prompt('请输入新的分类名称', '编辑分类', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputValue: row.name,
        inputPattern: /.+/,
        inputErrorMessage: '分类名称不能为空',
    }).then(async ({ value }) => {
        try {
            await updateCategory(row.id, { name: value });
            ElMessage.success('更新成功');
            fetchCategories();
        } catch (error) {
            ElMessage.error('更新失败: ' + error.message);
        }
    });
};

// 删除分类
const handleDelete = (row) => {
    ElMessageBox.confirm(`确定要删除分类“${row.name}”吗？`, '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
    }).then(async () => {
        try {
            await deleteCategory(row.id);
            ElMessage.success('删除成功');
            fetchCategories();
        } catch (error) {
            ElMessage.error('删除失败: ' + error.message);
        }
    });
};
</script>

<style scoped>
.page-title-container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    padding: 0 15px;
}
.article-category {
    background-color: var(--theme-pagecontent-bg-color);
}
.category-table {
    padding: 15px;
    min-width: 800px;
}
</style>
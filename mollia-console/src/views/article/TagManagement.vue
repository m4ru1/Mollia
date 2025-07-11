<template>
    <div class="tag-management">
        <PageContent class="page-content">
            <template #page-title>
                <div class="page-title-container">
                    <h2>文章标签</h2>
                    <el-button type="primary" @click="handleCreate">创建标签</el-button>
                </div>
            </template>
            <template #page-view>
                <div class="tag-table">
                    <el-table v-loading="loading" :data="tags" border style="width: 100%;">
                        <el-table-column prop="id" label="ID" width="100"></el-table-column>
                        <el-table-column prop="name" label="标签名"></el-table-column>
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
                    <el-empty v-if="!loading && !tags.length" description="这里空空如也"></el-empty>
                </div>
            </template>
        </PageContent>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import PageContent from '../../components/layout/PageContent.vue';
import { getTags, createTag, updateTag, deleteTag } from '../../api';

const tags = ref([]);
const loading = ref(false);

// 获取标签列表
const fetchTags = async () => {
    loading.value = true;
    try {
        tags.value = await getTags();
    } catch (error) {
        ElMessage.error('获取标签列表失败: ' + error.message);
    } finally {
        loading.value = false;
    }
};

onMounted(fetchTags);

// 创建标签
const handleCreate = () => {
    ElMessageBox.prompt('请输入新的标签名称', '创建标签', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputPattern: /.+/,
        inputErrorMessage: '标签名称不能为空',
    }).then(async ({ value }) => {
        try {
            await createTag({ name: value });
            ElMessage.success('创建成功');
            fetchTags();
        } catch (error) {
            ElMessage.error('创建失败: ' + error.message);
        }
    });
};

// 编辑标签
const handleEdit = (row) => {
    ElMessageBox.prompt('请输入新的标签名称', '编辑标签', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputValue: row.name,
        inputPattern: /.+/,
        inputErrorMessage: '标签名称不能为空',
    }).then(async ({ value }) => {
        try {
            await updateTag(row.id, { name: value });
            ElMessage.success('更新成功');
            fetchTags();
        } catch (error) {
            ElMessage.error('更新失败: ' + error.message);
        }
    });
};

// 删除标签
const handleDelete = (row) => {
    ElMessageBox.confirm(`确定要删除标签“${row.name}”吗？`, '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
    }).then(async () => {
        try {
            await deleteTag(row.id);
            ElMessage.success('删除成功');
            fetchTags();
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
.tag-management {
    background-color: var(--theme-pagecontent-bg-color);
}
.tag-table {
    padding: 15px;
    min-width: 800px;
}
</style> 
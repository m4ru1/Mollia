<template>
    <div class="article-editor">
        <PageContent class="page-content">
            <template #page-view>
                <div class="editor-header">
                    <el-input class="editor-title-input" v-model="articleForm.title" placeholder="请输入文章标题"></el-input>
                    <div class="header-actions">
                        <el-button type="primary" @click="saveArticle('published')" :loading="loading">发布文章</el-button>
                        <el-button type="success" @click="saveArticle('draft')" :loading="loading">保存草稿</el-button>
                    </div>
                </div>
                <div class="meta-form">
                     <el-form :model="articleForm" label-width="80px" inline>
                        <el-form-item label="分类">
                            <el-select v-model.number="articleForm.categoryId" placeholder="请选择分类" filterable>
                                <el-option v-for="category in categories" :key="category.id" :label="category.name" :value="category.id"></el-option>
                            </el-select>
                        </el-form-item>
                         <el-form-item label="标签">
                            <el-select v-model="articleForm.tags" multiple filterable allow-create default-first-option placeholder="请输入或创建标签">
                            </el-select>
                        </el-form-item>
                    </el-form>
                </div>
                <div id="editor"></div>
            </template>
        </PageContent>
    </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import PageContent from '../../components/layout/PageContent.vue';
import Cherry from 'cherry-markdown';
import 'cherry-markdown/dist/cherry-markdown.css';
import { getAdminArticleById, createArticle, updateArticle, getCategories, uploadFile } from '../../api';

const route = useRoute();
const router = useRouter();

const articleId = ref(route.params.id || null);
const isEditMode = computed(() => !!articleId.value);
const loading = ref(false);
let cherry;

// 文章表单数据
const articleForm = reactive({
    title: '',
    content: '',
    categoryId: null,
    tags: [],
    status: 'draft',
});

const categories = ref([]); // 存储分类列表

// 初始化编辑器
onMounted(() => {
    fetchCategories(); // 获取分类列表
    cherry = new Cherry({
        id: 'editor',
        value: articleForm.content,
        callback: {
            afterChange: (md) => {
                articleForm.content = md;
            }
        },
        fileUpload: async (file, callback) => {
            const formData = new FormData();
            formData.append('file', file);
            try {
                const res = await uploadFile(formData);
                // 调用回调函数，第一个参数是URL，第二个参数是文件名（可选）
                callback(res.url, file.name);
            } catch (error) {
                ElMessage.error('图片上传失败: ' + error.message);
            }
        },
    });

    if (isEditMode.value) {
        loadArticleData();
    }
});

// 加载文章数据（编辑模式）
const loadArticleData = async () => {
    loading.value = true;
    try {
        const data = await getAdminArticleById(articleId.value);
        articleForm.title = data.title;
        articleForm.content = data.content;
        articleForm.categoryId = data.category; // This is a name, need to find ID
        articleForm.tags = data.tags;
        cherry.setValue(data.content);

        // Find category ID from name
        const category = categories.value.find(cat => cat.name === data.category);
        if (category) {
            articleForm.categoryId = category.id;
        } else {
            ElMessage.warning('文章原分类未找到');
        }

    } catch (error) {
        ElMessage.error('加载文章失败: ' + error.message);
    } finally {
        loading.value = false;
    }
};

// 获取分类列表
const fetchCategories = async () => {
    try {
        categories.value = await getCategories();
        // 如果是编辑模式，且文章数据已加载，确保分类ID是有效的
        if (isEditMode.value && articleForm.categoryId) {
            const found = categories.value.some(cat => cat.id === articleForm.categoryId);
            if (!found) {
                ElMessage.warning('文章原分类不存在，请重新选择');
                articleForm.categoryId = null;
            }
        }
    } catch (error) {
        ElMessage.error('获取分类列表失败: ' + error.message);
    }
};

// 保存文章（发布或存为草稿）
const saveArticle = async (status) => {
    loading.value = true;
    articleForm.status = status;

    try {
        if (isEditMode.value) {
            // 更新文章
            await updateArticle(articleId.value, articleForm);
            ElMessage.success('文章更新成功');
        } else {
            // 创建新文章
            const response = await createArticle(articleForm);
            ElMessage.success('文章创建成功');
            router.push(`/article/${response.id}`); // 创建成功后跳转到编辑页
        }
    } catch (error) {
        ElMessage.error('保存失败: ' + error.message);
    } finally {
        loading.value = false;
    }
};

</script>

<style scoped>
.article-editor {
    display: flex;
    overflow: hidden;
    position: relative;
    background-color: var(--theme-pagecontent-bg-color);
}
.page-content {
    display: flex;
    flex-direction: column;
}
.editor-header {
    display: flex;
    align-items: center;
    padding: 20px 15px 15px 15px;
}
.editor-title-input {
    flex-grow: 1;
    margin-right: 20px;
}
.meta-form {
    padding: 0 15px 15px 15px;
}
#editor {
    flex-grow: 1;
    height: calc(100vh - 250px); /* Adjust height as needed */
    padding: 0 15px 10px 15px;
}
</style>
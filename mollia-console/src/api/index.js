import axios from 'axios';
import { useProfileStore } from '../stores/ProfileStore';

const apiClient = axios.create({
    baseURL: 'http://localhost:8080/api/v1',
    headers: {
        'Content-Type': 'application/json',
    },
});

// 请求拦截器：在每个请求前附加 Token
apiClient.interceptors.request.use(
    config => {
        const profileStore = useProfileStore();
        const token = profileStore.token;
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    error => {
        return Promise.reject(error);
    }
);

// 响应拦截器
apiClient.interceptors.response.use(
    response => {
        if (response.data.code === 200) {
            return response.data.data;
        }
        return Promise.reject(new Error(response.data.msg || 'Error'));
    },
    error => {
        // 在这里可以处理类似 401 未授权的全局错误
        if (error.response && error.response.status === 401) {
            // TODO: 跳转到登录页
            const profileStore = useProfileStore();
            profileStore.logout(); // 清除过期的 token
        }
        return Promise.reject(error);
    }
);

/**
 * 登录
 * @param {Object} credentials { username, password }
 * @returns {Promise<Object>}
 */
export const login = (credentials) => {
    return apiClient.post('/auth/login', credentials);
};

// --- 文章管理 ---

/**
 * 获取后台文章列表
 * @param {Object} params { page, size }
 */
export const getAdminArticles = (params) => {
    return apiClient.get('/admin/articles', { params });
};

/**
 * 根据ID获取文章详情（后台）
 * @param {number} id
 */
export const getAdminArticleById = (id) => {
    return apiClient.get(`/admin/articles/${id}`);
};

/**
 * 创建新文章
 * @param {Object} articleData
 */
export const createArticle = (articleData) => {
    return apiClient.post('/admin/articles', articleData);
};

/**
 * 更新文章
 * @param {number} id
 * @param {Object} articleData
 */
export const updateArticle = (id, articleData) => {
    return apiClient.put(`/admin/articles/${id}`, articleData);
};

/**
 * 删除文章
 * @param {number} id
 */
export const deleteArticle = (id) => {
    return apiClient.delete(`/admin/articles/${id}`);
};

// --- 分类管理 ---

/**
 * 获取所有分类
 */
export const getCategories = () => {
    return apiClient.get('/admin/categories');
};

/**
 * 创建新分类
 * @param {Object} categoryData { name }
 */
export const createCategory = (categoryData) => {
    return apiClient.post('/admin/categories', categoryData);
};

/**
 * 更新分类
 * @param {number} id
 * @param {Object} categoryData { name }
 */
export const updateCategory = (id, categoryData) => {
    return apiClient.put(`/admin/categories/${id}`, categoryData);
};

/**
 * 删除分类
 * @param {number} id
 */
export const deleteCategory = (id) => {
    return apiClient.delete(`/admin/categories/${id}`);
};

// --- 标签管理 ---

/**
 * 获取所有标签
 */
export const getTags = () => {
    return apiClient.get('/admin/tags');
};

/**
 * 创建新标签
 * @param {Object} tagData { name }
 */
export const createTag = (tagData) => {
    return apiClient.post('/admin/tags', tagData);
};

/**
 * 更新标签
 * @param {number} id
 * @param {Object} tagData { name }
 */
export const updateTag = (id, tagData) => {
    return apiClient.put(`/admin/tags/${id}`, tagData);
};

/**
 * 删除标签
 * @param {number} id
 */
export const deleteTag = (id) => {
    return apiClient.delete(`/admin/tags/${id}`);
};

// --- 文件上传 ---

/**
 * 上传文件
 * @param {FormData} formData
 */
export const uploadFile = (formData) => {
    return apiClient.post('/admin/storage/upload', formData, {
        headers: {
            'Content-Type': 'multipart/form-data',
        },
    });
};

export default apiClient; 
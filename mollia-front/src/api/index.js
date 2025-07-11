import axios from 'axios';

const apiClient = axios.create({
    baseURL: 'http://localhost:8080/api/v1', // 假设后端服务在8080端口
    headers: {
        'Content-Type': 'application/json',
    },
});

// 统一响应处理
apiClient.interceptors.response.use(
    response => {
        // 如果响应码是200，则直接返回 data 部分
        if (response.data.code === 200) {
            return response.data.data;
        }
        // 否则，返回一个被拒绝的 Promise
        return Promise.reject(new Error(response.data.msg || 'Error'));
    },
    error => {
        // 处理网络错误等
        console.error('API Error:', error);
        return Promise.reject(error);
    }
);

/**
 * 根据 ID 获取单篇文章
 * @param {number} id 文章ID
 * @returns {Promise<Object>}
 */
export const getArticleById = (id) => {
    return apiClient.get(`/articles/${id}`);
};

/**
 * 获取文章列表
 * @param {Object} params { page, size }
 * @returns {Promise<Object>}
 */
export const getArticles = (params) => {
    return apiClient.get('/articles', { params });
};

// 未来可以继续在这里添加其他 API...
// export const getTags = () => apiClient.get('/tags');

export default apiClient; 
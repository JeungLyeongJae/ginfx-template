import axios from 'axios';
import {message} from "ant-design-vue";

const apiClient = axios.create({
    baseURL: 'http://127.0.0.1:9060',
    headers: {
        'Content-Type': 'application/json',
    },
    timeout: 5000, // 请求超时时间
});

// 添加请求拦截器
apiClient.interceptors.request.use(
    config => {
        const token = sessionStorage.getItem('token'); // 从本地存储中获取 token
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    error => {
        // 处理请求错误
        return Promise.reject(error);
    }
);

apiClient.interceptors.response.use(
    response => {
        // 直接返回响应数据
        return response;
    },
    error => {
        const originalRequest = error.config;
        // 检查是否是登录请求
        if (originalRequest.url.includes("/login")) {
            // 可以处理登录特有的错误
            return Promise.reject(error);
        } else if (error.response && error.response.status === 401) {
            // 如果用户未认证或令牌失效，重定向到登录页面
            message.warn('登录已失效，请重新登录').then(() => {
                window.location.href = '/login';
            });
        }
        // 返回一个被拒绝的承诺，以便可以捕获错误并进一步处理
        return Promise.reject(error);
    }
);

export default apiClient;

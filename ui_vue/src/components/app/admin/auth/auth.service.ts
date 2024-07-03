import axios from 'axios';

const apiClient = axios.create({
    baseURL: 'http://10.18.16.33:9060', // 替换为你的 API 基础 URL
    headers: {
        'Content-Type': 'application/json',
    },
});


// 登录
export const login = async (params: { username: string, password: string }): Promise<any> => {
    return (await apiClient.post<any>('/login', params));
};
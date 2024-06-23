import axios from 'axios';
import {User} from "./user.ts";

export interface UserPagination {
    page_number?: number;
    page_size?: number;
    condition?: string;
    total_count?: number;
    users?: User[];
}

const apiClient = axios.create({
    baseURL: 'http://127.0.0.1:9060', // 替换为你的 API 基础 URL
    headers: {
        'Content-Type': 'application/json',
    },
});

// 获取用户列表
export const getUsers = async (params: UserPagination): Promise<UserPagination> => {
    return (await apiClient.get<UserPagination>('/api/user/list', { params })).data;
};

// 编辑用户
export const updateUser = async (params: User) => {
    return (await apiClient.post<User>('/api/user/update',  params));
};
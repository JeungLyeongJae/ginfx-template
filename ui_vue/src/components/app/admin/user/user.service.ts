import {User} from "./user.ts";
import apiClient from "../../../../axios.service.ts";

export interface UserPagination {
    page_number?: number;
    page_size?: number;
    condition?: string;
    total_count?: number;
    users?: User[];
}

// 判断用户名重复
export const isExistUserName = async (params: { username: string }): Promise<boolean> => {
    return (await apiClient.get<boolean>('/api/user/is_exist_username', {params})).data;
};

// 获取用户列表
export const getUsers = async (params: UserPagination): Promise<UserPagination> => {
    return (await apiClient.get<UserPagination>('/api/user/list', {params})).data;
};

// 新增用户
export const addUser = async (params: User) => {
    return (await apiClient.post<User>('/api/user/add', params));
};

// 编辑用户
export const updateUser = async (params: User) => {
    return (await apiClient.post<void>('/api/user/update', params));
};

// 删除用户
export const deleteUser = async (params: { id: number }) => {
    return (await apiClient.get<void>('/api/user/delete', {params}));
};
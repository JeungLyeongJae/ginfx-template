import apiClient from "../../../../axios.service.ts";

// 登录
export const login = async (params: { username: string, password: string }): Promise<any> => {
    return ((await apiClient.post<any>('/login', params)).data);
};
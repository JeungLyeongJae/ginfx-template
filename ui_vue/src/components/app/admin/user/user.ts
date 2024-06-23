export class User {
    id?: number;
    name?: string;
    phone?: string;
    username?: string;
    password?: string;
    plain_password?: string;
    roles?: string;
    enable?: number;
    created_at?: string;
    updated_at?: string;
    last_login?: string;
}

export const UserMapping: Map<string, string> = new Map();

UserMapping.set('name', '姓名');
UserMapping.set('phone', '手机号');
UserMapping.set('username', '用户名');
UserMapping.set('enable', '是否启用');
UserMapping.set('created_at', '创建日期');
UserMapping.set('updated_at', '更新日期');
UserMapping.set('last_login', '上传登录日期');
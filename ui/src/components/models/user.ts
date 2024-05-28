export interface User {
    id: number;
    name: string;
    phone: string;
    username: string;
    password: string;
    plain_password: string;
    roles: string[];
    enable: number;
    created_at: Date;
    updated_at: Date;
    last_login: Date;
    deleted_at: Date;
}
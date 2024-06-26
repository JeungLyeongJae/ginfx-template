import {isExistUserName} from "../../src/components/app/admin/user/user.service.ts";

// 通用 用户名校验
export const validateUserName = async (_: any, validate_user_name: string) => {

    const nameRegex = /^[A-Za-z][A-Za-z0-9]*$/; // 必须以英文字母开头，可以包含数字，但不能全是数字
    if (!validate_user_name) {
        return Promise.reject(new Error('用户名不能为空!'));
    }
    if (!nameRegex.test(validate_user_name)) {
        return Promise.reject(new Error('用户名只能由数字和英文组成，且以英文开头!'));
    }
    if (validate_user_name.length <= 3) {
        return Promise.reject(new Error('用户名不能小于三个字符!'));
    }
    const response = await isExistUserName({username: validate_user_name});
    if (response) {
        return Promise.reject(new Error('用户名已存在，请选择其他用户名！'));
    } else {
        return Promise.resolve();
    }
};

// 通用 密码校验
export const validatePassword = (_: any, validate_password: string) => {
    const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!"#$%&'()*+,-./:;<=>?@[\\\]^_`{|}~])[A-Za-z\d!"#$%&'()*+,-./:;<=>?@[\\\]^_`{|}~]{8,}$/; // 密码正则表达式
    if (!validate_password) {
        return Promise.reject(new Error('密码不能为空!'));
    }
    if (!passwordRegex.test(validate_password)) {
        return Promise.reject(new Error('密码至少8位，且必须包含数字，字母大小写以及特殊符号!'));
    }
    return Promise.resolve();
};

// 通用 中国大陆手机号校验
export const validatePhone = (_: any, validate_phone: string) => {
    const phoneRegex = /^1[3-9]\d{9}$/; // 中国手机号验证
    if (!validate_phone) {
        return Promise.reject(new Error('手机号不能为空!'));
    }
    if (!phoneRegex.test(validate_phone)) {
        return Promise.reject(new Error('请输入正确的手机号!'));
    }
    return Promise.resolve();
};

// 通用 中国姓名校验
export const validateName = (_: any, validate_name: string) => {
    const nameRegex = /^[\u4e00-\u9fa5]{2,4}$/; // 中国姓名正则表达式，2到4个汉字
    if (!validate_name) {
        return Promise.reject(new Error('姓名不能为空!'));
    }
    if (!nameRegex.test(validate_name)) {
        return Promise.reject(new Error('姓名必须是2到4个汉字!'));
    }
    return Promise.resolve();
};

// 通用 邮箱校验
export const validateEmail = (_: any, validate_email: string) => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/; // 简单的邮箱正则表达式
    if (!emailRegex.test(validate_email)) {
        return Promise.reject(new Error('邮箱格式有误!'));
    }
    return Promise.resolve();
};

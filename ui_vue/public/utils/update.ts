import {Button, notification} from "ant-design-vue";
import { h } from 'vue';
import {SmileOutlined} from "@ant-design/icons-vue";

let time = 0;            // 计算轮询次数
let version = "";        // 缓存的版本号
let timer: number | null = null;  // 定时器句柄
let notificationTimes = 0; // 提示框弹出次数

// 轮询用检测方法
const timerFunction = async (): Promise<void> => {
    // 次数超过的时候 停止轮询 防止用户挂着网页一直轮询
    if (time >= 5) {
        // 仅清除计时器
        clearInterval(timer!);
        timer = null;
        return;
    }

    try {
        // fetch 部署后同层级的version文件 并且加上时间戳参数 防止去访问本地硬盘的缓存
        const res = await fetch(`/version.txt?v=${new Date().getTime().toString()}`);
        const data = await res.json();

        // 首次加载网页的时候 存储第一份version
        if (!version) {
            version = data;
        } else if (version !== data) {
            if (notificationTimes == 0) {
                openNotification()
            }
            return;
        }
    } catch (error) {
        return;
    }

    time++;
}

// 检测鼠标是否移动 移动代表用户活跃中 把轮询比较用的次数一直清0
const moveFunction = (): void => {
    time = 0;
    // 长时间挂机后 不在轮询的网页 在鼠标活跃于窗口的时候重新检测
    if (!timer) {
        timer = window.setInterval(timerFunction, 1000);
    }
}

// 完全清除轮询 不轮询 不监听鼠标事件
// const clearTimer = (): void => {
//     if (timer !== null) {
//         clearInterval(timer);
//         timer = null;
//     }
//     window.removeEventListener("mousemove", moveFunction);
// }

// 当被 main.ts 引用的时候 开始轮询于监听鼠标移动事件
timer = window.setInterval(timerFunction, 5000);
window.addEventListener("mousemove", moveFunction);

const close = () => {
    notificationTimes = 0;
};

const refresh = (key: string) => {
    notification.close(key);
    location.reload();  // 刷新了缓存
}

const openNotification = () => {
    notificationTimes += 1;
    const key = `open${Date.now()}`;
    notification.open({
        message: '系统更新',
        description:
            '系统已更新，请手动刷新页面（请在刷新前注意保存当前页面数据）。',
        icon: () => h(SmileOutlined, { style: 'color: #108ee9' }),
        duration: 0,
        style: {
            width: '600px',
            marginLeft: `${335 - 600}px`,
        },
        btn: () =>
            h(
                Button,
                {
                    type: 'primary',
                    size: 'small',
                    onClick: () => refresh(key),
                },
                { default: () => '刷新' },
            ),
        key,
        onClose: close,
    });
};

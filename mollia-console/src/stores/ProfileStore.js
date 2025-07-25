import { defineStore } from 'pinia';

export const useProfileStore = defineStore('profile-store', {
    state: () => {
        return {
            token: localStorage.getItem('token') || null,
            userName: 'm4ru1',                          // 用户名
            userAvatarURL: '',                          // 用户头像地址
            websiteRecordInfo: '鲁ICP备2022040119号',    // 网站备案信息
            sinceYear: '2023',                          // 网站起始年份
            githubAccount: 'https://github.com/m4ru1',  //GitHub个人主页
        }
    },
    actions: {
        setToken(token) {
            this.token = token;
            localStorage.setItem('token', token);
        },
        logout() {
            this.token = null;
            localStorage.removeItem('token');
        },
    },
});
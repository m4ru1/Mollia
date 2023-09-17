import { defineStore } from 'pinia';

export const useThemeStore = defineStore('theme-store', {
    state: () => {
        return {
            isCollapse: false,  // 导航栏菜单是否折叠
            isDark: false,      // 是否为暗夜模式
        }
    }
})
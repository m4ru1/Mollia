import { defineStore } from 'pinia';

export const useThemeStore = defineStore('theme-store', {
    state: () => {
        isCollapse: false;
    }
})
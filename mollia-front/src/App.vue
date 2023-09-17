<script setup>
import AppHeader from "./components/layout/AppHeader.vue"
import AppFooter from "./components/layout/AppFooter.vue";
</script>

<template>
    <div class="app" v-cloak>
        <AppHeader class="app-header"></AppHeader>
        <div class="background-img"></div>
        <div class="main">
            <router-view v-slot="{ Component }">
                <transition name="fade" mode="out-in">
                    <component :is="Component" />
                </transition>
            </router-view>
        </div>
        <AppFooter class="app-footer"></AppFooter>
    </div>
</template>

<style>

.background-img{
    position: fixed;
    width: 100vw;
    height: 100vh;
    z-index: -1;
    background:url("./src/assets/01.jpg") center center no-repeat; 
    background-size:cover;
    filter: blur(2px);
}

.main{
    position: relative;
    top: 60px;
    margin-bottom: 60px;
}

/*滚动条美化，非标准语法，Firefox PC/Mobile以及Safari均不支持*/
body{
    overflow-y: scroll;
}

body::-webkit-scrollbar{ /*滚动条宽度*/
    width: 4px;
    height: 1px;
}

body::-webkit-scrollbar-track{ /*滚动条轨道*/
    background-color: #1b1b1b;
}

body::-webkit-scrollbar-thumb{ /*滚动条滑块*/
    background-color: #10b981;
    border-radius: 90px;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease-in;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>

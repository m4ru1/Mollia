<template>
  <div class="app">
    <div class="app-nav">
      <NavBoard></NavBoard>
    </div>
    <div class="app-content">
      <AppHeader class="page-header"></AppHeader>
      <router-view v-slot="{ Component }" class="page-content">
          <transition name="fade" mode="out-in">
              <component :is="Component" />
          </transition>
      </router-view>
    </div>
</div>
</template>

<script setup>
import NavBoard from '../src/components/nav/NavBoard.vue';
import AppHeader from '../src/components/layout/AppHeader.vue';

// 禁用全局鼠标滚动事件监听 TODO: 


</script>

<style scoped>
.app{
  /* 主题变量 */
  --theme-primary-color: #10b981;
  --theme-text-color: #000000;
  --theme-pagecontent-bg-color: #ccc;
  --theme-pagetitle-text-color: #333;
  --theme-pageview-bg-color: white;

  position: relative;
  display: flex;
  flex-direction: row;
  width: 100vw;
  overflow: hidden;
  .app-nav{
    height: 100%;
  }
  .app-content{
    flex: 1;
    display: flex;
    height: 100vh;
    flex-direction: column;
    .page-header{
      background-color: #444;
      width: 100%;
      height: 100px;
    }
    .page-content{
      flex-grow: 1;
      height: 100%;
    }
  }
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>

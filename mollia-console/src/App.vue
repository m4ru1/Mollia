<template>
  <div class="app">
    <div class="app-nav">
      <NavBoard></NavBoard>
    </div>
    <div class="app-content">
      <AppHeader class="header"></AppHeader>
      <router-view></router-view>
    </div>
</div>
</template>

<script setup>
import NavBoard from '../src/components/nav/NavBoard.vue';
import AppHeader from '../src/components/layout/AppHeader.vue';
</script>

<style scoped>
.app{
  /* */
  --theme-primary-color: #10b981;

  position: absolute;
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
    flex-direction: column;
    .header{
      background-color: #444;
      width: 100%;
      height: 100px;
    }
  }
}
</style>

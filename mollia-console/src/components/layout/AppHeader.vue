<template>
<div class="app-header">
    <div class="high-row">
        <Navigation class="navigation"></Navigation>
    </div>
    <div class="low-row">
        <PageDock class="page-dock"></PageDock>
    </div>
</div>    
</template>

<script setup>
import Navigation from './components/Navigation.vue';
import PageDock from './components/PageDock.vue';
</script>

<style>
.app-header{
    position: relative;
    display: flex;
    flex-direction: column;
    justify-content: center;
    user-select: none;
    .high-row, .low-row{
        flex: 1;
        display: flex;
        flex-direction: row;
        align-items: center;
    }
    .high-row{
        .navigation{
            flex-grow: 1;
            padding: 10px 0;
        }
    }
    .low-row{
        border-top: 1px solid #ccc;
        .page-dock{
            flex-grow: 1;
            margin: auto 0;
        }
    }
}
</style>